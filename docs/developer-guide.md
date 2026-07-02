# Developer Guide

## Welcome

Welcome to **Commerce AI Platform**!

This guide explains how to set up the project for local development, build it, run it, execute tests, and validate that everything is working correctly.

> **Note:** This project is under active development. Additional setup instructions for Ollama, RAG, OpenTelemetry, Phoenix, and evaluations will be added as those features are implemented.

---

# Prerequisites

Install the following software before getting started:

* Git
* Go 1.25 or later
* Visual Studio Code (recommended)

Future milestones will also require:

* Ollama
* Docker Desktop
* Arize Phoenix

---

# Clone the Repository

```bash
git clone https://github.com/mayankanup/commerce-ai-platform.git

cd commerce-ai-platform
```

---

# Install Dependencies

Download all Go module dependencies.

```bash
go mod tidy
```

---

# Project Structure

```
commerce-ai-platform/

cmd/                Application entry points
config/             Configuration files
docs/               Project documentation
internal/           Application source code
data/               Sample data
test/               Integration and end-to-end tests
```

---

# Build the Project

```bash
go build ./cmd/api
```

Or create an executable:

```bash
go build -o bin/commerce-ai-platform ./cmd/api
```

---

# Run the Application

```bash
go run ./cmd/api
```

The application starts on:

```
http://localhost:8080
```

---

# Validate the Application

Open a browser and navigate to:

```
http://localhost:8080/
```

A successful response should indicate that the service is running.

Health endpoint support will be added in a future milestone.

---

# Run Unit Tests

Execute all unit tests:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./internal/platform/config
```

---

# Format the Code

```bash
go fmt ./...
```

---

# Run Static Analysis

```bash
go vet ./...
```

---

# Recommended Development Workflow

1. Pull the latest changes from `main`.
2. Create a feature branch.
3. Implement your changes.
4. Run formatting and tests.
5. Commit using Conventional Commits.
6. Submit a Pull Request.

Example:

```
git checkout -b feature/CAP-002-config-loader
```

---

# Commit Message Convention

Examples:

```
feat(config): add YAML configuration loader

fix(api): handle invalid requests

docs: update developer guide

test(agent): add tool execution tests
```

---

# Troubleshooting

## Missing Go dependencies

Run:

```bash
go mod tidy
```

---

## Application fails to start

Verify that:

* Go is installed correctly.
* No other application is using port `8080`.
* `config/config.yaml` exists.

---

## Tests fail

Run:

```bash
go test ./...
```

Review the error messages and ensure your local environment is up to date.

---

# Additional Documentation

For more information, see:

* `docs/architecture.md`
* `docs/roadmap.md`
* `docs/backlog.md`

Additional guides for RAG, observability, evaluations, and deployment will be added in future milestones.
