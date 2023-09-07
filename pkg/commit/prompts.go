/*
Package commit provides functionality for creating Git commits.

This file includes utility functions for interacting with the user.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package commit

import (
	"commitsense/pkg/author"
	"commitsense/pkg/prompt"
	"strings"

	"github.com/manifoldco/promptui"
)

// PromptCommitType prompts the user to select a commit type.
func PromptCommitType() (string, error) {
	promptType := promptui.Select{
		Label: "Select a commit type",
		Items: []string{"feat", "fix", "chore", "docs", "style", "refactor", "perf", "test", "build", "ci"},
	}
	_, typeResult, err := promptType.Run()
	return typeResult, err
}

// PromptForBool prompts the user to enter a boolean value.
func PromptForBool(prompt prompt.Prompt) (bool, error) {
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
func PromptForString(prompt prompt.Prompt) (string, error) {
	promptString := promptui.Prompt{
		Label:    prompt.Label,
		Validate: prompt.Validate,
		Default:  prompt.Default,
	}
	return promptString.Run()
}

// PromptForMultilineString prompts the user for a multiline string input based on the provided prompt configuration.
// Users can enter multiple lines of text until they press Enter twice to finish.
func PromptForMultilineString(prompt prompt.Prompt) (string, error) {
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

// PromptForCoAuthors displays a prompt to select or enter co-authors for a Git commit.
//
// This function retrieves a list of suggested co-authors using the GetSuggestedCoAuthors function
// from the author package. It then presents the user with a selectable list of suggested co-authors
// and allows them to choose from the suggestions or enter custom co-authors.
func PromptForCoAuthors(prompt prompt.Prompt) ([]string, error) {
	suggestedCoAuthors, err := author.GetSuggestedCoAuthors()
	if err != nil {
		return nil, err
	}

	// Present the user with a list of suggested co-authors and allow them to select or enter custom co-authors.
	promptCoAuthors := promptui.Select{
		Label: prompt.Label,
		Items: suggestedCoAuthors,
	}

	_, authorResult, err := promptCoAuthors.Run()
	return []string{authorResult}, err
}
