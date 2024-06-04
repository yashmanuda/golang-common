[![codecov](https://codecov.io/gh/yashmanuda/golang-common/graph/badge.svg?token=LF0C83KVFL)](https://codecov.io/gh/yashmanuda/golang-common)

# Go Project

Welcome to our Go project repository, designed to facilitate efficient configuration management, logging, and middleware operations. This project leverages environment-based configurations and provides robust logging functionalities.

## Features

- **Environment-Based Configuration**: Supports dynamic configuration using environment variables, with "\_\_" used as both separator and delimiter.
- **Logging**: Streamlined logging capabilities that aid in monitoring and debugging the application.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation

Clone the repository and install the required dependencies:

```bash
git clone <repository-url>
cd <repository-directory>
go mod tidy
```

### Running Tests

Verify the setup and functionality by running the tests:

```bash
go test ./... -v -cover -coverpkg=./...
```
