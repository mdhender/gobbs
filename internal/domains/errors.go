// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package domains

// Error implements constant errors
type Error string

// Error implements the errors.Error interface
func (e Error) Error() string {
	return string(e)
}

const (
	ErrNotDirectory   = Error("not a directory")
	ErrNotFile        = Error("not a file")
	ErrNotImplemented = Error("not implemented")
	ErrServerShutdown = Error("server shutdown")
)
