package author

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func extractNameAndEmail(line string) []string {
	pattern := `(\w[\w\s]+)\s+<([^>]+)>`
	re := regexp.MustCompile(pattern)
	return re.FindStringSubmatch(line)
}

func GetSuggestedCoAuthors() ([]string, error) {
	// Use the `git rev-list` command to obtain a list of authors who have made commits in the Git repository.
	revlist := "git log --pretty='%an <%ae>' | sort -u"
	cmd := exec.Command("bash", "-c", revlist)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

	authorString := strings.TrimSpace(string(output))

	strings.Replace(authorString, `\n`, "\n", -1)

	// Parse the output to extract author names and email addresses.
	suggestedCoAuthors := strings.Split(authorString, "\n")

	fmt.Println("suggested authors", suggestedCoAuthors)

	return suggestedCoAuthors, nil
}
