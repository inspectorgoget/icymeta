# icymeta
[![Release](https://img.shields.io/github/v/release/inspectorgoget/icymeta?sort=semver)](https://github.com/inspectorgoget/icymeta/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/inspectorgoget/icymeta)
[![Build](https://github.com/inspectorgoget/icymeta/actions/workflows/build.yml/badge.svg)](https://github.com/inspectorgoget/icymeta/actions/workflows/build.yml)

Parses metadata from SHOUTcast audio streams.

```go
package main

import (
	"context"
	"fmt"
	"github.com/inspectorgoget/icymeta"
)

func main() {
	streamUrl := "..."
	title, err := icymeta.GetCurrentStreamTitle(context.Background(), streamUrl)

	if err != nil {
		panic(fmt.Sprintf("failed to get current stream title: %v", err))
	}

	fmt.Printf("Current stream title: %s\n", title)
}
```
