/*
Package commit provides functionality for creating Git commits.

This file includes utility functions for interacting with Git and creating Git commits.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package commit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
func createCommitMessage(commitInfo Info) string {
	commitMessage := commitInfo.CommitType
	if commitInfo.CommitScope != "" {
		commitMessage += "(" + commitInfo.CommitScope + ")"
	}

	if commitInfo.IsBreakingChange {
		commitMessage += "!"
	}

	commitMessage += ": " + commitInfo.CommitDescription

	if commitInfo.CommitBody != "" {
		commitMessage += "\n\n" + commitInfo.CommitBody
	}

	if commitInfo.IsBreakingChange {
		commitMessage += "\n\n"
	}

	if commitInfo.IsBreakingChange {
		commitMessage += "BREAKING CHANGE: " + commitInfo.BreakingChangeDescription
	}

	return commitMessage
}

// CreateGitCommit creates a Git commit with the given message and files.
func CreateGitCommit(commitInfo Info, files []string) error {
	commitMessage := createCommitMessage(commitInfo)

	fmt.Println(commitMessage, files)

	commitArgs := append([]string{"commit", "-m", commitMessage}, files...)

	commitGitCmd := exec.Command("git", commitArgs...) //nolint:gosec // because I do not think the users can do anything bad here
	commitGitCmd.Stdout = os.Stdout
	commitGitCmd.Stderr = os.Stderr

	return commitGitCmd.Run()
}
