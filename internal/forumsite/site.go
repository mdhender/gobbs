package forumsite

import (
	"bytes"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/mdhender/gobbs/internal/setupjson"
	"github.com/microcosm-cc/bluemonday"

	_ "modernc.org/sqlite"
)

//go:embed templates/*
var embeddedTemplates embed.FS

var rawHTMLPolicy = newRawHTMLPolicy()

type Config struct {
	SQLitePath   string
	SetupPath    string
	TablePrefix  string
	SiteTitle    string
	BaseURL      string
	TemplatesDir string
	LiveTemplate bool
	TextFormats  setupjson.TextFormats
	Debug        struct {
		HighlightRawHTML bool
	}
}

type Renderer struct {
	cfg Config
	db  *sql.DB
}

type siteData struct {
	Title       string
	BaseURL     string
	GeneratedAt time.Time
	Forums      []*forum
	Threads     []*thread
	Stats       boardStats
}

type boardStats struct {
	PostsTotal   int64
	ThreadsTotal int64
	MembersTotal int64
	NewestMember *user
}

type forum struct {
	ID               int64
	ParentID         int64
	Parent           *forum
	Name             string
	DescriptionHTML  template.HTML
	DescriptionClass string
	Type             string
	DisplayOrder     int
	ThreadsCount     int64
	PostsCount       int64
	LastPost         time.Time
	LastPostSubject  string
	LastPostThread   int64
	LastPostAuthor   string
	Children         []*forum
	Threads          []*thread
}

type thread struct {
	ID         int64
	ForumID    int64
	Subject    string
	AuthorID   int64
	AuthorName string
	CreatedAt  time.Time
	LastPostAt time.Time
	Replies    int64
	Views      int64
	Sticky     bool
	Posts      []*post
	Forum      *forum
}

type post struct {
	ID             int64
	ThreadID       int64
	AuthorID       int64
	AuthorName     string
	CreatedAt      time.Time
	EditedAt       time.Time
	EditReason     string
	BodyHTML       template.HTML
	BodyClass      string
	Attachments    []*attachment
	Author         *user
	EditedAtString string
}

type attachment struct {
	ID            int64
	PostID        int64
	FileName      string
	StoredPath    string
	ThumbnailPath string
	FileType      string
	FileSize      int64
}

type user struct {
	ID         int64
	Name       string
	Title      string
	AvatarPath string
	PostCount  int64
}

type pageData struct {
	Site       *siteData
	PageTitle  string
	Section    string
	ErrorPage  *errorPage
	Forum      *forum
	Thread     *thread
	Forums     []*forum
	Breadcrumb []crumb
}

type errorPage struct {
	StatusCode int
	Heading    string
	Message    string
}

type crumb struct {
	Label string
	Href  string
}

func New(cfg Config) (*Renderer, error) {
	if cfg.SQLitePath == "" {
		cfg.SQLitePath = "mybb.sqlite3"
	}
	if cfg.SetupPath == "" {
		cfg.SetupPath = "setup.json"
	}
	if cfg.SiteTitle == "" {
		cfg.SiteTitle = "PlayByMail Forums Archive"
	}
	cfg.BaseURL = normalizeBaseURL(cfg.BaseURL)
	if cfg.TemplatesDir == "" {
		cfg.TemplatesDir = filepath.Join("internal", "forumsite", "templates")
	}

	db, err := sql.Open("sqlite", cfg.SQLitePath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite archive: %w", err)
	}

	setupCfg, err := setupjson.Parse(cfg.SetupPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		db.Close()
		return nil, err
	}
	if cfg.TablePrefix == "" {
		cfg.TablePrefix = setupCfg.Database.TablePrefix
	}
	if len(cfg.TextFormats) == 0 {
		cfg.TextFormats = setupCfg.TextFormats
	}
	cfg.Debug.HighlightRawHTML = setupCfg.Debug.HighlightRawHTML

	if cfg.TablePrefix == "" {
		cfg.TablePrefix, err = detectPrefix(db)
		if err != nil {
			db.Close()
			return nil, err
		}
	}

	return &Renderer{cfg: cfg, db: db}, nil
}

