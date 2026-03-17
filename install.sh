#!/bin/bash
set -e

OS=$(uname -s | tr 'A-Z' 'a-z')

ARCH=$(uname -m)
case "$ARCH" in
	x86_64)
		ARCH=amd64
		;;
	aarch64|arm64)
		ARCH=arm64
		;;
	*)
		echo "Unsupported architecture: $ARCH" >&2
		exit 1
		;;
esac

echo "Resolving latest version of onda..."

VERSION=$(curl -sL https://api.github.com/repos/coalaura/onda/releases/latest | grep -Po '"tag_name": *"\K.*?(?=")')

if ! printf '%s\n' "$VERSION" | grep -Eq '^v[0-9]+\.[0-9]+\.[0-9]+$'; then
	echo "Error: '$VERSION' is not in vMAJOR.MINOR.PATCH format" >&2
	exit 1
fi

rm -f /tmp/onda

BIN="onda_${VERSION}_${OS}_${ARCH}"
URL="https://github.com/coalaura/onda/releases/download/${VERSION}/${BIN}"

echo "Downloading ${BIN}..."

if ! curl -sL "$URL" -o /tmp/onda; then
	echo "Error: failed to download $URL" >&2
	exit 1
fi

trap 'rm -f /tmp/onda' EXIT

chmod +x /tmp/onda

echo "Installing to /usr/local/bin/onda requires sudo"

if ! sudo install -m755 /tmp/onda /usr/local/bin/onda; then
	echo "Error: install failed" >&2
	exit 1
fi

echo "onda $VERSION installed to /usr/local/bin/onda"