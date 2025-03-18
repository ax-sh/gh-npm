# gh-npm

A GitHub CLI extension to manage and interact with npm packages hosted on GitHub
Packages.

## Overview

`gh-npm` seamlessly integrates with GitHub CLI to provide an easy interface for
viewing and managing your npm packages hosted on GitHub Packages. Use this
extension to list your packages, view details, and publish new packages directly
from the command line.

## Features

- **List npm packages** - View all your npm packages hosted on GitHub Packages
- **Package details** - (Coming soon) See important information about each
  package including repository information, version count, and ownership
- **Package publishing** - (Coming soon) Publish npm packages directly to GitHub
  Packages

## Installation

### Prerequisites

- [GitHub CLI](https://cli.github.com/) installed and authenticated
- Go 1.16 or higher (if installing from source)

### Install as GitHub CLI extension

```bash
gh extension install ax-sh/gh-npm
```

<details>
   <summary>Installing Manually</summary>

> If you want to install this extension **manually**, follow these steps:

1. Clone the repo

   ```shell
   # git
   git clone https://github.com/ax-sh/gh-npm
   ```

   ```shell
   # GitHub CLI
   gh repo clone ax-sh/gh-npm
   ```

2. Cd into it

   ```bash
   cd gh-npm
   ```

3. Build it

   ```bash
   go build
   ```

4. Install it locally
   ```bash
   gh extension install .
   ```

</details>

## Usage

### List all npm packages

```bash
gh npm list
# or
gh npm l
```

This command will display all npm packages owned by your authenticated GitHub
user, showing:

- Update timestamp
- Owner github info
- Package type
- Package name
- Repository URL

### Configuration (Coming soon)

```bash
gh npm config
```

Configure various settings for the extension.

### Publishing packages (Coming soon)

```bash
gh npm publish
```

Publish packages directly to GitHub Packages.

## Development

### Project Structure

```
└── ./
    ├── api
    │   ├── npm-package.go    # Package handling and API interactions
    │   └── user.go           # GitHub user information
    ├── cmd
    │   ├── config.go         # Configuration command
    │   ├── list.go           # List command
    │   ├── publish.go        # Publishing command
    │   └── root.go           # Root command
    └── main.go               # Entry point
```

### Build from source

```bash
go build
```

## Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for
details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Uses the [go-gh](https://github.com/cli/go-gh) library to interact with GitHub
  API
