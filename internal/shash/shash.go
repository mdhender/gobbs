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

// Package shash implements a simple hashing package for passwords.
// This is for testing only. Do not use it in a production environment.
// Do not use it if you care about the security of your data.
package shash

import (
	"crypto/sha1"
	"fmt"
)

func New(salt []byte) func([]byte) string {
	if len(salt) == 0 {
		panic("assert(len(salt) != 0)")
	} else if len(salt) < 128 {
		tmp := make([]byte, 128, 128)
		for i := 0; i < 128; i++ {
			tmp[i] = salt[i%len(salt)]
		}
		salt = sha1.New().Sum(tmp)
	} else {
		salt = sha1.New().Sum(salt)
	}
	return func(secret []byte) string {
		hash := sha1.New()
		hash.Write(salt)
		buf := hash.Sum(nil)
		hash.Reset()
		hash.Write(buf)
		hash.Write(secret)
		buf = hash.Sum(nil)
		hash.Reset()
		hash.Write(buf)
		hash.Write(salt)
		buf = hash.Sum(nil)
		return fmt.Sprintf("%d % x", len(buf), buf)
	}
}
