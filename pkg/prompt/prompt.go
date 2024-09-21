/*
Package prompt provides a struct for defining promptui prompts in CommitSense.

The prompt struct in this package represents a prompt object that can be used with the promptui library. It includes fields for the prompt's label, validation function, and default value.

Usage:
  - Create a prompt object to define custom prompts for user input.
  - Use the prompt object in your CommitSense commands for interactive prompts.

For more information on how to use prompts with CommitSense, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package prompt

import (
	"commitsense/pkg/author"
	"commitsense/pkg/config"
	"fmt"
	"os"
	"strings"


	goprompt "github.com/c-bata/go-prompt"
	"github.com/manifoldco/promptui"
)

// Item represents an item with an ID referring to a certain item in a multiselect prompt
type Item struct {
	ID         string
	IsSelected bool
}

// PromptCommitType prompts the user to select a commit type.
func PromptCommitType(label string) (string, error) {
	cfg, _ := config.ReadConfigFile()

	promptType := promptui.Select{
		Label: label,
		Items: cfg.CommitTypes,
	}

	_, typeResult, err := promptType.Run()

	return typeResult, err
}

// PromptForString prompts the user to enter a string.
func PromptForString(label string, validator promptui.ValidateFunc) (string, error) {
	promptString := promptui.Prompt{
		Label:    label,
		Validate: validator,
	}
	return promptString.Run()
}

// PromptForMultilineString prompts the user for a multiline string input based on the provided prompt configuration.
// Users can enter multiple lines of text until they press Enter twice to finish.
func PromptForMultilineString(label string) (string, error) {
	var lines []string
	for {
		line, err := PromptForString(label, nil)
		if err != nil || line == "" {
			break
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n"), nil
}

// PromptForCoAuthors displays a prompt to enter co-author names for a Git commit.
//
// This function provides real-time auto-completion suggestions based on the suggestedCoAuthors
// list. Users can choose from the suggestions or enter custom co-authors. It returns a slice
// of selected co-author names.
func PromptForCoAuthors(label string) ([]string, error) {
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
		goprompt.OptionPrefix(label),
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

func createSelectTemplates() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "→ {{if .IsSelected}}✔ {{end}} {{ .ID | cyan }}",
		Inactive: "{{if .IsSelected }}✔ {{ .ID | green }} {{else}}{{ .ID | faint }}{{end}} ",
	}
}

// Prepend the prompt.Items with the continue item
func prependItemsWithSpecialOptions(items []*Item) []*Item {
	continueItem := &Item{ID: "Continue"}

	items = append([]*Item{continueItem}, items...)

	return items
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
