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

type Post struct {
	PostID   string
	AuthorID string
	Title    string
}

func (ds *Store) CreateAuthor(name string) (string, error) {
	author, err := ds.createAuthor(name)
	if err != nil {
		return "", err
	}
	return author.id, nil
}

func (ds *Store) CreatePost(authorID, title string) (string, error) {
	post, err := ds.createPost(authorID, title)
	if err != nil {
		return "", err
	}
	return post.id, nil
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
