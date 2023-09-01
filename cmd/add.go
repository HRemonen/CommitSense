/*
Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package cmd

import (
	"commitsense/pkg/item"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// addCmd represents the add command used to stage files for commit
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Interactively select files to stage",
	Long: `Select files that are not tracked or files that are modified and add the selected files to staging

This command uses the native git add and git status commands under 
the hood. Use this command to stage selected files before making a 
commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := getChangedFiles()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		selectedFilePtrs, err := promptForFiles(0, files)

		for _, file := range selectedFilePtrs {
			stageFile(file.ID)
		}
	},
}

func init() {
	// Add the add-interactive subcommand to the root command
	rootCmd.AddCommand(addCmd)
}

func getChangedFiles() ([]*item.Item, error) {
	// Different git porcelain status codes for files
	gitPrefixes := []string{"M", "A", "D", "??"}

	// Simulate getting the list of changed files from Git
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var changedFiles []*item.Item
	for _, line := range lines {
		// Strip leading and trailing whitespace
		line = strings.TrimSpace(line)

		if containsPrefix(line, gitPrefixes) {
			// Extract the file path
			parts := strings.Fields(line)
			if len(parts) == 2 {
				var items = []*item.Item{
					{
						ID: parts[1],
					},
				}
				changedFiles = append(changedFiles, items...)
			}
		}
	}
	return changedFiles, nil
}

func promptForFiles(selectedPos int, allItems []*item.Item) ([]*item.Item, error) {
	const continueItem = "Continue"

	if len(allItems) > 0 && allItems[0].ID != continueItem {
		var items = []*item.Item{
			{
				ID: continueItem,
			},
		}

		allItems = append(items, allItems...)
	}

	// Define promptui template
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "→ {{if .IsSelected}}✔ {{end}} {{ .ID | cyan }}",
		Inactive: "{{if .IsSelected }}✔ {{ .ID | green }} {{else}}{{ .ID | faint }}{{end}} ",
	}

	prompt := promptui.Select{
		Label:        "Select files to stage",
		Items:        allItems,
		Templates:    templates,
		Size:         10,
		CursorPos:    selectedPos,
		HideSelected: true,
	}

	selectionIdx, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("prompt failed: %w", err)
	}

	chosenItem := allItems[selectionIdx]

	if chosenItem.ID != "Continue" {
		// If the user selected something other than "Continue",
		// toggle selection on this item and run the function again.
		chosenItem.IsSelected = !chosenItem.IsSelected

		return promptForFiles(selectionIdx, allItems)
	}

	var selectedItems []*item.Item
	for _, i := range allItems {
		if i.IsSelected {
			selectedItems = append(selectedItems, i)
		}
	}
	return selectedItems, nil
}

func stageFile(filename string) {
	// Simulate staging the file using Git add
	cmd := exec.Command("git", "add", filename)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error staging file %s: %v\n", filename, err)
	}
}

func containsPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
