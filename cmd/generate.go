package cmd

import (
	"fmt"
	"github.com/oli4maes/mediator-tools/internal/filegeneration"
	"github.com/spf13/cobra"
)

var feature string
var lang string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate 'handler', 'request' and 'response' files for a feature.",
	Long: `Generate all needed files for a given feature.
			These files are: handler, request and response`,
	Run: func(cmd *cobra.Command, args []string) {
		if lang != "" && feature != "" {
			langId, err := filegeneration.GetLangType(lang)
			if err != nil {
				fmt.Println(err)
			}

			err = filegeneration.GenerateFiles(feature, filegeneration.RequestFileType, langId)
			if err != nil {
				return
			}
			err = filegeneration.GenerateFiles(feature, filegeneration.ResponseFileType, langId)
			if err != nil {
				return
			}
			err = filegeneration.GenerateFiles(feature, filegeneration.HandlerFileType, langId)
			if err != nil {
				return
			}

			fmt.Println("files generated")
		}

		if lang == "" {
			fmt.Println("please provide a valid language with --language flag (supports: 'go', 'cs')")
		}
		if feature == "" {
			fmt.Println("please provide a feature name --feature flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&feature, "feature", "f", "", "feature name")
	generateCmd.Flags().StringVarP(&lang, "language", "l", "", "programming language (supports: 'go', 'cs')")
}
