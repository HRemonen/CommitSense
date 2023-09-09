/*
Package cmd provides commands for the commitsense application.

This package contains the main commands and functionality for the commitsense application. It includes commands for interactive file selection and staging for commits. Additionally, it provides commands for creating standardized commit messages.

Usage:
  - Use the 'add' command to interactively select files to stage for committing.
  - Use the 'commit' command to create a commit with a standardized commit message.

For more information, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	"commitsense/pkg/author"
	"commitsense/pkg/commit"
	"commitsense/pkg/prompt"
	"commitsense/pkg/validators"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CommitCmd represents the commit command.
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit with a standardized message",
	Run: func(cmd *cobra.Command, args []string) {
		stagedFiles, err := commit.GetStagedFiles()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		commitType, err := commit.PromptCommitType()
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitScope, err := commit.PromptForString(prompt.Prompt{
			Label: "Enter a commit scope (optional)",
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitDescription, err := commit.PromptForString(prompt.Prompt{
			Label:    "Enter a brief commit description",
			Validate: validators.ValidateStringNotEmpty,
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitBody, err := commit.PromptForMultilineString(prompt.Prompt{
			Label: "Enter a detailed commit body (press Enter twice to finish)",
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		isCoAuthored, err := commit.PromptForBool(prompt.Prompt{
			Label:    "Is this commit co-authored by someone?",
			Validate: validators.ValidateStringYesNo,
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		var coAuthors []string
		if isCoAuthored {
			suggestedCoAuthors, err := author.GetSuggestedCoAuthors()
			if err != nil {
				fmt.Println("Prompt failed:", err)
				os.Exit(1)
			}

			coAuthors, err = commit.PromptForCoAuthors(prompt.Prompt{
				Label:     "Select authors that are involded",
				Items:     suggestedCoAuthors,
				CursorPos: 0,
			})
			if err != nil {
				fmt.Println("Prompt failed:", err)
				os.Exit(1)
			}
		}

		isBreakingChange, err := commit.PromptForBool(prompt.Prompt{
			Label:    "Is this a breaking change?",
			Validate: validators.ValidateStringYesNo,
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		var breakingChangeDescription string
		if isBreakingChange {
			breakingChangeDescription, err = commit.PromptForString(prompt.Prompt{
				Label: "Enter a description of the breaking change",
			})
			if err != nil {
				fmt.Println("Prompt failed:", err)
				os.Exit(1)
			}
		}

		commitInfo := commit.Info{
			CommitType:                commitType,
			CommitScope:               commitScope,
			CommitDescription:         commitDescription,
			CommitBody:                commitBody,
			IsCoAuthored:              isCoAuthored,
			CoAuthors:                 coAuthors,
			IsBreakingChange:          isBreakingChange,
			BreakingChangeDescription: breakingChangeDescription,
		}

		if err := commit.CreateGitCommit(commitInfo, stagedFiles); err != nil {
			fmt.Println("Error creating commit:", err)
			os.Exit(1)
		}

		fmt.Println("Commit created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
