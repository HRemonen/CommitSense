/*
Package prompt provides a struct for defining promptui prompts in CommitSense.

The Prompt struct in this package represents a prompt object that can be used with the promptui library. It includes fields for the prompt's label, validation function, and default value.

Usage:
  - Create a Prompt object to define custom prompts for user input.
  - Use the Prompt object in your CommitSense commands for interactive prompts.

For more information on how to use prompts with CommitSense, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package prompt

// Prompt represents a promptui prompt object used for user input.
type Prompt struct {
	Label    string
	Validate func(string) error
	Default  string
}
