package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Note struct {
	Name string
	Path string
}

// readAllNotes reads all notes from the knowledge base directory
func readAllNotes(root string) ([]Note, error) {
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

// SelectNote lets the user select a note from a list using fzf and returns the path to the selected note
func SelectNote(knowledgeBasePath string) (noteName string) {

	notes, err := readAllNotes(knowledgeBasePath)

	if err != nil {
		panic(err)
	}

	var noteNames []string
	var notesMap = make(map[string]Note)

	for _, note := range notes {
		noteNames = append(noteNames, note.Name)
		notesMap[note.Name] = note
	}

	reader := strings.NewReader(strings.Join(noteNames, "\n"))

	noteName, err = FzfSelect(reader)

	if err != nil {
		panic(err)
	}

	notePath := notesMap[noteName].Path

	return notePath
}

// OpenNoteWithEditor opens a note in the default editor
func OpenNoteWithEditor(notePath string) {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		fmt.Println("Error: no editor found in environment")
		os.Exit(1)
	}

	editorCmd := exec.Command(editor, notePath)

	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	err := editorCmd.Start()

	if err != nil {
		panic(err)
	}

	err = editorCmd.Wait()

	if err != nil {
		panic(err)
	}
}
