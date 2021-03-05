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

package memory

import (
	"github.com/google/uuid"
	"sync"
)

type Store struct {
	locks struct {
		sync.RWMutex
	}
	authors map[string]*Author
	posts   map[string]*Post
}

func NewStore() (*Store, error) {
	return &Store{
		authors: make(map[string]*Author),
		posts:   make(map[string]*Post),
	}, nil
}

func (ds *Store) CreateAuthor(name string) (string, error) {
	auth := Author{
		ID:   uuid.New().String(),
		Name: name,
	}
	ds.locks.Lock()
	ds.authors[auth.ID] = &auth
	ds.locks.Unlock()
	return auth.ID, nil
}

func (ds *Store) CreatePost(title string) (string, error) {
	post := Post{
		ID: uuid.New().String(),
	}
	ds.locks.Lock()
	ds.posts[post.ID] = &post
	ds.locks.Unlock()
	return post.ID, nil
}

func (ds *Store) FindAuthorByID(id string) (Author, bool) {
	ds.locks.RLock()
	auth, ok := ds.authors[id]
	ds.locks.RUnlock()
	if !ok {
		return Author{}, false
	}
	return *auth, true
}

func (ds *Store) FindPostByID(id string) (Post, bool) {
	ds.locks.RLock()
	post, ok := ds.posts[id]
	ds.locks.RUnlock()
	if !ok {
		return Post{}, false
	}
	return *post, true
}
