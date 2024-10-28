// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"log"
	"net/http"
	"os"
	"syscall"
)

func (s *Server) serveAdminShutdownServer(key string) func(http.ResponseWriter, *http.Request) {
	keyHash := sha256.Sum256([]byte(key))
	log.Printf("server: %s/admin/shutdown-server/%s", s.BaseURL(), key)
	return func(w http.ResponseWriter, r *http.Request) {
		argHash := sha256.Sum256([]byte(r.PathValue("key")))
		if !bytes.Equal(keyHash[:], argHash[:]) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		log.Printf("%s %s: initiating shutdown\n", r.Method, r.URL.Path)

		s.admin.stop <- syscall.SIGTERM

		// important: calling http.Error() will close the connection and allow us to gracefully shut down the server.
		http.Error(w, "By your command, server is shutting down", http.StatusServiceUnavailable)
	}
}

// servePage serves a page from the file system.
// path must be the full path to the page to serve or relative to the server's working directory.
func (s *Server) servePage(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s: page %v\n", r.Method, r.URL.Path, path)

		// only serve files, never directories
		sb, err := os.Stat(path)
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
		fp, err := os.Open(path)
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
