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

package memory_test

import (
	"github.com/mdhender/gobbs/internal/store/memory"
	"testing"
)

func TestAuthor(t *testing.T) {
	// Specification: Author

	// When a new author is created
	// Then it has a unique ID
	// And it has the given name
	for _, tc := range []struct {
		name string
	}{
		{"James Joyce"},
	} {
		ds, _ := memory.NewStore()
		id, _ := ds.CreateAuthor(tc.name)
		o, ok := ds.FindAuthorByID(id)
		if !ok {
			t.Errorf("author does not have unique ID: expected id %q: got no author found\n", id)
		} else if id != o.ID {
			t.Errorf("author does not have unique ID: expected id %q: got %q\n", id, o.ID)
		} else {
			if tc.name != o.Name {
				t.Errorf("author does not have given name: expected %q: got %q\n", tc.name, o.Name)
			}
		}
	}
}

func TestPost(t *testing.T) {
	// Specification: Post

	// When a new post is created
	// Then it has a unique ID
	// And it has the given title
	for _, tc := range []struct {
		author string
		title  string
	}{
		{"James Joyce", "Test Post"},
	} {
		ds, _ := memory.NewStore()
		id, _ := ds.CreatePost(tc.title)
		o, ok := ds.FindPostByID(id)
		if !ok {
			t.Errorf("post does not have unique ID: expected id %q: got no post found\n", id)
		} else if id != o.ID {
			t.Errorf("post does not have unique ID: expected id %q: got %q\n", id, o.ID)
		} else {
			if tc.title != o.Title {
				t.Errorf("post does not have the given title: expected %q: got %q\n", tc.title, o.Title)
			}
		}
	}
}
