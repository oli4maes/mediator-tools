package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mediator-tools",
	Short: "mediator-tools generates handler, request and response files.",
	Long: `mediator-tools is a CLI library for Go that empowers applications
		that make use of the mediator pattern.
		This application is a tool to generate the needed files.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
