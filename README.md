# CommitSense

![CommitSense Logo](commit_sense_logo.png)

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