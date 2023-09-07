package commit

import (
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
