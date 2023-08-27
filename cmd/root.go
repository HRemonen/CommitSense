/*
Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "commitsense",
	Short: "A tool to improve commit messages and manage GitHub issues",
	Long:  `CommitSense is a command-line tool that helps you follow best practices for commit messages and manage GitHub issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is the default action when the tool is run without a subcommand.
		fmt.Println("Welcome to CommitSense!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.commitsense.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
