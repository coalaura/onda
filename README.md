# onda

onda is a tiny file type detector for Go.

## Features

- Fast magic-byte/signature based detection
- Supports many common formats (images, archives, audio, video, executables, documents, disk images)
- Includes custom parsers for formats that need structural checks (not just fixed headers)
- Works as both a CLI and a Go package

## Installation

Prebuilt releases are available [here](https://github.com/coalaura/onda/releases) or build it yourself.

## Usage

```bash
onda <file>
```

Example:

```bash
onda sample.png
```

## Go package

```go
package main

import (
	"fmt"

	"github.com/coalaura/onda/types"
	_ "github.com/coalaura/onda/types/all"
)

func main() {
	meta, err := types.Detect("sample.png", []byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a})
	if err != nil {
		fmt.Println("unknown")

		return
	}

	fmt.Println(meta.Name, meta.Type)
}
```
