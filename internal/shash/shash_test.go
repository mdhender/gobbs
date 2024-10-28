// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

package shash_test

import (
	"github.com/mdhender/gobbs/internal/shash"
	"testing"
)

func TestNew(t *testing.T) {
	// Specification: Insecure Password Hashing API

	// Given a salt
	// When I create a new hashing function
	// Then something
	h1 := shash.New([]byte("saltsalt"), 0)([]byte("secret"))
	h2 := shash.New([]byte("saltsalt"), 0)([]byte("secret"))
	h3 := shash.New([]byte("saltsalt"), 9)([]byte("secret"))
	h4 := shash.New([]byte("saltsalt"), 10)([]byte("secret"))
	h5 := shash.New([]byte("saltsalt"), 9)([]byte("secres"))
	h6 := shash.New([]byte("saltsals"), 10)([]byte("secret"))
	if h1 != h2 {
		t.Errorf("shash: h1<->h2 expected %q: got %q\n", h1, h2)
	}
	if h1 == h3 {
		t.Errorf("shash: h1<->h3 expected different hash: got %q\n", h3)
	}
	if h1 == h4 {
		t.Errorf("shash: h1<->h4 expected different hash: got %q\n", h4)
	}
	if h1 == h5 {
		t.Errorf("shash: h1<->h5 expected different hash: got %q\n", h5)
	}
	if h3 == h4 {
		t.Errorf("shash: h3<->h4 expected different hash: got %q\n", h4)
	}
	if h3 == h5 {
		t.Errorf("shash: h3<->h5 expected different hash: got %q\n", h5)
	}
	if h3 == h6 {
		t.Errorf("shash: h3<->h6 expected different hash: got %q\n", h6)
	}
}
