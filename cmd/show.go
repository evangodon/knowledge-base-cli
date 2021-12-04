package cmd

import (
	"fmt"
	"kb/utils"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

//  showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s"},
	Short:   "Show one note from knowledge base",
	Run: func(cmd *cobra.Command, args []string) {

		notePath := utils.SelectNote(KnowledgeBasePath)
		data, err := os.ReadFile(notePath)

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
