package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var KNOWLEDGE_BASE_PATH_ENV_VAR = "KB_PATH"
var KnowledgeBasePath = os.Getenv(KNOWLEDGE_BASE_PATH_ENV_VAR)

var rootCmd = &cobra.Command{
	Use:   "kb",
	Short: "A cli tool for managing my knowledge base",
}

func Execute() {
	if KnowledgeBasePath == "" {
		err := fmt.Errorf("ERROR: %s env variable is not set", KNOWLEDGE_BASE_PATH_ENV_VAR)
		fmt.Println(err)
		os.Exit(1)
	}

	cobra.CheckErr(rootCmd.Execute())
}
