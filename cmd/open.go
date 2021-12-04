package cmd

import (
	"fmt"
	"kb/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Open a note in knowledge base",
	Run: func(cmd *cobra.Command, args []string) {
		notePath, err := utils.SelectNote(KnowledgeBasePath)

		if err != nil {
			panic(err)
		}

		editor := os.Getenv("EDITOR")

		if editor == "" {
			fmt.Println("Error: no editor found in environment")
			os.Exit(1)
		}

		editorCmd := exec.Command(editor, notePath)

		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		err = editorCmd.Start()

		if err != nil {
			panic(err)
		}

		err = editorCmd.Wait()

		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
