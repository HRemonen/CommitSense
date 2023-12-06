/*
Package commit provides functionality for creating Git commits.

This file includes utility functions for interacting with Git and creating Git commits.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package commit

import (
	"commitsense/pkg/config"
	colorprinter "commitsense/pkg/printer"
	"errors"

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
	repo, err := git.PlainOpen(".")
	if err != nil {
		return errors.New("could not open the Git repository")
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return errors.New("could not get the Git worktree")
	}

	commitMessage := createCommitMessage(commit)

	createdCommit, err := worktree.Commit(commitMessage, &git.CommitOptions{})

	if err != nil {
		return errors.New("could not create the Git commit")
	}

	commitObj, err := repo.CommitObject(createdCommit)
	if err != nil {
		return errors.New("could not get the Git commit object")
	}

	colorprinter.ColorPrint("success", "Created new commit:")
	colorprinter.ColorPrint("success", commitObj.String())

	return nil
}
