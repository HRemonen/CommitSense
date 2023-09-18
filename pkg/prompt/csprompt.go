/*
Package csprompt provides a struct for defining promptui prompts in CommitSense.

The CSPrompt struct in this package represents a prompt object that can be used with the promptui library. It includes fields for the prompt's label, validation function, and default value.

Usage:
  - Create a CSPrompt object to define custom prompts for user input.
  - Use the CSPrompt object in your CommitSense commands for interactive prompts.

For more information on how to use prompts with CommitSense, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package csprompt

import "commitsense/pkg/item"

// CSPrompt represents a promptui prompt object used for user input.
type CSPrompt struct {
	Label     string
	Items     []*item.Item
	CursorPos int
	Validate  func(string) error
	Default   string
}
