// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/mdhender/gobbs/internal/domains"
	"os"
	"path/filepath"
)

func abspath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	} else if sb, err := os.Stat(absPath); err != nil {
		return "", err
	} else if !sb.IsDir() {
		return "", domains.ErrNotDirectory
	}
	return absPath, nil
}

func isDirExists(path string) (bool, error) {
	sb, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return sb.IsDir(), nil
}

func isFileExists(path string) (bool, error) {
	sb, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return !sb.IsDir() && sb.Mode().IsRegular(), nil
}
