/*
Package item provides a struct for representing files to be staged in a Git repository.

The Item struct in this package represents an item that corresponds to a file in a Git repository. It includes fields for the file's ID (referring to the file path) and a boolean flag IsSelected to track whether the file is selected to be staged.

Usage:
  - Create instances of the Item struct to represent individual files.
  - Use the IsSelected field to keep track of which files are selected for staging.

For more information on how to use the Item struct and its associated functionality, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package item

// Item represents an item with an ID referring to a file path in a Git repository
// and an IsSelected field to track its staging status.
type Item struct {
	ID         string
	IsSelected bool
}
