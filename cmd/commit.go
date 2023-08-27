package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// CommitPrompt represents a prompt for user input.
type CommitPrompt struct {
	Label    string
	Validate func(string) error
	Default  string
}

// SelectCommitType prompts the user to select a commit type.
func SelectCommitType() (string, error) {
	promptType := promptui.Select{
		Label: "Select a commit type",
		Items: []string{"feat", "fix", "chore", "docs", "style", "refactor", "perf", "test", "build", "ci"},
	}
	_, typeResult, err := promptType.Run()
	return typeResult, err
}

// PromptForString prompts the user to enter a string.
func PromptForString(prompt CommitPrompt) (string, error) {
	promptUI := promptui.Prompt{
		Label:    prompt.Label,
		Validate: prompt.Validate,
		Default:  prompt.Default,
	}
	return promptUI.Run()
}

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

		commitType, err := SelectCommitType()
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitScope, err := PromptForString(CommitPrompt{
			Label: "Enter a commit scope (optional):",
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitDescription, err := PromptForString(CommitPrompt{
			Label: "Enter a brief commit description:",
		})
		if err != nil {
			fmt.Println("Prompt failed:", err)
			os.Exit(1)
		}

		commitMessage := CreateCommitMessage(commitType, commitScope, commitDescription)

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
func CreateCommitMessage(commitType, commitScope, commitDescription string) string {
	commitMessage := commitType
	if commitScope != "" {
		commitMessage += "(" + commitScope + ")"
	}
	commitMessage += ": " + commitDescription
	return commitMessage
}

// CreateGitCommit creates a Git commit with the given message and files.
func CreateGitCommit(message string, files []string) error {
	commitArgs := append([]string{"commit", "-m", message}, files...)
	commitGitCmd := exec.Command("git", commitArgs...)
	commitGitCmd.Stdout = os.Stdout
	commitGitCmd.Stderr = os.Stderr

	return commitGitCmd.Run()
}