func (r *Renderer) Close() error {
	if r.db == nil {
		return nil
	}
	return r.db.Close()
}

func (r *Renderer) Build(outDir string) error {
	site, err := r.loadSite()
	if err != nil {
		return err
	}
	tmpl, err := r.loadTemplate()
	if err != nil {
		return err
	}
	css, err := r.loadCSS()
	if err != nil {
		return err
	}

	if err := os.RemoveAll(outDir); err != nil {
		return fmt.Errorf("clear output directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(outDir, "assets"), 0o755); err != nil {
		return fmt.Errorf("create assets directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(outDir, "f"), 0o755); err != nil {
		return fmt.Errorf("create forum directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(outDir, "t"), 0o755); err != nil {
		return fmt.Errorf("create thread directory: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "assets", "site.css"), css, 0o644); err != nil {
		return fmt.Errorf("write site stylesheet: %w", err)
	}

	indexData := pageData{
		Site:      site,
		PageTitle: site.Title,
		Section:   "home",
		Forums:    site.Forums,
	}
	if err := renderPage(tmpl, filepath.Join(outDir, "index.html"), indexData); err != nil {
		return err
	}
	for _, currentErrorPage := range errorPages(site) {
		if err := renderPage(tmpl, filepath.Join(outDir, fmt.Sprintf("%d.html", currentErrorPage.ErrorPage.StatusCode)), currentErrorPage); err != nil {
			return err
		}
	}

	for _, currentForum := range allForums(site.Forums) {
		dir := filepath.Join(outDir, "f", fmt.Sprintf("%d", currentForum.ID))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create forum output directory: %w", err)
		}
		if err := renderPage(tmpl, filepath.Join(dir, "index.html"), pageData{
			Site:       site,
			PageTitle:  currentForum.Name + " - " + site.Title,
			Section:    "forum",
			Forum:      currentForum,
			Breadcrumb: breadcrumbsForForum(site, currentForum),
		}); err != nil {
			return err
		}
	}

	for _, currentThread := range site.Threads {
		dir := filepath.Join(outDir, "t", fmt.Sprintf("%d", currentThread.ID))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create thread output directory: %w", err)
		}
		if err := renderPage(tmpl, filepath.Join(dir, "index.html"), pageData{
			Site:       site,
			PageTitle:  currentThread.Subject + " - " + site.Title,
			Section:    "thread",
			Thread:     currentThread,
			Breadcrumb: breadcrumbsForThread(site, currentThread),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (r *Renderer) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/assets/site.css", r.serveCSS)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	mux.HandleFunc("/", r.servePage)
	return mux
}

func (r *Renderer) serveCSS(w http.ResponseWriter, req *http.Request) {
	css, err := r.loadCSS()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Write(css)
}

func (r *Renderer) servePage(w http.ResponseWriter, req *http.Request) {
	site, err := r.loadSite()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := r.loadTemplate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := strings.TrimSuffix(req.URL.Path, "/")
	var data pageData
	switch {
	case path == "":
		data = pageData{
			Site:      site,
			PageTitle: site.Title,
			Section:   "home",
			Forums:    site.Forums,
		}
	case strings.HasPrefix(path, "/f/"):
		id, ok := parseID(path, "/f/")
		if !ok {
			renderHTTPErrorPage(w, tmpl, customErrorPageData(site, http.StatusNotFound))
			return
		}
		currentForum := findForum(site.Forums, id)
		if currentForum == nil {
			renderHTTPErrorPage(w, tmpl, customErrorPageData(site, http.StatusNotFound))
			return
		}
		data = pageData{
			Site:       site,
			PageTitle:  currentForum.Name + " - " + site.Title,
			Section:    "forum",
			Forum:      currentForum,
			Breadcrumb: breadcrumbsForForum(site, currentForum),
		}
	case strings.HasPrefix(path, "/t/"):
		id, ok := parseID(path, "/t/")
		if !ok {
			renderHTTPErrorPage(w, tmpl, customErrorPageData(site, http.StatusNotFound))
			return
		}
		currentThread := findThread(site.Threads, id)
		if currentThread == nil {
			renderHTTPErrorPage(w, tmpl, customErrorPageData(site, http.StatusNotFound))
			return
		}
		data = pageData{
			Site:       site,
			PageTitle:  currentThread.Subject + " - " + site.Title,
			Section:    "thread",
			Thread:     currentThread,
			Breadcrumb: breadcrumbsForThread(site, currentThread),
		}
	default:
		renderHTTPErrorPage(w, tmpl, customErrorPageData(site, http.StatusNotFound))
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buf.Bytes())
}

func (r *Renderer) loadTemplate() (*template.Template, error) {
	funcs := template.FuncMap{
		"forumLink": func(site *siteData, fid int64) string {
			return relLink(site.BaseURL, fmt.Sprintf("/f/%d/", fid))
		},
		"threadLink": func(site *siteData, tid int64) string {
			return relLink(site.BaseURL, fmt.Sprintf("/t/%d/", tid))
		},
		"initial": func(name string) string {
			name = strings.TrimSpace(name)
			if name == "" {
				return "?"
			}
			return strings.ToUpper(string([]rune(name)[0]))
		},
		"formatNumber": func(v int64) string {
			sign := ""
			if v < 0 {
				sign = "-"
				v = -v
			}
			s := fmt.Sprintf("%d", v)
			if len(s) <= 3 {
				return sign + s
			}
			var parts []string
			for len(s) > 3 {
				parts = append([]string{s[len(s)-3:]}, parts...)
				s = s[:len(s)-3]
			}
			parts = append([]string{s}, parts...)
			return sign + strings.Join(parts, ",")
		},
	}

	if r.cfg.LiveTemplate {
		path := filepath.Join(r.cfg.TemplatesDir, "site.html")
		return template.New("site.html").Funcs(funcs).ParseFiles(path)
	}

	data, err := fs.ReadFile(embeddedTemplates, "templates/site.html")
	if err != nil {
		return nil, err
	}
	return template.New("site.html").Funcs(funcs).Parse(string(data))
}

func (r *Renderer) loadCSS() ([]byte, error) {
	if r.cfg.LiveTemplate {
		return os.ReadFile(filepath.Join(r.cfg.TemplatesDir, "site.css"))
	}
	return fs.ReadFile(embeddedTemplates, "templates/site.css")
}

func (r *Renderer) loadSite() (*siteData, error) {
	users, err := loadUsers(r.db, r.cfg.TablePrefix)
	if err != nil {
		return nil, err
	}
	forumsByID, forums, err := loadForums(r.db, r.cfg.TablePrefix, r.cfg.TextFormats, r.cfg.Debug.HighlightRawHTML)
	if err != nil {
		return nil, err
	}
	attachmentsByPost, err := loadAttachments(r.db, r.cfg.TablePrefix)
	if err != nil {
		return nil, err
	}
	threads, err := loadThreads(r.db, r.cfg.TablePrefix, r.cfg.TextFormats, r.cfg.Debug.HighlightRawHTML, forumsByID, users, attachmentsByPost)
	if err != nil {
		return nil, err
	}
	stats := computeBoardStats(users, threads)

	for _, currentThread := range threads {
		currentForum := forumsByID[currentThread.ForumID]
		if currentForum != nil {
			currentForum.Threads = append(currentForum.Threads, currentThread)
		}
	}

	for _, currentForum := range allForums(forums) {
		sort.Slice(currentForum.Threads, func(i, j int) bool {
			left, right := currentForum.Threads[i], currentForum.Threads[j]
			if left.Sticky != right.Sticky {
				return left.Sticky
			}
			return left.LastPostAt.After(right.LastPostAt)
		})
		sort.Slice(currentForum.Children, func(i, j int) bool {
			return currentForum.Children[i].DisplayOrder < currentForum.Children[j].DisplayOrder
		})
	}
	sort.Slice(forums, func(i, j int) bool {
		return forums[i].DisplayOrder < forums[j].DisplayOrder
	})

	return &siteData{
		Title:       r.cfg.SiteTitle,
		BaseURL:     r.cfg.BaseURL,
		GeneratedAt: time.Now().UTC(),
		Forums:      forums,
		Threads:     threads,
		Stats:       stats,
	}, nil
}

func detectPrefix(db *sql.DB) (string, error) {
	const query = `
SELECT substr(name, 1, length(name) - length('forums'))
FROM sqlite_master
WHERE type = 'table' AND name LIKE '%forums'
ORDER BY name
LIMIT 1`
	var prefix string
	if err := db.QueryRow(query).Scan(&prefix); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("could not detect table prefix; pass -table-prefix explicitly")
		}
		return "", fmt.Errorf("detect table prefix: %w", err)
	}
	return prefix, nil
}

func loadUsers(db *sql.DB, prefix string) (map[int64]*user, error) {
	query := fmt.Sprintf(`SELECT uid, username, usertitle, avatar, postnum FROM %s`, tableName(prefix, "users"))
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("load users: %w", err)
	}
	defer rows.Close()

	users := map[int64]*user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Title, &u.AvatarPath, &u.PostCount); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		u.AvatarPath = normalizeAssetPath(u.AvatarPath)
		users[u.ID] = &u
	}
	return users, rows.Err()
}

