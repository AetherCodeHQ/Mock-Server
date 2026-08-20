# Mock Server

![CI](https://github.com/Qyroxen/Mock-Server/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Mock-Server/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Mock-Server?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Mock-Server)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Mock-Server)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Mock-Server?style=social)](https://github.com/Qyroxen/Mock-Server/stargazers)

## What is it?

Mock Server is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Mock-Server.git
cd Mock-Server
go build -o mockserver .

# Run
./mockserver --help
```

## CLI Usage

```bash
# Basic usage
./mockserver

# With flags
./mockserver --verbose --output json

# Get help
./mockserver --help
```

## Examples

```bash
# Example 1
./mockserver example1

# Example 2
./mockserver example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o mockserver .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Mock-Server/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Mock-Server?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Mock-Server/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Mock-Server?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Mock-Server/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Mock-Server" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Mock-Server/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Mock-Server" alt="Pull Requests">
  </a>
</p>
