package cmd

import (
	"fmt"
	"kb/utils"
	"os"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Open a note in knowledge base",
	Run: func(cmd *cobra.Command, args []string) {
		notePath := utils.SelectNote(KnowledgeBasePath)

		editor := os.Getenv("EDITOR")

		if editor == "" {
			fmt.Println("Error: no editor found in environment")
			os.Exit(1)
		}

		utils.OpenNoteWithEditor(notePath)

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
