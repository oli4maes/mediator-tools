package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "mediator-tools version",
	Long:  `Every application has a version, this is mediator-tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mediator-tools v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
