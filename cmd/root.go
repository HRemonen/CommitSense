/*
Package cmd provides the main entry point for the CommitSense command-line tool.

CommitSense is a command-line tool designed to improve commit messages. This package contains the root command and initializes the application. It also defines global flags and configuration settings.

For more information on how to use CommitSense, run 'commitsense help'.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	"commitsense/pkg/config"
	colorprinter "commitsense/pkg/printer"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	showVersion bool
	showConfig  bool
	setConfig   bool
	validArgs   = []string{"add", "commit"}
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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !config.Exists() {
			err := config.CreateDefaultConfig()
			if err != nil {
				return err
			}
			colorprinter.ColorPrint("info", "\nCould not find an existing configuration file")
			colorprinter.ColorPrint("success", "Created default configuration file at ~/.commitsense.yaml")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if showConfig {
			return showConfigSettings()
		}

		return cmd.Help()
	},
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

func showConfigSettings() error {
	colorprinter.ColorPrint("success", "\nShowing current configuration settings")

	config, err := config.ReadConfigFile()
	if err != nil {
		colorprinter.ColorPrint("error", "Error reading configuration file: %v", err)

		return err
	}

	colorprinter.ColorPrint("bold", "\nAllowed commit types:")
	printYAML(config.CommitTypes)

	colorprinter.ColorPrint("bold", "Skipping CI on types:")
	printYAML(config.SkipCITypes)

	return nil
}

func printYAML(data interface{}) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		colorprinter.ColorPrint("error", "Error printing YAML: %v", err)

		return
	}

	// Use strings.Replace to add proper indentation
	indentedYAML := strings.Replace(string(yamlData), "\n", "\n  ", -1)
	fmt.Println("  " + indentedYAML)
}
