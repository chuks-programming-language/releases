#!/bin/bash
set -e

# Chuks Programming Language Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/chuks-programming-language/releases/main/install.sh | bash

REPO="chuks-programming-language/releases"
INSTALL_DIR="$HOME/chuks/bin"

# Detect OS and architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Darwin) PLATFORM="darwin" ;;
    Linux)  PLATFORM="linux" ;;
    MINGW*|MSYS*|CYGWIN*) PLATFORM="windows" ;;
    *) echo "Error: Unsupported operating system: $OS"; exit 1 ;;
esac

case "$ARCH" in
    x86_64|amd64)  GOARCH="amd64" ;;
    arm64|aarch64) GOARCH="arm64" ;;
    *) echo "Error: Unsupported architecture: $ARCH"; exit 1 ;;
esac

if [ "$PLATFORM" = "windows" ]; then
    ARCHIVE="chuks-windows-amd64.zip"
else
    ARCHIVE="chuks-${PLATFORM}-${GOARCH}.tar.gz"
fi

URL="https://github.com/${REPO}/releases/latest/download/${ARCHIVE}"
EXTRACT_DIR="chuks-${PLATFORM}-${GOARCH}"

echo ""
echo "  Chuks Installer"
echo "  ────────────────"
echo "  Platform:     ${PLATFORM}/${GOARCH}"
echo "  Archive:      ${ARCHIVE}"
echo "  Install to:   ${INSTALL_DIR}"
echo ""

# Create temp directory
TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

echo "→ Downloading ${ARCHIVE}..."
curl -fsSL "$URL" -o "$TMPDIR/$ARCHIVE"

echo "→ Extracting..."
cd "$TMPDIR"
if [ "$PLATFORM" = "windows" ]; then
    unzip -q "$ARCHIVE"
else
    tar xzf "$ARCHIVE"
fi

echo "→ Installing to ${INSTALL_DIR}..."
mkdir -p "$INSTALL_DIR"

if [ "$PLATFORM" = "windows" ]; then
    cp "$EXTRACT_DIR/chuks.exe" "$INSTALL_DIR/"
else
    cp "$EXTRACT_DIR/chuks" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/chuks"

    # macOS: strip quarantine/provenance xattrs and ad-hoc codesign
    if [ "$PLATFORM" = "darwin" ]; then
        echo "→ Clearing macOS quarantine flags..."
        xattr -cr "$INSTALL_DIR/chuks" 2>/dev/null || true
        xattr -dr com.apple.quarantine "$INSTALL_DIR" 2>/dev/null || true
        xattr -dr com.apple.provenance "$INSTALL_DIR" 2>/dev/null || true
        codesign --force --sign - "$INSTALL_DIR/chuks" 2>/dev/null || true
    fi

    # Verify the binary actually runs
    echo "→ Verifying installation..."
    if ! "$INSTALL_DIR/chuks" --version >/dev/null 2>&1; then
        echo ""
        echo "  ⚠  The binary was installed but macOS is blocking execution."
        echo "  Run this command to allow it:"
        echo ""
        echo "    xattr -cr $INSTALL_DIR/chuks && codesign --force --sign - $INSTALL_DIR/chuks"
        echo ""
        echo "  Then verify with: chuks --version"
        echo ""
        exit 1
    fi
fi

# Add to PATH if not already there
add_to_path() {
    local shell_rc="$1"
    local export_line='export PATH="$HOME/chuks/bin:$PATH"'
    if ! grep -q 'chuks/bin' "$shell_rc" 2>/dev/null; then
        # Create the rc file if it doesn't exist (e.g. fresh macOS has no ~/.zshrc)
        touch "$shell_rc"
        echo "" >> "$shell_rc"
        echo "# Chuks Programming Language" >> "$shell_rc"
        echo "$export_line" >> "$shell_rc"
        echo "  Added to $shell_rc"
    fi
}

if [ "$PLATFORM" != "windows" ]; then
    PATH_ADDED=false
    CURRENT_SHELL="$(basename "$SHELL" 2>/dev/null || echo "")"

    case "$CURRENT_SHELL" in
        zsh)
            add_to_path "$HOME/.zshrc"
            PATH_ADDED=true
            ;;
        bash)
            if [ -f "$HOME/.bashrc" ]; then
                add_to_path "$HOME/.bashrc"
            elif [ -f "$HOME/.bash_profile" ]; then
                add_to_path "$HOME/.bash_profile"
            fi
            PATH_ADDED=true
            ;;
    esac

    if [ "$PATH_ADDED" = false ]; then
        # Fallback: add to the default shell's rc file, creating it if needed
        for rc in "$HOME/.zshrc" "$HOME/.bashrc" "$HOME/.profile"; do
            add_to_path "$rc"
            PATH_ADDED=true
            break
        done
    fi
fi

echo ""
echo "  ✓ Chuks installed successfully!"
echo ""

# Check if already on PATH
if command -v chuks >/dev/null 2>&1; then
    VERSION="$(chuks --version 2>/dev/null || echo "unknown")"
    echo "  Version: ${VERSION}"
else
    echo "  To start using Chuks, run:"
    echo ""
    echo "    source ~/.zshrc"
    echo ""
    echo "  Or open a new terminal. Then verify with:"
    echo ""
    echo "    chuks --version"
fi

echo ""
echo "  Get started:  https://chuks.org/getting-started/hello-world/"
echo "  Docs:         https://chuks.org"
echo ""
