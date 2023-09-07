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
	"commitsense/pkg/commit"
	"commitsense/pkg/prompt"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// CommitCmd represents the commit command.
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit with a standardized message",
	Run: func(cmd *cobra.Command, args []string) {
		stagedFiles, err := GetStagedFiles()
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
			Label: "Enter a brief commit description",
			Validate: func(s string) error {
				if len(s) > 0 {
					return nil
				}
				return fmt.Errorf("please enter a valid string")
			},
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitBody, err := commit.PromptForMultilineString(prompt.Prompt{
			Label: "Enter a detailed commit body (press Enter twice to finish)",
			Validate: func(s string) error {
				// Accept any input
				return nil
			},
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		isBreakingChange, err := commit.PromptForBool(prompt.Prompt{
			Label: "Is this a breaking change?",
			Validate: func(s string) error {
				if s == "Y" || s == "N" || s == "y" || s == "n" {
					return nil
				}
				return fmt.Errorf("please enter Y or N")
			},
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

		commitMessage := CreateCommitMessage(commitType, commitScope, commitDescription, commitBody, isBreakingChange, breakingChangeDescription)

		if err := CreateGitCommit(commitMessage, stagedFiles); err != nil {
			fmt.Println("Error creating commit:", err)
			os.Exit(1)
		}

		fmt.Println("Commit created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

// GetStagedFiles returns a list of staged files.
func GetStagedFiles() ([]string, error) {
	statusCmd := exec.Command("git", "status", "--porcelain")
	output, err := statusCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var stagedFiles []string
	for _, line := range lines {
		if strings.HasPrefix(line, "M") || strings.HasPrefix(line, "A") {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				stagedFiles = append(stagedFiles, parts[1])
			}
		}
	}
	return stagedFiles, nil
}

// CreateCommitMessage creates a commit message in the Conventional Commits format.
func CreateCommitMessage(commitType, commitScope, commitDescription string, commitBody string, isBreakingChange bool, breakingChangeDescription string) string {
	commitMessage := commitType
	if commitScope != "" {
		commitMessage += "(" + commitScope + ")"
	}

	if isBreakingChange {
		commitMessage += "!"
	}

	commitMessage += ": " + commitDescription

	if commitBody != "" {
		commitMessage += "\n\n" + commitBody
	}

	if isBreakingChange {
		commitMessage += "\n\nBREAKING CHANGE: " + breakingChangeDescription
	}

	return commitMessage
}

// CreateGitCommit creates a Git commit with the given message and files.
func CreateGitCommit(message string, files []string) error {
	commitArgs := append([]string{"commit", "-m", message}, files...)
	commitGitCmd := exec.Command("git", commitArgs...) //nolint:gosec // because I do not think the users can do anything bad here
	commitGitCmd.Stdout = os.Stdout
	commitGitCmd.Stderr = os.Stderr

	return commitGitCmd.Run()
}
