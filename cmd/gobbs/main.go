/*
 * gobbs - threaded forum server
 *
 * Copyright (c) 2021 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package main implements a GoBBS server.
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mdhender/gobbs/internal/config"
	"github.com/mdhender/gobbs/internal/dot"
	"github.com/mdhender/gobbs/internal/mybb"
	"log"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	defer func(started time.Time) {
		log.Printf("[main] elapsed time %v\n", time.Now().Sub(started))
	}(time.Now())
	log.Println("[main] entered")

	if err := dot.Load("GOBBS", false, false); err != nil {
		log.Fatalf("main: %+v\n", err)
	}

	cfg := config.Default()
	if err := cfg.Load(); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(2)
	}

	if err := run(cfg); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(2)
	}
}

func run(cfg *config.Config) error {
	// set up database connection
	mb := mydb{
		ctx: context.Background(),
	}
	db, err := sql.Open("mysql", cfg.DB.DSN())
	if err != nil {
		return err
	}
	mb.queries = mybb.New(db)

	forums, err := mb.fetchForums()
	if err != nil {
		return fmt.Errorf("fetchForums: %w", err)
	}

	// filter the forums
	var tmp Forums
	for _, forum := range forums {
		if forum.Name == "Galac-Tac" {
			//tmp = append(tmp, forum)
		} else if forum.Name == "Midgard" {
			tmp = append(tmp, forum)
		}
	}
	forums = tmp

	// fetch threads and posts for the forums
	for _, forum := range forums {
		if err := mb.fetchForum(forum); err != nil {
			return err
		}
	}

	// dump the forums
	if buf, err := json.MarshalIndent(forums, "", "  "); err != nil {
		return err
	} else if err = os.WriteFile("forums.json", buf, 0644); err != nil {
		return err
	}

	//if err := mb.listAllTemplates(); err != nil {
	//	return fmt.Errorf("listAllTemplates: %w", err)
	//}

	//// dump the Galac-Tac forum
	//forumId := 28
	//if err := mb.fetchThreadsForForum(forumId); err != nil {
	//	return fmt.Errorf("fetchThreadsForForum: %d: %w", forumId, err)
	//}
	//if err := mb.fetchThreadUsersForForum(forumId); err != nil {
	//	return fmt.Errorf("fetchThreadsForForum: %d: %w", forumId, err)
	//}
	//threadId := 130481
	//if err := mb.fetchThread(threadId); err != nil {
	//	return fmt.Errorf("fetchThreadsForForum: %d: %w", forumId, err)
	//}

	//// dump the Midgard forum
	//forumId = 44
	//if err := mb.fetchThreadsForForum(forumId); err != nil {
	//	return fmt.Errorf("fetchThreadsForForum: %d: %w", forumId, err)
	//}
	//if err := mb.fetchThreadUsersForForum(forumId); err != nil {
	//	return fmt.Errorf("fetchThreadsForForum: %d: %w", forumId, err)
	//}

	//srv, err := server.New(cfg)
	//if err != nil {
	//	return err
	//} else if srv == nil {
	//	return fmt.Errorf("assert(srv != nil)")
	//}
	return nil
}

type mydb struct {
	ctx     context.Context
	queries *mybb.Queries
}

func (m mydb) fetchThread(threadId int64) error {
	started := time.Now()
	t, err := m.queries.GetThread(m.ctx, threadId)
	log.Printf("query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return err
	}
	log.Printf("thread %d\n", t.Tid)
	return nil
}

func (m mydb) fetchThreadsForForum(fid int64) error {
	started := time.Now()
	t, err := m.queries.GetThreadsForForum(m.ctx, fid)
	log.Printf("query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return err
	}
	log.Printf("forum %d threads %d\n", fid, t.Threads)
	return nil
}

func (m mydb) fetchThreadUsersForForum(fid int64) error {
	started := time.Now()
	t, err := m.queries.GetThreadUsers(m.ctx, fid)
	log.Printf("query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return err
	}
	log.Printf("forum %d threads %d\n", fid, len(t))
	for n, tt := range t {
		log.Printf("forum %d %d thread %d subj %s\n", fid, n, tt.Tid, tt.Subject)
	}
	return nil
}

func (m mydb) listAllTemplates() error {
	started := time.Now()
	tmps, err := m.queries.GetTemplates(m.ctx)
	log.Printf("query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return err
	}
	log.Printf("length(tmps) is %d\n", len(tmps))
	for n, t := range tmps {
		log.Printf("%3d: %-30s %8d\n", n, t.Title, len(t.Template))
	}
	return nil
}

type Forum struct {
	Id          int64
	Name        string
	Description string
	Threads     []*Thread
}

func (m mydb) fetchForum(forum *Forum) error {
	log.Printf("forums: fetching %q\n", forum.Name)
	threads, err := m.fetchThreads(forum.Id)
	if err != nil {
		return fmt.Errorf("forum %d: thread: %w", forum.Id, err)
	}
	forum.Threads = threads
	return nil
}

type Forums []*Forum

func (m mydb) fetchForums() (Forums, error) {
	started := time.Now()
	rows, err := m.queries.FetchForums(m.ctx)
	log.Printf("fetchForums: query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return nil, err
	}
	var forums Forums
	for _, row := range rows {
		forum := &Forum{
			Id:          row.Fid,
			Name:        row.Name,
			Description: row.Description,
		}
		forums = append(forums, forum)
	}
	return forums, nil
}

type Threads []*Thread

type Thread struct {
	Id      int64
	From    string
	Date    time.Time
	Subject string
	Posts   Posts
}

func (m mydb) fetchThreads(forumId int64) (Threads, error) {
	started := time.Now()
	rows, err := m.queries.FetchThreads(m.ctx, forumId)
	log.Printf("fetchThreads: query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return nil, err
	}
	var threads Threads
	for _, row := range rows {
		thread := &Thread{
			Id:      row.Tid,
			From:    row.Username,
			Date:    time.Unix(int64(row.Dateline), 0).UTC(),
			Subject: row.Subject,
		}
		if thread.Posts, err = m.fetchPosts(thread.Id); err != nil {
			return nil, fmt.Errorf("fetchThread: %w", err)
		}
		threads = append(threads, thread)
		log.Printf("forum %d thread %d subject %q\n", forumId, thread.Id, thread.Subject)
	}
	return threads, nil
}

type Posts []*Post

type Post struct {
	Id      int
	From    string
	Date    time.Time
	Subject string
	ReplyTo int
	Message string
}

func (m mydb) fetchPosts(threadId int64) (Posts, error) {
	started := time.Now()
	rows, err := m.queries.FetchPosts(m.ctx, threadId)
	log.Printf("fetchThreads: query ran %v\n", time.Now().Sub(started))
	if err != nil {
		return nil, err
	}
	var posts Posts
	for _, row := range rows {
		post := &Post{
			Id:      int(row.Pid),
			From:    row.Username,
			Date:    time.Unix(row.Dateline, 0).UTC(),
			Subject: row.Subject,
			ReplyTo: int(row.Replyto),
			Message: row.Message,
		}
		posts = append(posts, post)
		log.Printf("thread %d post %d subject %q\n", threadId, post.Id, post.Subject)
	}
	return posts, nil
}
