package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mdhender/gobbs/internal/forumsite"
)

func main() {
	var cfg forumsite.Config
	var listenAddr string

	flag.StringVar(&cfg.SQLitePath, "sqlite-path", "mybb.sqlite3", "path to the SQLite archive")
	flag.StringVar(&cfg.SetupPath, "setup-file", "setup.json", "path to setup.json")
	flag.StringVar(&cfg.TablePrefix, "table-prefix", "", "MyBB table prefix, auto-detected when empty")
	flag.StringVar(&cfg.SiteTitle, "site-title", "PlayByMail Forums Archive", "site title for generated pages")
	flag.StringVar(&cfg.BaseURL, "base-url", "/", "base URL prefix for generated links")
	flag.StringVar(&cfg.UploadsDir, "uploads-dir", "uploads", "directory containing archived uploads and avatars")
	flag.StringVar(&cfg.TemplatesDir, "templates-dir", "internal/forumsite/templates", "directory containing live templates and assets")
	flag.StringVar(&listenAddr, "listen", "127.0.0.1:8080", "address for the preview server")
	flag.Parse()

	cfg.LiveTemplate = true

	renderer, err := forumsite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Close()

	log.Printf("preview server listening on http://%s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, renderer.Handler()))
}
