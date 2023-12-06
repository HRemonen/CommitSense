/*
Package item provides a struct for representing prompt items.

For more information on how to use the Item struct and its associated functionality, refer to the package-specific functions and commands.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package prompt

// Item represents an item with an ID referring to a certain item in a multiselect prompt
type Item struct {
	ID         string
	IsSelected bool
}
