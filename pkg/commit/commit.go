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

	"github.com/go-git/go-git/v5"
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
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, errors.New("could not open the Git repository")
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, errors.New("could not get the Git worktree")
	}

	status, err := worktree.Status()
	if err != nil {
		return nil, errors.New("could not get the Git status")
	}

	var stagedFiles []string
	for file, status := range status {
		if status.Staging == git.Added || status.Staging == git.Modified || status.Staging == git.Deleted || status.Staging == git.Renamed || status.Staging == git.Copied {
			stagedFiles = append(stagedFiles, file)
		}
	}

	if len(stagedFiles) == 0 {
		return nil, errors.New("could not get staged files, is the files added for staging?")
	}

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

// CreateGitCommit creates a Git commit with the given message and files.
func CreateGitCommit(commit Commit, files []string) error {
	commitMessage := createCommitMessage(commit)

	commitArgs := append([]string{"commit", "-m", commitMessage}, files...)

	commitGitCmd := exec.Command("git", commitArgs...) //nolint:gosec // because I do not think the users can do anything bad here
	commitGitCmd.Stdout = os.Stdout
	commitGitCmd.Stderr = os.Stderr

	return commitGitCmd.Run()
}
