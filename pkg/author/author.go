package author

import (
	"os/exec"
	"strings"
)

func GetSuggestedCoAuthors() ([]string, error) {
	// Use the `git rev-list` command to obtain a list of authors who have made commits in the Git repository.
	cmd := exec.Command("git", "log", "--format='%aN <%aE>'", "--all", "--no-merges", "|", "sort", "--unique")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// Parse the output to extract author names and email addresses.
	lines := strings.Split(string(output), "\n")
	var suggestedCoAuthors []string
	for _, line := range lines {
		if strings.HasPrefix(line, "'") && strings.HasSuffix(line, "'") {
			suggestedCoAuthors = append(suggestedCoAuthors, strings.Trim(line, "'"))
		}
	}

	return suggestedCoAuthors, nil
}
