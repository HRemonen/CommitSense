/*
Package commit provides functionality for creating Git commits.

This file includes utility functions for interacting with the user.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package commit

import (
	"commitsense/pkg/author"
	"commitsense/pkg/config"
	"fmt"
	"os"
	"strings"

	csprompt "commitsense/pkg/prompt"

	goprompt "github.com/c-bata/go-prompt"
	"github.com/manifoldco/promptui"
)

// PromptCommitType prompts the user to select a commit type.
func PromptCommitType(prompt csprompt.Prompt) (string, error) {
	config, _ := config.ReadConfigFile()

	promptType := promptui.Select{
		Label: prompt.Label,
		Items: config.CommitTypes,
	}

	_, typeResult, err := promptType.Run()

	return typeResult, err
}

// PromptForBool prompts the user to enter a boolean value.
func PromptForBool(prompt csprompt.Prompt) (bool, error) {
	promptBool := promptui.Prompt{
		Label:    prompt.Label,
		Validate: prompt.Validate,
		Default:  prompt.Default,
	}

	result, err := promptBool.Run()
	if err != nil {
		return false, err
	}

	return result == "Y" || result == "y", nil
}

// PromptForString prompts the user to enter a string.
func PromptForString(prompt csprompt.Prompt) (string, error) {
	promptString := promptui.Prompt{
		Label:    prompt.Label,
		Validate: prompt.Validate,
		Default:  prompt.Default,
	}
	return promptString.Run()
}

// PromptForMultilineString prompts the user for a multiline string input based on the provided prompt configuration.
// Users can enter multiple lines of text until they press Enter twice to finish.
func PromptForMultilineString(prompt csprompt.Prompt) (string, error) {
	var lines []string
	for {
		line, err := PromptForString(prompt)
		if err != nil || line == "" {
			break
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n"), nil
}

// Prepend the prompt.Items with the continue item
func prependItemsWithSpecialOptions(items []*csprompt.Item) []*csprompt.Item {
	continueItem := &csprompt.Item{ID: "Continue"}

	items = append([]*csprompt.Item{continueItem}, items...)

	return items
}

func createSelectTemplates() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "→ {{if .IsSelected}}✔ {{end}} {{ .ID | cyan }}",
		Inactive: "{{if .IsSelected }}✔ {{ .ID | green }} {{else}}{{ .ID | faint }}{{end}} ",
	}
}

func promptForMultipleItems(prompt csprompt.Prompt) ([]*csprompt.Item, error) {
	promptItems := prependItemsWithSpecialOptions(prompt.Items)

	promptMultiple := promptui.Select{
		Label:        prompt.Label,
		Items:        promptItems,
		Templates:    createSelectTemplates(),
		Size:         10,
		CursorPos:    prompt.CursorPos,
		HideSelected: true,
	}

	var selectedItems []*csprompt.Item

	for {
		selectionIdx, _, err := promptMultiple.Run()
		if err != nil {
			return nil, fmt.Errorf("prompt failed: %w", err)
		}

		chosenItem := prompt.Items[selectionIdx]

		if chosenItem.ID == "Continue" {
			break
		}

		// If the user selected something other than "Continue",
		// toggle selection on this item.
		chosenItem.IsSelected = !chosenItem.IsSelected
		prompt.CursorPos = selectionIdx
	}

	// Collect selected items.
	for _, item := range prompt.Items {
		if item.IsSelected {
			selectedItems = append(selectedItems, item)
		}
	}

	return selectedItems, nil
}

func coAuthorCompleter(suggestedCoAuthors []string) goprompt.Completer { // Use "p" as the alias
	return func(d goprompt.Document) []goprompt.Suggest {
		suggestions := []goprompt.Suggest{}
		text := d.TextBeforeCursor()
		for _, coAuthor := range suggestedCoAuthors {
			if strings.HasPrefix(coAuthor, text) {
				suggestions = append(suggestions, goprompt.Suggest{Text: coAuthor})
			}
		}
		return goprompt.FilterHasPrefix(suggestions, text, true)
	}
}

// PromptForCoAuthors displays a prompt to enter co-author names for a Git commit.
//
// This function provides real-time auto-completion suggestions based on the suggestedCoAuthors
// list. Users can choose from the suggestions or enter custom co-authors. It returns a slice
// of selected co-author names.
func PromptForCoAuthors(prompt csprompt.Prompt) ([]string, error) {
	suggestedCoAuthors, err := author.GetSuggestedCoAuthors()
	if err != nil {
		fmt.Println("Error getting the suggested co-authors:", err)
		os.Exit(1)
	}

	fmt.Println("Enter Co-authors:")
	fmt.Println("Press 'Tab' to auto-complete.")

	pr := goprompt.New(
		func(_ string) { /* No-op executor */ },
		coAuthorCompleter(suggestedCoAuthors),
		goprompt.OptionPrefix(prompt.Label),
	)

	coAuthors := []string{}
	for {
		coAuthor := pr.Input()
		if coAuthor == "" {
			break
		}
		coAuthors = append(coAuthors, coAuthor)
	}

	return coAuthors, nil
}
