/*
Package commit provides functionality for creating Git commits.

The CommitInfo struct in this package represents information needed for creating a Git commit. It includes fields for specifying the commit type, scope, description, body, co-authors, and breaking change details.

Usage:
  - Create instances of the CommitInfo struct to define commit information.
  - Set the CommitType, CommitDescription, and other fields to configure the commit.
  - Optionally, add co-authors to the CoAuthors field if the commit is co-authored.
  - Use the IsBreakingChange field to indicate whether the commit introduces a breaking change.
  - For more information on how to use the CommitInfo struct and create Git commits, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package commit

// CommitInfo represents information needed for creating a Git commit.
type CommitInfo struct {
	CommitType                string
	CommitScope               string
	CommitDescription         string
	CommitBody                string
	IsCoAuthored              bool
	CoAuthors                 []string
	IsBreakingChange          bool
	BreakingChangeDescription string
}
