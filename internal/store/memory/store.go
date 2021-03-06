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
	"time"
)

func NewStore() (*Store, error) {
	return &Store{
		authors:  make(map[string]*author),
		messages: make(map[string]*message),
		threads:  make(map[string]*thread),
	}, nil
}

type Store struct {
	locks struct {
		sync.RWMutex
	}
	authors  map[string]*author
	messages map[string]*message
	threads  map[string]*thread
}

type author struct {
	id    string
	name  string
	roles map[string]bool
}

type message struct {
	id      string
	author  *author
	subject string
	created time.Time
	body    string
}

type thread struct {
	id       string
	messages []*message
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

func (ds *Store) createMessage(authorID, subject, body string) (*message, error) {
	author := ds.findAuthorByID(authorID)
	if author == nil {
		return nil, fmt.Errorf("author %q: %w", authorID, ErrNoDataFound)
	}
	message := message{
		id:      uuid.New().String(),
		author:  author,
		subject: subject,
		body:    body,
	}
	ds.locks.Lock()
	ds.messages[message.id] = &message
	ds.locks.Unlock()
	return &message, nil
}

func (ds *Store) createThread(authorID, subject, body string) (*thread, error) {
	msg, err := ds.createMessage(authorID, subject, body)
	if err != nil {
		return nil, err
	}
	thread := &thread{
		id:       uuid.New().String(),
		messages: []*message{msg},
	}
	ds.locks.Lock()
	ds.threads[thread.id] = thread
	ds.locks.Unlock()
	return thread, nil
}

func (ds *Store) findAuthorByID(id string) *author {
	ds.locks.RLock()
	author, _ := ds.authors[id]
	ds.locks.RUnlock()
	return author
}

func (ds *Store) findMessageByID(id string) *message {
	ds.locks.RLock()
	message, _ := ds.messages[id]
	ds.locks.RUnlock()
	return message
}

func (ds *Store) findThreadByID(id string) *thread {
	ds.locks.RLock()
	thread, _ := ds.threads[id]
	ds.locks.RUnlock()
	return thread
}