func loadForums(db *sql.DB, prefix string, textFormats setupjson.TextFormats, highlightRawHTML bool) (map[int64]*forum, []*forum, error) {
	query := fmt.Sprintf(`
SELECT fid, pid, name, description, type, disporder, threads, posts, lastpost, lastpostsubject, lastposttid, lastposter
FROM %s
WHERE active = 1
ORDER BY pid, disporder, fid`, tableName(prefix, "forums"))
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("load forums: %w", err)
	}
	defer rows.Close()

	forumsByID := map[int64]*forum{}
	var ordered []*forum
	for rows.Next() {
		var (
			f                forum
			description      string
			lastPostUnix     int64
			lastPostSubject  string
			lastPostThreadID int64
			lastPoster       string
		)
		if err := rows.Scan(&f.ID, &f.ParentID, &f.Name, &description, &f.Type, &f.DisplayOrder, &f.ThreadsCount, &f.PostsCount, &lastPostUnix, &lastPostSubject, &lastPostThreadID, &lastPoster); err != nil {
			return nil, nil, fmt.Errorf("scan forum: %w", err)
		}
		format := textFormats.Format(configTableName(prefix, "forums"), "description")
		f.DescriptionHTML = renderFormattedText(format, description)
		f.DescriptionClass = contentDebugClass(format, highlightRawHTML)
		f.LastPost = unixTime(lastPostUnix)
		f.LastPostSubject = lastPostSubject
		f.LastPostThread = lastPostThreadID
		f.LastPostAuthor = lastPoster
		forumsByID[f.ID] = &f
		ordered = append(ordered, &f)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	var roots []*forum
	for _, currentForum := range ordered {
		if currentForum.ParentID == 0 {
			roots = append(roots, currentForum)
			continue
		}
		parent := forumsByID[currentForum.ParentID]
		if parent == nil {
			roots = append(roots, currentForum)
			continue
		}
		currentForum.Parent = parent
		parent.Children = append(parent.Children, currentForum)
	}
	return forumsByID, roots, nil
}

