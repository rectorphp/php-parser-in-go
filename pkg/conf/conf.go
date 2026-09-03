package conf

import (
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(parserError *errors.Error)
}
