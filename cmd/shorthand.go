/*
Package cmd provides commands for the commitsense application.

This package contains the main commands and functionality for the commitsense application. It includes commands for interactive file selection and staging for commits. Additionally, it provides commands for creating standardized commit messages.

Usage:
  - Use the 'commit' command to create a commit with a standardized commit message.

For more information, refer to the package-specific functions and commands.

Copyright © 2024 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	colorprinter "commitsense/internal/printer"
	"commitsense/pkg/commit"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	commitScope    string
	breakingChange bool
)

func shorthandCmd(commitType string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   commitType + " [message]",
		Short: fmt.Sprintf("Create a git commit with type %s", commitType),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			stagedFiles, err := commit.GetStagedFiles()
			if err != nil {
				colorprinter.ColorPrint("error", "Error: %v", err)
				os.Exit(1)
			}

			commitDescription := strings.Join(args, " ")

			c := commit.Commit{
				CommitType:        commitType,
				CommitScope:       commitScope,
				CommitDescription: commitDescription,
				IsCoAuthored:      isCoAuthored,
				IsBreakingChange:  breakingChange,
				StagedFiles:       stagedFiles,
			}

			if err := c.CreateGitCommit(); err != nil {
				colorprinter.ColorPrint("error", "Error creating a commit: %v", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&commitScope, "scope", "s", "", "Commit scope")
	cmd.Flags().BoolVarP(&breakingChange, "is-breaking", "b", false, "Commit is introducing a breaking change")

	return cmd
}

func init() {
	commitTypes := []string{"build", "ci", "chore", "docs", "feat", "fix", "perf", "refactor", "revert", "style", "test"}

	for _, commitType := range commitTypes {
		rootCmd.AddCommand(shorthandCmd(commitType))
	}
}
