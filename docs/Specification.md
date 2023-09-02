# CommitSense Specification

CommitSense follows the Conventional Commits specification to create meaningful and standardized commit messages for version control. This specification outlines the format and conventions for CommitSense commit messages.

## Format

A CommitSense commit message consists of a **header**, an optional **body**, and an optional **footer**, with each part separated by a blank line. Here's the format:

```bash
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Header

- **Type (Required):** Describes the purpose of the commit. It should be one of the following:

  - `feat`: A new feature or enhancement.
  - `fix`: A bug fix.
  - `chore`: Routine tasks, maintenance, or refactoring.
  - `docs`: Documentation updates.
  - `style`: Code style and formatting changes (not affecting functionality).
  - `refactor`: Code refactorings without adding new features or fixing bugs.
  - `perf`: Performance improvements.
  - `test`: Adding or modifying tests.
  - `build`: Build-related changes (e.g., dependencies, configurations).
  - `ci`: Continuous integration and deployment changes.

- **Scope (Optional):** Describes the part of the codebase affected by the commit.

- **Description (Required):** A brief and concise description of the change in the present tense. Start with a capital letter and do not end with a period.

### Body

- **Optional:** A more detailed explanation of the commit. Use this section when the header alone doesn't provide enough context.

### Footer

- **Optional:** Additional metadata about the commit, such as references to issue trackers, breaking changes, or related commits.

## Benefits

- **Clarity:**  Conventional Commits make it easier to understand the purpose and impact of each commit.
- **Automation:** Tools and scripts can parse commit messages to automate tasks like versioning and generating changelogs.
- **Consistency:** Enforces a consistent commit message format across a project, making it more maintainable.

## References

- [Conventional Commits Website](https://www.conventionalcommits.org/)
- [Conventional Commits GitHub Repository](https://github.com/conventional-commits/conventionalcommits.org)

By following this Conventional Commits-based specification, CommitSense ensures that its commit messages are meaningful, standardized, and aligned with best practices in version control.
