// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package app

import (
	"bytes"
	"github.com/mdhender/gobbs/internal/domains"
	"html/template"
	"log"
	"net/http"
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

func renderFragment(payload any, templateName string, templateFiles ...string) ([]byte, error) {
	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Printf("%s: %v\n", templateName, err)
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := t.ExecuteTemplate(buf, templateName, payload); err != nil {
		log.Printf("%s: %v\n", templateName, err)
		return nil, err
	}

	return buf.Bytes(), nil
}

// writeFragments writes the fragments to the response.
// any errors returned need to be handled by the caller.
// note that the caller should call http.Error() to close
// the connection if there is an error writing, but the
// header and any prior fragment writes have already been
// sent to the client.
func writeFragments(w http.ResponseWriter, r *http.Request, fragments ...[]byte) (bytesWritten int, err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	for _, fragment := range fragments {
		n, err := w.Write(fragment)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}
	return bytesWritten, nil
}
