package cmd

import (
	"fmt"
	"io/ioutil"
	"kb/utils"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Create a new note",
	Run: func(cmd *cobra.Command, args []string) {

		noteName := utils.GetUserInput("Enter note title: ")

		if noteName == "" {
			fmt.Println("No note title provided")
			os.Exit(0)
		}

		noteName = strings.TrimSpace(noteName)
		fileName := strings.Replace(noteName, " ", "_", -1)

		currentDate := time.Now().Format("02-Jan-2006")

		content := fmt.Sprintf("---\ntitle: %s\ncreatedAt: %s\n---\n", noteName, currentDate)
		filePath := fmt.Sprintf("%s/%s.md", KnowledgeBasePath, fileName)

		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("Note with name '%s' already exists\n", noteName)
			utils.ConfirmOperation("Do you want to overwrite it?")

		}

		err := ioutil.WriteFile(filePath, []byte(content), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

		utils.OpenNoteWithEditor(filePath)

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
