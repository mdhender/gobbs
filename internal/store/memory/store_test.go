// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

package memory_test

import (
	"errors"
	"github.com/mdhender/gobbs/internal/store/memory"
	"testing"
)

func TestAuthor(t *testing.T) {
	// Specification: Author

	for _, tc := range []struct {
		name string
	}{
		{"James Joyce"},
	} {
		ds, _ := memory.NewStore()

		// When a new author is created
		id, _ := ds.CreateAuthor(tc.name)

		// Then it has a unique ID
		o, ok := ds.FindAuthorByID(id)
		if !ok {
			t.Errorf("author does not have unique ID: expected id %q: got no author found\n", id)
			continue
		}
		if id != o.AuthorID {
			t.Errorf("author does not have unique ID: expected id %q: got %q\n", id, o.AuthorID)
			continue
		}

		// And it has the given name
		if tc.name != o.Name {
			t.Errorf("author does not have given name: expected %q: got %q\n", tc.name, o.Name)
		}
	}

	for _, tc := range []struct {
		name string
	}{
		{"James Joyce"},
	} {
		ds, _ := memory.NewStore()
		_, _ = ds.CreateAuthor(tc.name)

		// When a new author is created with an existing name
		id, err := ds.CreateAuthor(tc.name)

		// Then it is rejected as a duplicate key
		if err == nil {
			t.Errorf("author with existing name is not rejected: expected error: got id %q\n", id)
			continue
		}
		if !errors.Is(err, memory.ErrDuplicateKey) {
			t.Errorf("author with existing name is rejected with unexpected error: expected %q: got %q\n", memory.ErrDuplicateKey, err)
			continue
		}
	}
}

func TestMessage(t *testing.T) {
	// Specification: Message

	for _, tc := range []struct {
		author  string
		subject string
		body    string
	}{
		{"James Joyce", "Test Message", "Lorem quicksand tenor tomato."},
	} {
		ds, _ := memory.NewStore()

		authorID, _ := ds.CreateAuthor(tc.author)

		// When a new message is created
		id, _ := ds.CreateMessage(authorID, tc.subject, tc.body)

		// Then it has a unique ID
		o, ok := ds.FindMessageByID(id)
		if !ok {
			t.Errorf("message does not have unique ID: expected id %q: got no message found\n", id)
			continue
		}
		if id != o.MessageID {
			t.Errorf("message does not have unique ID: expected id %q: got %q\n", id, o.MessageID)
			continue
		}

		// And it has the given title
		if tc.subject != o.Subject {
			t.Errorf("message does not have the given subject: expected %q: got %q\n", tc.subject, o.Subject)
		}

		// And it has the given author
		if authorID != o.AuthorID {
			t.Errorf("message does not have the given author: expected %q: got %q\n", authorID, o.AuthorID)
		}

		// And it has the given body
		if tc.body != o.Body {
			t.Errorf("message does not have the given body: expected %q: got %q\n", tc.body, o.Body)
		}
	}
}

func TestThread(t *testing.T) {
	// Specification: Thread

	for _, tc := range []struct {
		author  string
		subject string
		body    string
	}{
		{"James Joyce", "Test Thread", "Lorem quicksand tenor tomato."},
	} {
		ds, _ := memory.NewStore()
		authorID, _ := ds.CreateAuthor(tc.author)

		// When a new thread is created
		id, _ := ds.CreateThread(authorID, tc.subject, tc.body)

		// Then it has a unique ID
		thread, ok := ds.FindThreadByID(id)
		if !ok {
			t.Errorf("thread does not have unique ID: expected id %q: got no thread found\n", id)
			continue
		}
		if id != thread.ThreadID {
			t.Errorf("thread does not have unique ID: expected id %q: got %q\n", id, thread.ThreadID)
			continue
		}
		if id2, _ := ds.CreateThread(authorID, tc.subject, tc.body); id == id2 {
			t.Errorf("thread does not have unique ID: id %q is not unique\n", id)
			continue
		}

		// And it contains the given message
		if len(thread.Messages) == 0 {
			t.Errorf("thread does not contain the given message: found 0 messages\n")
			continue
		} else if len(thread.Messages) != 1 {
			t.Errorf("thread contains unexpected messages: expected %d: found %d messages\n", 1, len(thread.Messages))
			continue
		}
		if authorID != thread.Messages[0].AuthorID {
			t.Errorf("thread does not contain the given message: expected author %q: found %q\n", authorID, thread.Messages[0].AuthorID)
		}
		if tc.subject != thread.Messages[0].Subject {
			t.Errorf("thread does not contain the given message: expected subject %q: found %q\n", tc.subject, thread.Messages[0].Subject)
		}
	}
}
