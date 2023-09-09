/*
Package author provides functions for working with Git commit authors and co-authors.

This package includes functions for retrieving suggested co-authors who have made commits in the Git repository.
It uses the `git log` command to extract author names and email addresses from the commit history. The resulting list
of authors can be used when creating Git commits with co-authors.

Usage:
  - Call the GetSuggestedCoAuthors function to obtain a list of suggested co-authors based on the Git commit history.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package author

import (
	"os/exec"
	"strings"
)

// GetSuggestedCoAuthors retrieves a list of suggested co-authors who have made commits in the Git repository.
//
// This function uses the `git log` command to obtain a list of authors who have made commits in the Git repository.
// It executes the command and processes the output to extract author names and email addresses. The resulting
// list represents suggested co-authors for Git commits.
func GetSuggestedCoAuthors() ([]string, error) {
	// Use the `git rev-list` command to obtain a list of authors who have made commits in the Git repository.
	revlist := "git log --pretty='%an <%ae>' | sort -u"
	cmd := exec.Command("bash", "-c", revlist)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	authorString := strings.TrimSpace(string(output))

	authorString = strings.ReplaceAll(authorString, `\n`, "\n")

	// Parse the output to extract author names and email addresses.
	suggestedCoAuthors := strings.Split(authorString, "\n")

	return suggestedCoAuthors, nil
}
