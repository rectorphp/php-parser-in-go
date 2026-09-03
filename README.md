# php-parser-in-go

A PHP parser written in Go. It lexes and parses PHP 5, 7 and 8 source into an AST that downstream tools (linters, refactoring, metrics, formatters) can traverse.

## Install

```bash
go get github.com/rectorphp/php-parser-in-go
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/dumper"
)

func main() {
	phpVersion, _ := version.New("8.3")

	root, err := parser.Parse([]byte("<?php echo 'hi';"), conf.Config{
		Version: phpVersion,
	})
	if err != nil {
		panic(err)
	}

	// walk / print / dump the AST
	fmt.Print(dumper.Dump(root))
}
```

## Packages

- `pkg/parser` — `parser.Parse(source, conf.Config{...})` entry point.
- `pkg/ast` — AST node definitions.
- `pkg/visitor` — traverser, printer, dumper, namespace and class resolvers, formatter.
- `pkg/token`, `pkg/position`, `pkg/version`, `pkg/errors`, `pkg/conf`.

## Generated code

`internal/*/php*.go` and `internal/*/scanner.go` are generated. Edit the grammar source instead:

- `*.y` files are regenerated with [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc).
- `*.rl` files are regenerated with [ragel](https://www.colm.net/open-source/ragel/).

Run `make build` to regenerate and build (both `goyacc` and `ragel` must be installed). `make test` runs the tests.

## Credits

Based on [z7zmey/php-parser](https://github.com/z7zmey/php-parser), the original PHP parser written in Go.
