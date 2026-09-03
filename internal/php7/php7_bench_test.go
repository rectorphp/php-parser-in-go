package php7_test

import (
	"os"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php7"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func BenchmarkPhp7(benchmark *testing.B) {
	src, err := os.ReadFile("test.php")

	if err != nil {
		benchmark.Fatal("can not read test.php: " + err.Error())
	}

	for number := 0; number < benchmark.N; number++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 7,
				Minor: 4,
			},
		}
		lexer := scanner.NewLexer(src, config)
		php7parser := php7.NewParser(lexer, config)
		php7parser.Parse()
	}
}

func BenchmarkPhp7Heredoc(benchmark *testing.B) {
	src, err := os.ReadFile("heredoc.php")

	if err != nil {
		benchmark.Fatal("can not read heredoc.php: " + err.Error())
	}

	for number := 0; number < benchmark.N; number++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 7,
				Minor: 4,
			},
		}
		lexer := scanner.NewLexer(src, config)
		php7parser := php7.NewParser(lexer, config)
		php7parser.Parse()
	}
}