func loadAttachments(db *sql.DB, prefix string) (map[int64][]*attachment, error) {
	query := fmt.Sprintf(`
SELECT aid, pid, filename, attachname, thumbnail, filetype, filesize
FROM %s
WHERE visible = 1
ORDER BY aid`, tableName(prefix, "attachments"))
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("load attachments: %w", err)
	}
	defer rows.Close()

	attachmentsByPost := map[int64][]*attachment{}
	for rows.Next() {
		var a attachment
		if err := rows.Scan(&a.ID, &a.PostID, &a.FileName, &a.StoredPath, &a.ThumbnailPath, &a.FileType, &a.FileSize); err != nil {
			return nil, fmt.Errorf("scan attachment: %w", err)
		}
		a.StoredPath = normalizeUploadAttachmentPath(a.StoredPath)
		a.ThumbnailPath = normalizeUploadAttachmentPath(a.ThumbnailPath)
		attachmentsByPost[a.PostID] = append(attachmentsByPost[a.PostID], &a)
	}
	return attachmentsByPost, rows.Err()
}

func loadThreads(db *sql.DB, prefix string, textFormats setupjson.TextFormats, highlightRawHTML bool, forumsByID map[int64]*forum, users map[int64]*user, attachmentsByPost map[int64][]*attachment) ([]*thread, error) {
	threadQuery := fmt.Sprintf(`
SELECT tid, fid, subject, uid, username, dateline, lastpost, replies, views, sticky
FROM %s
WHERE visible = 1
ORDER BY dateline`, tableName(prefix, "threads"))
	rows, err := db.Query(threadQuery)
	if err != nil {
		return nil, fmt.Errorf("load threads: %w", err)
	}
	defer rows.Close()

	threadByID := map[int64]*thread{}
	var threads []*thread
	for rows.Next() {
		var t thread
		var createdAtUnix, lastPostUnix int64
		var sticky int
		if err := rows.Scan(&t.ID, &t.ForumID, &t.Subject, &t.AuthorID, &t.AuthorName, &createdAtUnix, &lastPostUnix, &t.Replies, &t.Views, &sticky); err != nil {
			return nil, fmt.Errorf("scan thread: %w", err)
		}
		t.CreatedAt = unixTime(createdAtUnix)
		t.LastPostAt = unixTime(lastPostUnix)
		t.Sticky = sticky == 1
		t.Forum = forumsByID[t.ForumID]
		threadByID[t.ID] = &t
		threads = append(threads, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	postQuery := fmt.Sprintf(`
SELECT pid, tid, uid, username, dateline, edittime, editreason, message
FROM %s
WHERE visible = 1
ORDER BY tid, dateline, pid`, tableName(prefix, "posts"))
	postRows, err := db.Query(postQuery)
	if err != nil {
		return nil, fmt.Errorf("load posts: %w", err)
	}
	defer postRows.Close()

	for postRows.Next() {
		var p post
		var createdAtUnix, editedAtUnix int64
		var message string
		if err := postRows.Scan(&p.ID, &p.ThreadID, &p.AuthorID, &p.AuthorName, &createdAtUnix, &editedAtUnix, &p.EditReason, &message); err != nil {
			return nil, fmt.Errorf("scan post: %w", err)
		}
		p.CreatedAt = unixTime(createdAtUnix)
		p.EditedAt = unixTime(editedAtUnix)
		format := textFormats.Format(configTableName(prefix, "posts"), "message")
		p.BodyHTML = renderFormattedText(format, message)
		p.BodyClass = contentDebugClass(format, highlightRawHTML)
		p.Attachments = attachmentsByPost[p.ID]
		p.Author = users[p.AuthorID]
		if !p.EditedAt.IsZero() {
			p.EditedAtString = p.EditedAt.Format("Jan 2, 2006 15:04 MST")
		}
		currentThread := threadByID[p.ThreadID]
		if currentThread != nil {
			currentThread.Posts = append(currentThread.Posts, &p)
		}
	}
	if err := postRows.Err(); err != nil {
		return nil, err
	}
	return threads, nil
}

func renderPage(tmpl *template.Template, path string, data pageData) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("render %s: %w", path, err)
	}
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func renderHTTPErrorPage(w http.ResponseWriter, tmpl *template.Template, data pageData) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(data.ErrorPage.StatusCode)
	w.Write(buf.Bytes())
}

