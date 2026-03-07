# Chuks Releases

Official release binaries for the [Chuks Programming Language](https://chuks.org).

## Downloads

| Platform | Architecture | File |
|---|---|---|
| macOS | Apple Silicon (ARM64) | [`chuks-darwin-arm64.tar.gz`](https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-darwin-arm64.tar.gz) |
| macOS | Intel (x86_64) | [`chuks-darwin-amd64.tar.gz`](https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-darwin-amd64.tar.gz) |
| Linux | x86_64 | [`chuks-linux-amd64.tar.gz`](https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-linux-amd64.tar.gz) |
| Linux | ARM64 | [`chuks-linux-arm64.tar.gz`](https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-linux-arm64.tar.gz) |
| Windows | x86_64 | [`chuks-windows-amd64.zip`](https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-windows-amd64.zip) |

## Quick Install

```bash
# macOS Apple Silicon
curl -L https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-darwin-arm64.tar.gz | tar xz
cd chuks-darwin-arm64 && ./scripts/install.sh

# macOS Intel
curl -L https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-darwin-amd64.tar.gz | tar xz
cd chuks-darwin-amd64 && ./scripts/install.sh

# Linux x86_64
curl -L https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-linux-amd64.tar.gz | tar xz
cd chuks-linux-amd64 && ./scripts/install.sh

# Linux ARM64
curl -L https://github.com/chuks-programming-language/releases/releases/latest/download/chuks-linux-arm64.tar.gz | tar xz
cd chuks-linux-arm64 && ./scripts/install.sh
```

Then open a new terminal and verify:

```bash
chuks --version
```

## All Releases

See the [Releases page](https://github.com/chuks-programming-language/releases/releases) for all versions.

## Installation Guide

For detailed instructions, visit the [installation page](https://chuks.org/getting-started/installation/) on the official website.

## License

See the [Chuks website](https://chuks.org) for license details.
