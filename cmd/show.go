package cmd

import (
	"fmt"
	"kb/utils"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

//  showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s"},
	Short:   "Show one note from knowledge base",
	Run: func(cmd *cobra.Command, args []string) {

		notes, err := utils.ReadAllNotes(KnowledgeBasePath)

		if err != nil {
			panic(err)
		}

		var noteNames []string
		var notesMap = make(map[string]utils.Note)

		for _, note := range notes {
			noteNames = append(noteNames, note.Name)
			notesMap[note.Name] = note
		}

		reader := strings.NewReader(strings.Join(noteNames, "\n"))

		notePath, err := utils.Fzf(reader)

		if err != nil {
			panic(err)
		}

		data, err := os.ReadFile(notesMap[notePath].Path)

		if err != nil {
			panic(err)
		}

		out, err := glamour.Render(string(data), "dark")

		if err != nil {
			panic(err)
		}

		fmt.Print(out)

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
