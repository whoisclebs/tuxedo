# Contributing to Tuxedo

Thank you for your interest in contributing to **Tuxedo**! We welcome all contributions, whether it's reporting issues, suggesting new features, improving documentation, or submitting code changes. This document will guide you on how to get started.

## How to Contribute

### 1. Reporting Issues
If you find a bug, inconsistency, or have a suggestion, please open an issue:
- Clearly describe the problem or suggestion.
- Provide steps to reproduce the issue (if applicable).
- Include relevant logs, screenshots, or error messages.

### 2. Fork & Clone the Repository
To contribute code, start by forking the repository and cloning it to your local machine:
```sh
# Fork the repository on GitHub, then clone it locally
$ git clone https://github.com/whoisclebs/tuxedo.git
$ cd tuxedo
```

### 3. Create a Feature Branch
Before making changes, create a new branch:
```sh
$ git checkout -b feature/your-feature-name
```

### 4. Writing Code
- Follow Go best practices and keep the code clean and maintainable.
- Try to keep changes small and focused on a single feature or fix.
- Ensure your code is well-documented and adheres to the project structure.
- Follow the **Conventional Commits** specification for commit messages.
- Use meaningful variable and function names.
- Keep functions and methods small and modular.

### 5. Testing
Before submitting changes, make sure the tests pass. All tests are located in the root of the project:
```sh
$ go test ./...
```
If adding new features, include relevant test cases to maintain high coverage.

### 6. Commit and Push Changes
Follow the **Conventional Commits** format when committing your changes:
```sh
$ git commit -m "feat: add feature XYZ to improve HTTP handling"
$ git push origin feature/your-feature-name
```
Example commit types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation updates
- `test`: Adding or updating tests
- `refactor`: Code refactoring without changing functionality
- `chore`: Maintenance tasks, like build process updates

### 7. Open a Pull Request (PR)
Once your code is ready:
- Go to the [Tuxedo GitHub repository](https://github.com/whoisclebs/tuxedo)
- Click **New Pull Request** and select your branch
- Provide a clear description of the changes and why they are necessary
- Ensure your PR follows the project's coding standards and guidelines

### 8. Code Review & Merging
Your PR will be reviewed by maintainers. If requested, make necessary changes and update your PR. Once approved, it will be merged into the main branch.

## Best Practices for Open Source Contributions
- Keep PRs small and focused.
- Follow Go naming conventions and maintain consistent formatting.
- Write meaningful commit messages.
- Ensure all tests pass before submitting PRs.
- Document any changes in the **README** or relevant documentation files.
- Use **error handling** properly to avoid unexpected crashes.
- Avoid unnecessary dependencies and keep the library lightweight.
- Be respectful and collaborative in discussions.

## Getting Help
If you need any help, feel free to ask in the issues section or reach out to the maintainers.

Thank you for contributing to **Tuxedo**! 🚀
