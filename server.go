// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mdhender/gobbs/internal/domains"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func newServer(options ...Option) (*Server, error) {
	s := &Server{
		scheme: "http",
	}
	s.Addr = net.JoinHostPort(s.host, s.port)
	s.MaxHeaderBytes = 1 << 20
	s.IdleTimeout = 10 * time.Second
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 10 * time.Second

	// create a channel to listen for OS signals.
	s.admin.keys.shutdown = uuid.NewString()
	s.admin.stop = make(chan os.Signal, 1)
	signal.Notify(s.admin.stop, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	for _, option := range options {
		if err := option(s); err != nil {
			return nil, err
		}
	}

	s.Handler = s.Routes()

	return s, nil
}

type Server struct {
	http.Server
	scheme, host, port string
	admin              struct {
		ctx  context.Context
		stop chan os.Signal // channel to stop the server
		keys struct {
			shutdown string // key to stop the server
		}
	}
	paths struct {
		assets     string
		components string
		database   string
	}
}

func (s *Server) BaseURL() string {
	return fmt.Sprintf("%s://%s", s.scheme, s.Addr)
}

func (s *Server) PrintAdminRoutes() {
	log.Printf("shutdown server: %s/admin/shutdown-server/%s\n", s.BaseURL(), s.admin.keys.shutdown)
}

func (s *Server) Serve() error {
	// start the server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("listening on %s\n", s.BaseURL())
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server: %v\n", err)
		}
		log.Printf("server: shutdown and closed\n")
	}()

	// server is running; block until we receive a signal.
	sig := <-s.admin.stop

	started := time.Now()
	log.Printf("signal: received %v (%v)\n", sig, time.Since(started))

	// graceful shutdown with a timeout.
	timeout := time.Second * 10
	log.Printf("creating context with %v timeout (%v)\n", timeout, time.Since(started))
	ctx, cancel := context.WithTimeout(s.admin.ctx, timeout)
	defer cancel()

	// cancel any idle connections.
	log.Printf("canceling idle connections (%v)\n", time.Since(started))
	s.SetKeepAlivesEnabled(false)

	log.Printf("sending signal to shut down the server (%v)\n", time.Since(started))
	if err := s.Shutdown(ctx); err != nil {
		return errors.Join(domains.ErrServerShutdown, err)
	}

	log.Printf("server stopped ¡gracefully! (%v)\n", time.Since(started))
	return nil
}

type Options []Option
type Option func(*Server) error

func withAssets(path string) Option {
	return func(s *Server) error {
		if path == "" {
			return fmt.Errorf("path is required\n")
		} else if ok, err := isDirExists(path); err != nil {
			return err
		} else if !ok {
			return domains.ErrNotDirectory
		} else if s.paths.assets, err = filepath.Abs(path); err != nil {
			return err
		}
		return nil
	}
}

func withComponents(path string) Option {
	return func(s *Server) error {
		if path == "" {
			return fmt.Errorf("path is required\n")
		} else if ok, err := isDirExists(path); err != nil {
			return err
		} else if !ok {
			return domains.ErrNotDirectory
		} else if s.paths.components, err = filepath.Abs(path); err != nil {
			return err
		}
		return nil
	}
}

func withContext(ctx context.Context) Option {
	return func(s *Server) error {
		s.admin.ctx = ctx
		return nil
	}
}

func withDatabase(path string) Option {
	return func(s *Server) error {
		if path == "" {
			return fmt.Errorf("path is required\n")
		} else if ok, err := isFileExists(path); err != nil {
			return err
		} else if !ok {
			return domains.ErrNotFile
		} else if s.paths.database, err = filepath.Abs(path); err != nil {
			return err
		}
		return nil
	}
}

func withHost(host string) Option {
	return func(s *Server) error {
		s.host = host
		s.Addr = net.JoinHostPort(s.host, s.port)
		return nil
	}
}

func withPort(port string) Option {
	return func(s *Server) error {
		s.port = port
		s.Addr = net.JoinHostPort(s.host, s.port)
		return nil
	}
}
