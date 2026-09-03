/*
A Parser for PHP written in Go

Package usage example:

	package main

	import (
		"log"
		"os"

		"github.com/rectorphp/php-parser-in-go/pkg/conf"
		"github.com/rectorphp/php-parser-in-go/pkg/errors"
		"github.com/rectorphp/php-parser-in-go/pkg/parser"
		"github.com/rectorphp/php-parser-in-go/pkg/version"
		"github.com/rectorphp/php-parser-in-go/pkg/visitor/dumper"
	)

	func main() {
		source := []byte(`<? echo "Hello world";`)

		// Error handler

		var parserErrors []*errors.Error
		errorHandler := func(parseError *errors.Error) {
			parsmakeerErrors = append(parserErrors, parseError)
		}

		// Parse

		rootNode, err := parser.Parse(source, conf.Config{
			Version:          &version.Version{Major: 5, Minor: 6},
			ErrorHandlerFunc: errorHandler,
		})

		if err != nil {
			log.Fatal("Error:" + err.Error())
		}

		// Dump

		goDumper := dumper.NewDumper(os.Stdout).
			WithTokens().
			WithPositions()

		rootNode.Accept(goDumper)
	}
*/
package parser
