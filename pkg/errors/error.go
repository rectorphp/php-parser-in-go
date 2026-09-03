package errors

import (
	"fmt"

	"github.com/rectorphp/php-parser-in-go/pkg/position"
)

// Error parsing error
type Error struct {
	Msg string
	Pos *position.Position
}

// NewError creates and returns new Error
func NewError(message string, errorPosition *position.Position) *Error {
	return &Error{
		Msg: message,
		Pos: errorPosition,
	}
}

func (parseError *Error) String() string {
	atLine := ""
	if parseError.Pos != nil {
		atLine = fmt.Sprintf(" at line %d", parseError.Pos.StartLine)
	}

	return fmt.Sprintf("%s%s", parseError.Msg, atLine)
}
