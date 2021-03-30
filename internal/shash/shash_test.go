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
