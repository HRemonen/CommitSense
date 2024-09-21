/*
Package cmd provides the main entry point for the CommitSense command-line tool.

CommitSense is a command-line tool designed to improve commit messages. This package contains the root command and initializes the application. It also defines global flags and configuration settings.

For more information on how to use CommitSense, run 'commitsense help'.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	"commitsense/pkg/config"
	"fmt"
	"os"

	colorprinter "commitsense/pkg/printer"

	"github.com/spf13/cobra"
)

var (
	showConfig bool
	setConfig  bool
	validArgs  = []string{"commit", "help"}
)

var rootCmd = &cobra.Command{
	Use:   "commitsense",
	Short: "A tool to improve commit messages",
	Long: `
CommitSense is a command-line tool that simplifies Git 
version control by providing an interactive and standardized way to stage 
files and create commit messages following the Conventional Commits specification.
`,
	TraverseChildren:   true,
	DisableSuggestions: false,
	Args:               cobra.OnlyValidArgs,
	ValidArgs:          validArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if showConfig {
			return config.ShowConfigSettings()
		}

		return cmd.Help()
	},
}

// SetVersion sets the version and build date for the application.
func SetVersion(version string, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s)", version, date)
}

func init() {
	cobra.OnInitialize()

	rootCmd.Flags().BoolVarP(&showConfig, "show-config", "s", false, "Show current configuration settings")
	rootCmd.Flags().BoolVarP(&setConfig, "set-config", "c", false, "Set new configuration settings")
}

// Execute command for the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		colorprinter.ColorPrint("error", "Error while executing: %v", err)

		os.Exit(1)
	}
}
