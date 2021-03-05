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

type Author struct {
	AuthorID string
	Name     string
}

type Messages []Message

type Message struct {
	MessageID string
	AuthorID  string
	Subject   string
	Body      string
}

type Thread struct {
	ThreadID string
	Messages Messages
}

func (ds *Store) CreateAuthor(name string) (string, error) {
	author, err := ds.createAuthor(name)
	if err != nil {
		return "", err
	}
	return author.id, nil
}

func (ds *Store) CreateMessage(authorID, subject, body string) (string, error) {
	message, err := ds.createMessage(authorID, subject, body)
	if err != nil {
		return "", err
	}
	return message.id, nil
}

func (ds *Store) CreateThread(authorID, subject, body string) (string, error) {
	thread, err := ds.createThread(authorID, subject, body)
	if err != nil {
		return "", err
	}
	return thread.id, nil
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

func (ds *Store) FindMessageByID(id string) (Message, bool) {
	message := ds.findMessageByID(id)
	if message == nil {
		return Message{}, false
	}
	return Message{
		MessageID: message.id,
		AuthorID:  message.author.id,
		Subject:   message.subject,
		Body:      message.body,
	}, true
}

func (ds *Store) FindThreadByID(id string) (Thread, bool) {
	var t Thread
	thread := ds.findThreadByID(id)
	if thread == nil {
		return t, false
	}
	t.ThreadID = thread.id
	for _, msg := range thread.messages {
		t.Messages = append(t.Messages, Message{
			MessageID: msg.id,
			AuthorID:  msg.author.id,
			Subject:   msg.subject,
			Body:      msg.body,
		})
	}
	return t, true
}
