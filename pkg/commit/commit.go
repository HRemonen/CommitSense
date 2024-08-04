/*
Package commit provides functionality for creating Git commits.

This file includes utility functions for interacting with Git and creating Git commits.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package commit

import (
	"commitsense/pkg/config"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// Commit represents information needed for creating a Git commit.
type Commit struct {
	CommitType                string
	CommitScope               string
	CommitDescription         string
	CommitBody                string
	IsCoAuthored              bool
	CoAuthors                 []string
	IsBreakingChange          bool
	BreakingChangeDescription string
}

// getStringsFromTerminalOutput takes the os/exec functions returned byte array
// and transforms the bytes into an array of lines
func getStagedFilesFromTerminalOutput(output []byte) []string {
	lines := strings.Split(string(output), "\n")

	var stagedFiles []string
	for _, line := range lines {
		// Strip leading and trailing whitespace
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)
		if len(parts) == 2 {
			stagedFiles = append(stagedFiles, parts[1])
		}
	}

	return stagedFiles
}

// GetStagedFiles returns a list of staged files.
func GetStagedFiles() ([]string, error) {
	statusCmd := "git status --porcelain --untracked-files=all | grep '^[A|C|M|D|R]'"
	cmd := exec.Command("bash", "-c", statusCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New("could not get staged files, is the files added for staging?")
	}

	stagedFiles := getStagedFilesFromTerminalOutput(output)

	return stagedFiles, nil
}

// CreateCommitMessage creates a commit message in the Conventional Commits format.
func createCommitMessage(commit Commit) string {
	var commitMessage string

	commitMessage = commit.CommitType
	config, _ := config.ReadConfigFile()

	if commit.CommitScope != "" {
		commitMessage += "(" + commit.CommitScope + ")"
	}

	if commit.IsBreakingChange {
		commitMessage += "!"
	}

	commitMessage += ": " + commit.CommitDescription

	if commit.CommitBody != "" {
		commitMessage += "\n\n" + commit.CommitBody
	}

	for _, skipType := range config.SkipCITypes {
		if commit.CommitType == skipType {
			commitMessage += "\n[skip ci]"
			break
		}
	}

	if commit.IsBreakingChange {
		commitMessage += "\n"
		commitMessage += "\nBREAKING CHANGE: " + commit.BreakingChangeDescription
	}

	if commit.IsCoAuthored {
		commitMessage += "\n"
		for _, coauth := range commit.CoAuthors {
			commitMessage += "\nCo-authored-by: " + coauth
		}
	}

	return commitMessage
}

// CreateGitCommit creates a Git commit from the commit struct.
func CreateGitCommit(commit Commit, stagedFiles []string) error {
	commitMessage := createCommitMessage(commit)

	commitArgs := append([]string{"commit", "-m", commitMessage}, stagedFiles...)

	commitGitCmd := exec.Command("git", commitArgs...) //nolint:gosec // because I do not think the users can do anything bad here
	commitGitCmd.Stdout = os.Stdout
	commitGitCmd.Stderr = os.Stderr

	if err := commitGitCmd.Run(); err != nil {
		return err
	}

	updateIndexCmd := exec.Command("git", "update-index", "-g")
	updateIndexCmd.Stdout = os.Stdout
	updateIndexCmd.Stderr = os.Stderr

	return updateIndexCmd.Run()
}