func errorPages(site *siteData) []pageData {
	return []pageData{
		customErrorPageData(site, http.StatusNotFound),
		customErrorPageData(site, http.StatusInternalServerError),
	}
}

func customErrorPageData(site *siteData, statusCode int) pageData {
	page := errorPage{
		StatusCode: statusCode,
	}
	switch statusCode {
	case http.StatusNotFound:
		page.Heading = "Page Not Found"
		page.Message = "The archive could not find the page you requested. The thread, forum, or file may not exist in this snapshot."
	case http.StatusInternalServerError:
		page.Heading = "Archive Error"
		page.Message = "Something went wrong while serving this page. Please try again, or return to the archive index."
	default:
		page.Heading = "Archive Notice"
		page.Message = "This archive page is not available right now."
	}

	return pageData{
		Site:      site,
		PageTitle: fmt.Sprintf("%d %s - %s", page.StatusCode, page.Heading, site.Title),
		Section:   "error",
		ErrorPage: &page,
	}
}

func allForums(roots []*forum) []*forum {
	var out []*forum
	var walk func(items []*forum)
	walk = func(items []*forum) {
		for _, current := range items {
			out = append(out, current)
			if len(current.Children) > 0 {
				walk(current.Children)
			}
		}
	}
	walk(roots)
	return out
}

