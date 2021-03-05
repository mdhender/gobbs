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

// Specification: Author

func TestAuthor(t *testing.T) {
	// When a new author is created
	// Then it has a unique ID
	ds, _ := memory.NewStore()
	id, _ := ds.CreateAuthor("James Joyce")
	o, ok := ds.FindAuthorByID(id)
	if !ok {
		t.Errorf("author does not have unique ID: expected id %q: got no author found\n", id)
	} else if o.ID != id {
		t.Errorf("author does not have unique ID: expected id %q: got %q\n", id, o.ID)
	} else {
		// And it has the given name
		if o.Name != "James Joyce" {
			t.Errorf("author does not have given name: expected %q: got %q\n", "James Joyce", o.Name)
		}
	}
}
