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
	"crypto/sha256"
	"encoding/hex"
)

func New(salt []byte, rounds int) func([]byte) string {
	if len(salt) < 8 {
		panic("assert(len(salt) >= 8)")
	}
	if rounds < 8 {
		rounds = 8
	}
	saltHash := sha256.New()
	saltHash.Write(salt)
	salt = saltHash.Sum(nil)
	return func(secret []byte) string {
		hash := sha256.New()
		hash.Write(salt)
		hash.Write(secret)
		buf := hash.Sum(nil)
		for i := 0; i < rounds; i++ {
			hash.Reset()
			hash.Write(buf)
			hash.Write(salt)
			hash.Write(secret)
			buf = hash.Sum(nil)
		}
		return hex.EncodeToString(buf)
	}
}