func computeBoardStats(users map[int64]*user, threads []*thread) boardStats {
	stats := boardStats{
		ThreadsTotal: int64(len(threads)),
		MembersTotal: int64(len(users)),
	}
	var newest *user
	for _, currentUser := range users {
		if newest == nil || currentUser.ID > newest.ID {
			newest = currentUser
		}
	}
	stats.NewestMember = newest
	for _, currentThread := range threads {
		stats.PostsTotal += int64(len(currentThread.Posts))
	}
	return stats
}

func findForum(roots []*forum, id int64) *forum {
	for _, current := range allForums(roots) {
		if current.ID == id {
			return current
		}
	}
	return nil
}

func findThread(threads []*thread, id int64) *thread {
	for _, current := range threads {
		if current.ID == id {
			return current
		}
	}
	return nil
}

func breadcrumbsForForum(site *siteData, f *forum) []crumb {
	var chain []*forum
	for current := f; current != nil; current = current.Parent {
		chain = append(chain, current)
	}
	reverseForums(chain)
	crumbs := []crumb{{Label: "Forums", Href: relLink(site.BaseURL, "/")}}
	for _, current := range chain {
		crumbs = append(crumbs, crumb{
			Label: current.Name,
			Href:  relLink(site.BaseURL, forumHref(current.ID)),
		})
	}
	return crumbs
}

func breadcrumbsForThread(site *siteData, t *thread) []crumb {
	crumbs := breadcrumbsForForum(site, t.Forum)
	return append(crumbs, crumb{Label: t.Subject, Href: relLink(site.BaseURL, threadHref(t.ID))})
}

func reverseForums(items []*forum) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

func parseID(path, prefix string) (int64, bool) {
	value := strings.TrimPrefix(path, prefix)
	if value == "" || strings.Contains(value, "/") {
		return 0, false
	}
	var id int64
	_, err := fmt.Sscanf(value, "%d", &id)
	return id, err == nil
}

func unixTime(v int64) time.Time {
	if v <= 0 {
		return time.Time{}
	}
	return time.Unix(v, 0).UTC()
}

func tableName(prefix, base string) string {
	return `"` + strings.ReplaceAll(configTableName(prefix, base), `"`, `""`) + `"`
}

func configTableName(prefix, base string) string {
	return prefix + base
}

func normalizeAssetPath(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if idx := strings.Index(value, "?"); idx >= 0 {
		value = value[:idx]
	}
	value = strings.TrimPrefix(value, ".")
	if !strings.HasPrefix(value, "/") {
		value = "/" + value
	}
	return value
}

func normalizeUploadAttachmentPath(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if !strings.HasPrefix(value, "/") {
		value = "/uploads/" + strings.TrimPrefix(value, "/")
	}
	return value
}

func normalizeBaseURL(base string) string {
	base = strings.TrimSpace(base)
	if base == "" {
		return "/"
	}
	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}
	return strings.TrimRight(base, "/") + "/"
}

