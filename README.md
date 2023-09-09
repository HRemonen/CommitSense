# CommitSense

CommitSense is a command-line tool that simplifies Git version control by providing an interactive and standardized way to stage files and create commit messages following the Conventional Commits specification.

## Features

- Interactive file selection for staging.
- Conventional Commits-based commit message generation.
- Improved commit message consistency.
- Streamlined Git workflow.

Here's the cool part: CommitSense plays well with native Git commands under the hood. So, while you're using Git commands like a console wizard, CommitSense is right there, ensuring compatibility and helping you create those commits with style when you're ready!

But, there's one thing to note: CommitSense doesn't support chunking files when adding. So, for those complex file-staging tasks, you might want to stick to the classic `git add` method.

## Usage

### Adding Files

```bash
./commitsense add
```

This command launches an interactive interface for selecting files to stage.

### Creating Commits

```bash
./commitsense commit
```

This command guides you through creating a commit message according to the Conventional Commits format.

## Development

### golangci-lint

[golangci-lint](https://golangci-lint.run/) is a fast and customizable Go linter. It provides a wide range of checks for various aspects of your Go code.

To install golangci-lint, run the following command:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Running golangci-lint

To run golangci-lint on your project, navigate to your project's root directory and execute:

```bash
golangci-lint run
```
This command will analyze your Go code, check for issues, and display the results.

To run and fix autofixable problems, run the following command:

```bash
golangci-lint run --fix
```

### gofumpt

gofumpt is a stricter Go code formatter that follows the [gofumpt style](https://github.com/mvdan/gofumpt).

To install gofumpt, run:

```bash
go install mvdan.cc/gofumpt@latest
```

### Running gofumpt

To format and organize your import statements, run:

```bash
gofumpt -l -w .
```

## Contributing
We welcome contributions to CommitSense! To contribute, follow these steps:

- Fork this repository.
- Create a new branch for your feature or bug fix: git checkout -b my-feature.
- Commit your changes following the Conventional Commits format.
- Push your branch to your fork: git push origin my-feature.
- Open a pull request to the main repository.
- Please ensure your code follows best practices and includes appropriate tests.

## License

CommitSense is released under the MIT License. See [LICENSE](LICENSE) for details.

## Acknowledgments
This project is inspired by the Conventional Commits specification. Learn more at [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
