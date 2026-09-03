// Corpus benchmark: parse every .php file under the given path and report timing.
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	phpVersion, err := version.New("8.3")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	config := conf.Config{Version: phpVersion}

	files := collectPHPFiles(root)

	start := time.Now()
	for _, path := range files {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		// broken fixtures return an error; the parse work is what we time
		parser.Parse(content, config)
	}
	elapsed := time.Since(start)

	fmt.Printf("parsed %d files in %d ms\n", len(files), elapsed.Milliseconds())
}

func collectPHPFiles(root string) []string {
	var files []string
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && filepath.Ext(path) == ".php" {
			files = append(files, path)
		}
		return nil
	})
	return files
}
