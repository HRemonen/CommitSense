/*
Package cmd provides commands for the commitsense application.

This package contains the main commands and functionality for the commitsense application. It includes commands for interactive file selection and staging for commits. Additionally, it provides commands for creating standardized commit messages.

Usage:
  - Use the 'commit' command to create a commit with a standardized commit message.

For more information, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	"commitsense/internal/validators"
	"commitsense/pkg/commit"
	"os"

	colorprinter "commitsense/internal/printer"
	csprompt "commitsense/pkg/prompt"

	"github.com/spf13/cobra"
)

var (
	isCoAuthored     bool
	isBreakingChange bool
)

// CommitCmd represents the commit command.
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit with a standardized message",
	Run: func(_ *cobra.Command, _ []string) {
		stagedFiles, err := commit.GetStagedFiles()
		if err != nil {
			colorprinter.ColorPrint("error", "Error: %v", err)
			os.Exit(1)
		}

		commitType, err := csprompt.CommitType("Select a commit type")
		if err != nil {
			colorprinter.ColorPrint("error", "Error prompting for the commit type: %v", err)
			os.Exit(1)
		}

		commitScope, err := csprompt.String("Enter a commit scope (optional)", nil)
		if err != nil {
			colorprinter.ColorPrint("error", "Error prompting for the commit scope: %v", err)
			os.Exit(1)
		}

		commitDescription, err := csprompt.String(
			"Enter a brief commit description",
			validators.ValidateStringNotEmpty,
		)
		if err != nil {
			colorprinter.ColorPrint("error", "Error prompting for the commit description: %v", err)
			os.Exit(1)
		}

		commitBody, err := csprompt.MultilineString(
			"Enter a detailed commit body (press Enter twice to finish)",
		)
		if err != nil {
			colorprinter.ColorPrint("error", "Error prompting for the commit body: %v", err)
			os.Exit(1)
		}

		var coAuthors []string
		if isCoAuthored {
			coAuthors, err = csprompt.CoAuthors(
				"Enter Co-Author information ",
			)
			if err != nil {
				colorprinter.ColorPrint("error", "Error prompting for the co-authors: %v", err)
				os.Exit(1)
			}
		}

		var breakingChangeDescription string
		if isBreakingChange {
			breakingChangeDescription, err = csprompt.String(
				"Enter a description of the breaking change",
				nil,
			)
			if err != nil {
				colorprinter.ColorPrint("error", "Error prompting for the breaking change description: %v", err)
				os.Exit(1)
			}
		}

		c := commit.Commit{
			CommitType:                commitType,
			CommitScope:               commitScope,
			CommitDescription:         commitDescription,
			CommitBody:                commitBody,
			IsCoAuthored:              isCoAuthored,
			CoAuthors:                 coAuthors,
			IsBreakingChange:          isBreakingChange,
			BreakingChangeDescription: breakingChangeDescription,
			StagedFiles:               stagedFiles,
		}

		if err := c.CreateGitCommit(); err != nil {
			colorprinter.ColorPrint("error", "Error creating a commit: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().BoolVarP(&isCoAuthored, "is-coauthored", "a", false, "Commit is co-authored")
	commitCmd.Flags().BoolVarP(&isBreakingChange, "is-breaking", "b", false, "Commit is introducing a breaking change")
}