func relLink(baseURL, path string) string {
	path = strings.TrimPrefix(path, "/")
	return normalizeBaseURL(baseURL) + path
}

func forumHref(fid int64) string {
	return fmt.Sprintf("/f/%d/", fid)
}

func threadHref(tid int64) string {
	return fmt.Sprintf("/t/%d/", tid)
}

func renderBBCode(source string) template.HTML {
	if strings.TrimSpace(source) == "" {
		return ""
	}
	escaped := html.EscapeString(strings.ReplaceAll(source, "\r\n", "\n"))
	replacements := []struct {
		pattern *regexp.Regexp
		replace string
	}{
		{regexp.MustCompile(`(?is)\[b\](.*?)\[/b\]`), "<strong>$1</strong>"},
		{regexp.MustCompile(`(?is)\[i\](.*?)\[/i\]`), "<em>$1</em>"},
		{regexp.MustCompile(`(?is)\[u\](.*?)\[/u\]`), "<span class=\"bb-underline\">$1</span>"},
		{regexp.MustCompile(`(?is)\[s\](.*?)\[/s\]`), "<span class=\"bb-strike\">$1</span>"},
		{regexp.MustCompile(`(?is)\[code\](.*?)\[/code\]`), "<pre class=\"bb-code\"><code>$1</code></pre>"},
		{regexp.MustCompile(`(?is)\[quote='([^']+)'.*?\](.*?)\[/quote\]`), "<blockquote><header>$1 wrote</header><p>$2</p></blockquote>"},
		{regexp.MustCompile(`(?is)\[quote=&quot;([^"]+)&quot;.*?\](.*?)\[/quote\]`), "<blockquote><header>$1 wrote</header><p>$2</p></blockquote>"},
		{regexp.MustCompile(`(?is)\[quote\](.*?)\[/quote\]`), "<blockquote><p>$1</p></blockquote>"},
		{regexp.MustCompile(`(?is)\[url=(https?://[^\]]+)\](.*?)\[/url\]`), "<a href=\"$1\">$2</a>"},
		{regexp.MustCompile(`(?is)\[url\](https?://.*?)\[/url\]`), "<a href=\"$1\">$1</a>"},
		{regexp.MustCompile(`(?is)\[img\](https?://.*?)\[/img\]`), "<img src=\"$1\" alt=\"embedded image\" loading=\"lazy\">"},
		{regexp.MustCompile(`(?is)\[email=([^\]]+)\](.*?)\[/email\]`), "<a href=\"mailto:$1\">$2</a>"},
		{regexp.MustCompile(`(?is)\[email\](.*?)\[/email\]`), "<a href=\"mailto:$1\">$1</a>"},
		{regexp.MustCompile(`(?is)\[size=[^\]]+\](.*?)\[/size\]`), "<span>$1</span>"},
		{regexp.MustCompile(`(?is)\[color=[^\]]+\](.*?)\[/color\]`), "<span>$1</span>"},
		{regexp.MustCompile(`(?is)\[align=[^\]]+\](.*?)\[/align\]`), "<div class=\"bb-align\">$1</div>"},
	}
	for _, replacement := range replacements {
		escaped = replacement.pattern.ReplaceAllString(escaped, replacement.replace)
	}
	escaped = renderLists(escaped)
	escaped = regexp.MustCompile(`(?is)\[/?(?:font|hr)\b[^\]]*\]`).ReplaceAllString(escaped, "")
	escaped = regexp.MustCompile(`(?is)\[(?:left|center|right)\](.*?)\[/(?:left|center|right)\]`).ReplaceAllString(escaped, "<div class=\"bb-align\">$1</div>")
	escaped = regexp.MustCompile(`(?is)\[attach\][^\[]*?\[/attach\]`).ReplaceAllString(escaped, "")
	escaped = regexp.MustCompile(`(?is)\[video=.*?\].*?\[/video\]`).ReplaceAllString(escaped, "<p><em>Embedded video omitted in static export.</em></p>")

	paragraphs := strings.Split(escaped, "\n\n")
	for i, paragraph := range paragraphs {
		trimmed := strings.TrimSpace(paragraph)
		if strings.Contains(paragraph, "<pre ") || strings.HasPrefix(trimmed, "<blockquote") || strings.HasPrefix(trimmed, "<ul") || strings.HasPrefix(trimmed, "<ol") || strings.HasPrefix(trimmed, "<div ") {
			paragraphs[i] = strings.ReplaceAll(paragraph, "\n", "<br>\n")
			continue
		}
		paragraphs[i] = "<p>" + strings.ReplaceAll(paragraph, "\n", "<br>\n") + "</p>"
	}
	rendered := strings.Join(paragraphs, "\n")
	rendered = strings.ReplaceAll(rendered, "<p></p>", "")
	return template.HTML(rendered)
}

