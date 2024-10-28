// Copyright (c) 2021-2024 Michael D Henderson. All rights reserved.

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
