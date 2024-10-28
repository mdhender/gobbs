// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

// Package config_test implementes tests against the exposed Config API
//
// Apparently https://github.com/golang/go/issues/31859#issuecomment-489889428
// kind of broke tests that need command line flags. Limiting test to just the
// default values.
package config_test

import (
	"github.com/mdhender/gobbs/internal/config"
	"testing"
)

func TestDefault(t *testing.T) {
	// Specification: Config API

	// Given an empty environment
	// When we create a new configuration
	cfg := config.Default()
	// Then we should have a valid Config
	if cfg == nil {
		t.Fatalf("default: expected config to be non-nil: got nil\n")
	}
	// And default values should be ...
	if expected := false; cfg.Debug != expected {
		t.Errorf("default: expected debug to be %v: got %v\n", expected, cfg.Debug)
	}
	if expected := "http"; cfg.Server.Scheme != expected {
		t.Errorf("default: expected server.scheme to be %q: got %q\n", expected, cfg.Server.Scheme)
	}
	if expected := "localhost"; cfg.Server.Host != expected {
		t.Errorf("default: expected server.host to be %q: got %q\n", expected, cfg.Server.Host)
	}
	if expected := "3000"; cfg.Server.Port != expected {
		t.Errorf("default: expected server.port to be %q: got %q\n", expected, cfg.Server.Port)
	}
}
