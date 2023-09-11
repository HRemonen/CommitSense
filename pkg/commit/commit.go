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
	statusCmd := "git status --porcelain --untracked-files=all | grep '^[A|C|M|D|R]'"
	cmd := exec.Command("bash", "-c", statusCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var stagedFiles []string
	for _, line := range lines {
		// Strip leading and trailing whitespace
		line = strings.TrimSpace(line)
	}
	return stagedFiles, nil
}

// CreateCommitMessage creates a commit message in the Conventional Commits format.
func createCommitMessage(commitInfo Info) string {
	var commitMessage string

	if commitInfo.CommitType == "docs" {
		commitMessage = "[skip ci] " + commitInfo.CommitType
	} else {
		commitMessage = commitInfo.CommitType
	}

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
		commitMessage += "\n"
		commitMessage += "\nBREAKING CHANGE: " + commitInfo.BreakingChangeDescription
	}

	if commitInfo.IsCoAuthored {
		commitMessage += "\n"
		for _, coauth := range commitInfo.CoAuthors {
			commitMessage += "\nCo-authored-by: " + coauth
		}
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