func renderFormattedText(format setupjson.TextFormat, source string) template.HTML {
	switch format {
	case setupjson.TextFormatRawHTML:
		return sanitizeHTMLFragment(source)
	case setupjson.TextFormatBBCodes:
		return renderBBCode(source)
	case setupjson.TextFormatTaintedText:
		fallthrough
	default:
		return renderPlainText(source)
	}
}

func contentDebugClass(format setupjson.TextFormat, highlightRawHTML bool) string {
	if highlightRawHTML && format == setupjson.TextFormatRawHTML {
		return " debug-raw-html"
	}
	return ""
}

func renderPlainText(source string) template.HTML {
	if strings.TrimSpace(source) == "" {
		return ""
	}
	escaped := html.EscapeString(strings.ReplaceAll(html.UnescapeString(source), "\r\n", "\n"))
	paragraphs := strings.Split(escaped, "\n\n")
	for i, paragraph := range paragraphs {
		paragraphs[i] = "<p>" + strings.ReplaceAll(paragraph, "\n", "<br>\n") + "</p>"
	}
	rendered := strings.Join(paragraphs, "\n")
	rendered = strings.ReplaceAll(rendered, "<p></p>", "")
	return template.HTML(rendered)
}

func sanitizeHTMLFragment(source string) template.HTML {
	if strings.TrimSpace(source) == "" {
		return ""
	}
	return template.HTML(rawHTMLPolicy.Sanitize(source))
}

func newRawHTMLPolicy() *bluemonday.Policy {
	policy := bluemonday.UGCPolicy()
	policy.AllowImages()
	policy.RequireNoFollowOnFullyQualifiedLinks(true)
	policy.RequireNoReferrerOnFullyQualifiedLinks(true)
	policy.AddTargetBlankToFullyQualifiedLinks(true)
	return policy
}

func renderLists(source string) string {
	listRE := regexp.MustCompile(`(?is)\[list(?:=1)?\](.*?)\[/list\]`)
	return listRE.ReplaceAllStringFunc(source, func(match string) string {
		ordered := strings.HasPrefix(strings.ToLower(match), "[list=1]")
		content := regexp.MustCompile(`(?is)^\[list(?:=1)?\]`).ReplaceAllString(match, "")
		content = regexp.MustCompile(`(?is)\[/list\]$`).ReplaceAllString(content, "")
		rawItems := strings.Split(content, "[*]")
		if len(rawItems) <= 1 {
			return content
		}
		tag := "ul"
		if ordered {
			tag = "ol"
		}
		var b strings.Builder
		b.WriteString("<" + tag + ">")
		for _, item := range rawItems[1:] {
			item = strings.TrimSpace(item)
			if item == "" {
				continue
			}
			b.WriteString("<li>")
			b.WriteString(item)
			b.WriteString("</li>")
		}
		b.WriteString("</" + tag + ">")
		return b.String()
	})
}

func (r *Renderer) CopyAsset(name string, w io.Writer) error {
	data, err := r.loadCSS()
	if err != nil {
		return err
	}
	if name != "site.css" {
		return fmt.Errorf("unknown asset %q", name)
	}
	_, err = w.Write(data)
	return err
}
