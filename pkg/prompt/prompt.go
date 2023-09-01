package prompt

type Prompt struct {
	Label    string
	Validate func(string) error
	Default  string
}
