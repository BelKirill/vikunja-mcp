

# Contributing to Vikunja MCP

Thank you for considering contributing to the Vikunja MCP project! We welcome bug reports, feature requests, and pull requests. To make the process smooth for everyone, please follow these guidelines.

## Code of Conduct

Please read and adhere to our [Code of Conduct](CODE_OF_CONDUCT.md). Be respectful, inclusive, and constructive.

## Getting Started

1. **Fork** the repository on GitHub.
2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/BelKirill/vikunja-mcp.git
   cd vikunja-mcp
   ```
3. **Install prerequisites**:
   - Go 1.20+ (with `GOPATH` and `GOROOT` configured)
   - [golangci-lint](https://github.com/golangci/golangci-lint) v2.x (`make lint` will use it)
   - [swag](https://github.com/swaggo/swag) for API docs generation (if modifying API)
   - Docker & Docker Compose (for local Vikunja integration, optional)

4. **Set up dependencies**:
   ```bash
   go mod download
   ```

## Development Workflow

### Branching

- Create a descriptive branch name:
  - Feature: `feature/<short-description>`
  - Bugfix: `bugfix/<short-description>`
- Branch off from `main`:
  ```bash
  git checkout main
  git pull upstream main
  git checkout -b feature/awesome-endpoint
  ```

### Coding Standards

- Follow Go idioms and conventions.
- Keep functions small and focused.
- Write clear, concise comments for exported types and functions.
- Use [dependency injection](https://github.com/golang/go/wiki/CodeReviewComments#dependency-injection) where appropriate.

### Formatting & Linting

- Run formatting before committing:
  ```bash
  go fmt ./...
  ```
- Run lint checks:
  ```bash
  make lint
  ```
- Fix any reported issues. The CI pipeline enforces these checks.

### Testing

- Write unit tests alongside your code (`*_test.go`).
- Run tests locally:
  ```bash
  go test ./... -v
  ```
- Aim for meaningful coverage. We enforce a minimum threshold in CI.

### API Documentation

- If you update or add endpoints, regenerate OpenAPI docs:
  ```bash
  swag init --output ./docs
  ```
- Commit changes to the generated `docs` directory.

## Pull Requests

1. Push your branch to your fork:
   ```bash
   git push origin feature/awesome-endpoint
   ```
2. Open a Pull Request against the `main` branch.
3. Ensure the pipeline passes all checks (lint, tests, build).
4. Fill out the PR template:
   - **Description**: What and why.
   - **Related issue**: Link to an existing issue or create one.
   - **Checklist**:
     - [ ] Code compiles and tests pass.
     - [ ] Relevant documentation updated.
     - [ ] API docs regenerated (if applicable).
5. Request reviews from project maintainers.

## Reporting Issues

- Search existing issues before creating a new one.
- Use a clear, descriptive title.
- Provide steps to reproduce, expected vs. actual behavior, and any logs or screenshots.

## Commit Message Guidelines

Follow [Conventional Commits](https://www.conventionalcommits.org/):
```
<type>(<scope>): <subject>
```
- **type**: `feat`, `fix`, `docs`, `chore`, `refactor`, etc.
- **scope**: optional module or component.
- **subject**: short summary in imperative mood.

Example:
```
feat(api): add support for /healthz endpoint
```

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).

---
_Thank you for making Vikunja MCP even better!_