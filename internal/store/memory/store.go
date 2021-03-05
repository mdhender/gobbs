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
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Store struct {
	locks struct {
		sync.RWMutex
	}
	authors map[string]*author
	posts   map[string]*post
}

func NewStore() (*Store, error) {
	return &Store{
		authors: make(map[string]*author),
		posts:   make(map[string]*post),
	}, nil
}

func (ds *Store) createAuthor(name string) (*author, error) {
	ds.locks.Lock()
	defer ds.locks.Unlock()
	for _, author := range ds.authors {
		if author.name == name {
			return nil, fmt.Errorf("duplicate name %q: %w", name, ErrDuplicateKey)
		}
	}
	author := author{
		id:   uuid.New().String(),
		name: name,
	}
	ds.authors[author.id] = &author
	return &author, nil
}

func (ds *Store) CreateAuthor(name string) (string, error) {
	author, err := ds.createAuthor(name)
	if err != nil {
		return "", err
	}
	return author.id, nil
}

func (ds *Store) CreatePost(authorID, title string) (string, error) {
	author := ds.findAuthorByID(authorID)
	if author == nil {
		return "", fmt.Errorf("author %q: %w", authorID, ErrNoDataFound)
	}
	post := post{
		id:     uuid.New().String(),
		author: author,
		title:  title,
	}
	ds.locks.Lock()
	ds.posts[post.id] = &post
	ds.locks.Unlock()
	return post.id, nil
}

func (ds *Store) findAuthorByID(id string) *author {
	ds.locks.RLock()
	author, _ := ds.authors[id]
	ds.locks.RUnlock()
	return author
}

func (ds *Store) FindAuthorByID(id string) (Author, bool) {
	author := ds.findAuthorByID(id)
	if author == nil {
		return Author{}, false
	}
	return Author{
		AuthorID: author.id,
		Name:     author.name,
	}, true
}

func (ds *Store) findPostByID(id string) *post {
	ds.locks.RLock()
	post, _ := ds.posts[id]
	ds.locks.RUnlock()
	return post
}

func (ds *Store) FindPostByID(id string) (Post, bool) {
	post := ds.findPostByID(id)
	if post == nil {
		return Post{}, false
	}
	return Post{
		PostID:   post.id,
		AuthorID: post.author.id,
		Title:    post.title,
	}, true
}
