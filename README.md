# mcleaner

[![CI](https://github.com/cdigiuseppe/mcleaner/workflows/CI/badge.svg)](https://github.com/cdigiuseppe/mcleaner/actions)
[![Build](https://github.com/cdigiuseppe/mcleaner/workflows/Build/badge.svg)](https://github.com/cdigiuseppe/mcleaner/actions)
[![Release](https://github.com/cdigiuseppe/mcleaner/workflows/Release/badge.svg)](https://github.com/cdigiuseppe/mcleaner/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A powerful command-line utility to clean junk files and cache on macOS.

## Features

- **Scan Mode**: Identify junk files without removing them
- **Cache Cleaning**: Remove system and user cache files
- **DS_Store Cleanup**: Find and remove all .DS_Store files
- **Temporary Files**: Clean temporary and old downloaded files
- **System Maintenance**: Run macOS daily, weekly, and monthly maintenance scripts

## Installation

### Download Binary

Download the latest release for your platform from the [releases page](https://github.com/cdigiuseppe/mcleaner/releases).

### Build from Source

```bash
git clone https://github.com/cdigiuseppe/mcleaner.git
cd mcleaner
make build
```

### Install via Go

```bash
go install github.com/cdigiuseppe/mcleaner@latest
```

## Usage

### Basic Commands

```bash
# Scan for junk files (no removal)
mcleaner scan

# Clean system cache
mcleaner clean-cache

# Remove all .DS_Store files
mcleaner clean-dsstore

# Clean temporary files
mcleaner clean-temp
```

### Maintenance Commands

```bash
# Run daily maintenance
mcleaner maintenance daily

# Run weekly maintenance
mcleaner maintenance weekly

# Run monthly maintenance
mcleaner maintenance monthly
```

### Help

```bash
# Show general help
mcleaner --help

# Show command-specific help
mcleaner scan --help
```

## Command Details

### `scan`
Performs a comprehensive scan of your system to identify:
- Cache files in system and user directories
- .DS_Store files throughout the filesystem
- Temporary files and old downloads
- Reports total recoverable space

### `clean-cache`
Removes cache files from:
- `~/Library/Caches`
- `/Library/Caches`
- `/tmp`

### `clean-dsstore`
Recursively finds and removes all `.DS_Store` files starting from your home directory.

### `clean-temp`
Removes temporary files including:
- Files with `.tmp` and `.temp` extensions
- Files containing "temp" in the name
- Old installer files in Downloads (older than 30 days)

### `maintenance`
Runs macOS system maintenance scripts:
- **daily**: Runs daily periodic scripts
- **weekly**: Runs weekly periodic scripts  
- **monthly**: Runs monthly periodic scripts

*Note: Maintenance commands require sudo privileges.*

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile)

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Run all checks (format, vet, lint, test)
make check
```

### Available Make Targets

```bash
make help  # Show all available targets
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Releasing

To create a new release:

1. Ensure all changes are committed
2. Run `make release` and enter the version tag (e.g., `v1.0.0`)
3. The GitHub Actions workflow will automatically build and publish the release

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Safety

This tool modifies and deletes files on your system. While it targets known safe-to-remove locations, always:

- Run `scan` first to see what would be removed
- Backup important data before running cleanup commands
- Test on a non-production system first

## Platform Support

- ✅ macOS (primary target)
- ✅ Linux (basic support)
- ⚠️ Windows (limited functionality)

*Note: This tool is specifically designed for macOS and some features may not work on other platforms.*