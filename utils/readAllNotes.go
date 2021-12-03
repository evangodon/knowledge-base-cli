package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type Note struct {
	Name string
	Path string
}

func ReadAllNotes(root string) ([]Note, error) {
	var notes []Note

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".git") {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		notes = append(notes, Note{
			Name: info.Name(),
			Path: path,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return notes, nil

}
