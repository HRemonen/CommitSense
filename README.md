# CommitSense

CommitSense is a command-line tool that simplifies Git version control by providing an interactive and standardized way to stage files and create commit messages following the Conventional Commits specification.

## Features

- Interactive file selection for staging.
- Conventional Commits-based commit message generation.
- Improved commit message consistency.
- Streamlined Git workflow.

## 🎩 Neckbeard Mode

If you've got a neckbeard or you've been using the console since before the dinosaurs were alive, we've got you covered! You can totally go old school and add files for staging using the `git add <filename>` command. Some even say it's faster than an interactive CLI.

Here's the cool part: CommitSense plays well with native Git commands under the hood. So, while you're using Git commands like a console wizard, CommitSense is right there, ensuring compatibility and helping you create those commits with style when you're ready!

But, there's one thing to note: CommitSense doesn't support chunking files when adding. So, for those complex file-staging tasks, you might want to stick to the classic `git add` method.

## Development

## Linters and Code Formatting

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

## Pre-Commit Hooks

This project utilizes pre-commit hooks to ensure code quality and consistency before commits are made. Pre-commit hooks are automated checks that run locally on your development environment before you commit changes to the repository. They help maintain code quality and consistency across the project.

### Available Pre-Commit Hooks

We use the following pre-commit hooks from various repositories to check and format our code:

#### [pre-commit/pre-commit-hooks](https://github.com/pre-commit/pre-commit-hooks)

1. **check-added-large-files**: Prevents large files from being added to the repository.
2. **trailing-whitespace**: Ensures there are no trailing whitespace characters at the end of lines.
3. **end-of-file-fixer**: Appends a newline to the end of files if it's missing.
4. **check-yaml**: Checks YAML files for syntax errors.
5. **check-json**: Checks JSON files for syntax errors.

#### [Bahjat/pre-commit-golang](https://github.com/Bahjat/pre-commit-golang)

1. **gofumpt**: Enforces the gofumpt code style for Go code.
2. **golangci-lint**: Runs linting checks on Go code using golangci-lint.
3. **go-unit-tests**: Runs Go unit tests to ensure code correctness.

### How to Use Pre-Commit Hooks


1. **Install Pre-Commit**: If you haven't already installed pre-commit, you can do so by following the [installation guide](https://pre-commit.com/#install) on the Pre-Commit website.

2. **Install Pre-Commit Hooks Locally**: Run the following command to install the pre-commit hooks locally in your project:

    ```bash
    pre-commit install
    ```

3. **Run Pre-Commit**: Before committing your changes, run the pre-commit hooks using:

    ```bash
    pre-commit run --all-files
    ```

    This command will run all configured hooks on your changes. If any hooks fail, they will provide feedback on what needs to be fixed.

4. **Commit with Confidence**: Once all hooks pass without errors, you can commit your changes knowing that they meet the code quality and style standards defined in the pre-commit hooks.

**Note for Contributors**: To maintain consistency and adhere to project standards, please do not modify the .pre-commit-config.yaml file. Changes to this file are not to be made if you want to contribute to the project.


## Installation

To install CommitSense, you'll need [Go](https://golang.org/) installed on your system.

1. Clone this repository:

    ```bash
    git clone https://github.com/yourusername/CommitSense.git
    ```

2. Build the binary:
    ```bash
    cd CommitSense
    go build
    ```

3. Run CommitSense
    ```bash
    ./commitsense
    ```

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
