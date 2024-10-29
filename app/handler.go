// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package app

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// getIndex serves the index page.
func getIndex(components string) func(http.ResponseWriter, *http.Request) {
	files := []string{
		filepath.Join(components, "index.gohtml"),
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s: entered\n", r.Method, r.URL.Path)

		var fragments [][]byte

		var payload string
		if frag, err := renderFragment(payload, "index.php", files...); err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL.Path, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			fragments = append(fragments, frag)
		}

		if _, err := writeFragments(w, r, fragments...); err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL.Path, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func getShowthread(components string) func(http.ResponseWriter, *http.Request) {
	files := []string{
		filepath.Join(components, "showthread.gohtml"),
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s: entered\n", r.Method, r.URL.Path)
		action, pid, tid := r.URL.Query().Get("action"), r.URL.Query().Get("pid"), r.URL.Query().Get("tid")
		if action == "lastpost" && tid == "130855" {
			http.Redirect(w, r, "/showthread.php?tid=130855&pid=138816#pid138816", http.StatusSeeOther)
			return
		}
		if !(tid == "130855" && pid == "138816") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("%s %s: entered\n", r.Method, r.URL.Path)

		var fragments [][]byte

		var payload string
		if frag, err := renderFragment(payload, "showthread.php", files...); err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL.Path, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			fragments = append(fragments, frag)
		}

		if _, err := writeFragments(w, r, fragments...); err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL.Path, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// getTasks was used by MyBB to initiate background tasks.
// it is currently a no-op.
func getTasks() func(http.ResponseWriter, *http.Request) {
	log.Printf("getTasks: not implemented\n")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<p>ok</p>"))
	}
}

func serveAdminShutdownServer(key string, stop chan os.Signal) func(http.ResponseWriter, *http.Request) {
	keyHash := sha256.Sum256([]byte(key))
	return func(w http.ResponseWriter, r *http.Request) {
		argHash := sha256.Sum256([]byte(r.PathValue("key")))
		if !bytes.Equal(keyHash[:], argHash[:]) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		log.Printf("%s %s: initiating shutdown\n", r.Method, r.URL.Path)

		stop <- syscall.SIGTERM

		// important: calling http.Error() will close the connection and allow us to gracefully shut down the server.
		http.Error(w, "By your command, server is shutting down", http.StatusServiceUnavailable)
	}
}

// serveAssets returns a handler that uses http.ServeContent to serve files in the assets directory.
// As we all know, Go treats the "/" path as a wild-card. If we see it here, we smile, nod, and send
// the client off to fetch the index page.
func serveAssets(assets string, debugAssets bool, indexPage string) http.HandlerFunc {
	if ok, _ := isDirExists(assets); !ok {
		log.Printf("assets %q: is not a directory\n", assets)
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		} else if r.URL.Path == "/" {
			http.Redirect(w, r, indexPage, http.StatusSeeOther)
		}

		file := filepath.Join(assets, filepath.Clean(strings.TrimPrefix(r.URL.Path, "")))
		if debugAssets {
			log.Printf("%s: %s: assets\n", r.Method, r.URL.Path)
		}

		// only serve regular files, never directories or directory listings.
		sb, err := os.Stat(file)
		if err != nil || sb.IsDir() || !sb.Mode().IsRegular() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// create a reader because ServeContent() requires one.
		rdr, err := os.Open(file)
		if err != nil {
			log.Printf("%s %s: file disappeared: %v\n", r.Method, r.URL.Path, err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		defer func(r io.ReadCloser) {
			_ = r.Close()
		}(rdr)

		// let Go serve the file. it does magic things like content-type, etc.
		http.ServeContent(w, r, file, sb.ModTime(), rdr)
	}
}

// servePage serves a page from the file system.
// path must be the full path to the page to serve or relative to the server's working directory.
func servePage(assets, page string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s: page %v\n", r.Method, r.URL.Path, page)

		// only serve files, never directories
		sb, err := os.Stat(filepath.Join(assets, page))
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Printf("%s %s: %v\n", r.Method, r.URL.Path, err)
			}
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if sb.IsDir() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("%s %s: mod %v\n", r.Method, r.URL.Path, sb.ModTime())

		// we have to open the file ourselves, because ServeContent doesn't support serving from a file
		fp, err := os.Open(filepath.Join(assets, page))
		if err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL.Path, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// and we must close the file ourselves or leak the file descriptor
		defer fp.Close()

		// let ServeContent do the rest, including setting the Content-Type header
		http.ServeContent(w, r, r.URL.Path, sb.ModTime(), fp)
	}
}
