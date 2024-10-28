// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

// Package server implements an HTTP server for GoBBS
package server

import (
	"fmt"
	"github.com/mdhender/gobbs/internal/config"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) (*Server, error) {
	return &Server{cfg: cfg}, fmt.Errorf("!implemented")
}
