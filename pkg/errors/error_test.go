package errors_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
)

func TestConstructor(test *testing.T) {
	pos := position.NewPosition(1, 2, 3, 4)

	actual := errors.NewError("message", pos)

	expected := &errors.Error{
		Msg: "message",
		Pos: pos,
	}

	assert.DeepEqual(test, expected, actual)
}

func TestPrint(test *testing.T) {
	pos := position.NewPosition(1, 2, 3, 4)

	Error := errors.NewError("message", pos)

	actual := Error.String()

	expected := "message at line 1"

	assert.DeepEqual(test, expected, actual)
}

func TestPrintWithotPos(test *testing.T) {
	Error := errors.NewError("message", nil)

	actual := Error.String()

	expected := "message"

	assert.DeepEqual(test, expected, actual)
}
