# Go Interface & Enum Explorer

A command-line educational tool for teaching Golang interfaces and enums through progressive examples.

## Overview

Go Interface & Enum Explorer is an interactive command-line application designed to help developers learn about two important concepts in Go:

1. **Interfaces** - Go's powerful way to define behavior and enable polymorphism
2. **Enums** - Idiomatic patterns for implementing enumerated types in Go

The tool provides a structured, progressive learning path with explanations, code examples, and outputs to help you understand these concepts from basics to advanced usage.

## Features

- **Progressive Tutorial Mode**: Learn concepts in a logical, step-by-step order
- **Example Browser**: Directly access specific topics you're interested in
- **Interactive Learning**: See explanations, code examples, and their output together
- **Comprehensive Coverage**: From basic interface definitions to advanced enum patterns

## Topics Covered

### Interfaces

- Basic Interfaces: Learn the fundamental concept of interfaces in Go
- Interface Implementation: How types implement interfaces implicitly
- Empty Interface: Understanding the empty interface (`interface{}`) and its uses
- Type Assertion: Extract concrete types from interfaces safely
- Interface Composition: Build complex interfaces from simpler ones

### Enums

- Basic Enums: Simple constants as enumerated values
- Iota Enums: Creating sequential constants with `iota`
- String Enums: Adding string representation to enum values
- Behavior Enums: Attaching methods to enum types for rich behavior

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/go-interface-enum-explorer.git
   cd go-interface-enum-explorer
   ```

2. Build the application:
   ```
   go build -o go-explorer
   ```

3. Run the application:
   ```
   ./go-explorer
   ```

## Usage

After starting the application, you'll see a main menu with the following options:

1. **Start Tutorial**: Begin a guided journey through all concepts in a progressive order
2. **Browse Examples**: Pick specific topics you're interested in exploring
3. **Help**: View information about how to use the tool and learn about Go interfaces and enums
4. **Quit**: Exit the application

Navigate through the application using the on-screen prompts.

## Learning Path

For the best learning experience, we recommend starting with the "Tutorial" mode, which guides you through the concepts in this order:

1. Basic Interfaces
2. Interface Implementation
3. Empty Interface
4. Type Assertion
5. Interface Composition
6. Basic Enums
7. Iota Enums
8. String Enums
9. Behavior Enums

This progression builds on previous concepts to provide a comprehensive understanding of interfaces and enums in Go.

## Requirements

- Go 1.11 or later

## Contributing

Contributions are welcome! If you'd like to add more examples, improve explanations, or fix issues, please feel free to submit a pull request.

### GitHub Actions

This project includes GitHub Actions workflows for continuous integration:

1. **Build Workflow**: Runs on every push to main/master and pull requests, ensuring the code builds correctly.
   - Runs `go build`, `go test`, and code quality checks
   - Automatically uses the Go version specified in go.mod

2. **Release Workflow**: Triggered when a tag starting with 'v' is pushed.
   - Builds binaries for Linux, macOS, and Windows
   - Creates a GitHub Release with the binaries attached
   - Sets version information in the binaries
   - Automatically uses the Go version specified in go.mod

### Creating a Release

To create a new release:

```bash
git tag v1.0.0
git push origin v1.0.0
```

This will trigger the release workflow, which builds binaries and creates a GitHub Release.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
