package cmd

import (
	"fmt"
	"kb/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

//  showCmd represents the show command
// TODO: extract reading kb folder, https://flaviocopes.com/go-list-files/
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show one note from knowledge base",
	Run: func(cmd *cobra.Command, args []string) {

		var files []string

		root := "/home/evan/notes/knowledge-base"
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, ".git") {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			files = append(files, path)
			return nil
		})

		if err != nil {
			panic(err)
		}

		reader := strings.NewReader(strings.Join(files, "\n"))

		notePath, err := utils.Fzf(reader)

		if err != nil {
			panic(err)
		}

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
