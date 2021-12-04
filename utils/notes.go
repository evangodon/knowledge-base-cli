package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// OpenNote opens a note in the default editor
func OpenNote(notePath string) {
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
