// Corpus benchmark: parse every .php file under the given path and report timing.
// Parsing runs across GOMAXPROCS workers with the GC disabled for the short-lived run.
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
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

	// short-lived process: parse throughput matters, not steady-state memory
	debug.SetGCPercent(-1)

	files := collectPHPFiles(root)

	start := time.Now()
	parsed := parseAll(files, config)
	elapsed := time.Since(start)

	fmt.Printf("parsed %d files in %d ms\n", parsed, elapsed.Milliseconds())
}

// parseAll parses every file across GOMAXPROCS workers and returns the count parsed.
func parseAll(files []string, config conf.Config) int64 {
	jobs := make(chan string, runtime.GOMAXPROCS(0))
	var parsed int64

	var wg sync.WaitGroup
	for range runtime.GOMAXPROCS(0) {
		wg.Go(func() {
			for path := range jobs {
				content, err := os.ReadFile(path)
				if err != nil {
					continue
				}
				// broken fixtures return an error; the parse work is what we time
				parser.Parse(content, config)
				atomic.AddInt64(&parsed, 1)
			}
		})
	}

	for _, path := range files {
		jobs <- path
	}
	close(jobs)
	wg.Wait()

	return parsed
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
