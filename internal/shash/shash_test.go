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
	h1, h2 := shash.New([]byte("salt"))([]byte("secret")), shash.New([]byte("salt"))([]byte("secret"))
	if h1 != h2 {
		t.Errorf("shash: expected %q: got %q\n", h1, h2)
	}
	h3 := shash.New([]byte("sals"))([]byte("secret"))
	if h1 == h3 {
		t.Errorf("shash: expected different hash: got %q\n", h3)
	}
	h4 := shash.New([]byte("salt"))([]byte("secres"))
	if h1 == h4 {
		t.Errorf("shash: expected different hash: got %q\n", h4)
	}
}
