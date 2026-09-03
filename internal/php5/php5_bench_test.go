package php5_test

import (
	"os"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php5"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func BenchmarkPhp5(benchmark *testing.B) {
	src, err := os.ReadFile("test.php")
	if err != nil {
		benchmark.Fatal("can not read test.php: " + err.Error())
	}

	for number := 0; number < benchmark.N; number++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 5,
				Minor: 6,
			},
		}
		lexer := scanner.NewLexer(src, config)
		php5parser := php5.NewParser(lexer, config)
		php5parser.Parse()
	}
}
