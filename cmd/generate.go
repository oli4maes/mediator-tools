package cmd

import (
	"fmt"
	"github.com/oli4maes/mediator-tools/internal/filegeneration"
	"github.com/spf13/cobra"
)

var feature string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate 'handler', 'request' and 'response' files for a feature.",
	Long: `Generate all needed files for a given feature.
These files are: handler, request and response`,
	Run: func(cmd *cobra.Command, args []string) {
		err := filegeneration.GenerateFiles(feature, filegeneration.RequestFileType)
		if err != nil {
			return
		}
		err = filegeneration.GenerateFiles(feature, filegeneration.ResponseFileType)
		if err != nil {
			return
		}
		err = filegeneration.GenerateFiles(feature, filegeneration.HandlerFileType)
		if err != nil {
			return
		}

		fmt.Println("files generated")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//generateCmd.PersistentFlags().String("feature", "", "The name of the feature.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().StringVarP(&feature, "feature", "f", "", "usage")
}
