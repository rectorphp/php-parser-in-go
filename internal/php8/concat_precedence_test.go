package php8_test

import (
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/ast"
)

func concatAssignedValue(test *testing.T, source string) ast.Vertex {
	test.Helper()
	root, parserErrors := parsePHP85(source)
	if len(parserErrors) != 0 {
		test.Fatalf("unexpected parser errors: %v", parserErrors)
	}
	return root.(*ast.Root).Stmts[0].(*ast.StmtExpression).Expr.(*ast.ExprAssign).Expr
}

// Since PHP 8.0 `+` and `-` bind tighter than `.`: "Sum: " . $a + $b is "Sum: " . ($a + $b).
func TestConcatBindsLooserThanPlus(test *testing.T) {
	value := concatAssignedValue(test, `<?php $label = "Sum: " . $a + $b;`)

	concat, ok := value.(*ast.ExprBinaryConcat)
	if !ok {
		test.Fatalf("expected the concatenation at the top, got %T", value)
	}
	if _, ok := concat.Right.(*ast.ExprBinaryPlus); !ok {
		test.Errorf("expected the addition under the concatenation, got %T", concat.Right)
	}
}

// Shifts bind tighter than `.` too: "x" . 1 << 2 is "x" . (1 << 2).
func TestConcatBindsLooserThanShift(test *testing.T) {
	value := concatAssignedValue(test, `<?php $label = "x" . 1 << 2;`)

	concat, ok := value.(*ast.ExprBinaryConcat)
	if !ok {
		test.Fatalf("expected the concatenation at the top, got %T", value)
	}
	if _, ok := concat.Right.(*ast.ExprBinaryShiftLeft); !ok {
		test.Errorf("expected the shift under the concatenation, got %T", concat.Right)
	}
}

// `.` stays left-associative and above comparison.
func TestConcatChainAndComparison(test *testing.T) {
	value := concatAssignedValue(test, `<?php $same = $a . $b . $c === $d;`)

	identical, ok := value.(*ast.ExprBinaryIdentical)
	if !ok {
		test.Fatalf("expected the comparison at the top, got %T", value)
	}
	outer, ok := identical.Left.(*ast.ExprBinaryConcat)
	if !ok {
		test.Fatalf("expected a concatenation under the comparison, got %T", identical.Left)
	}
	if _, ok := outer.Left.(*ast.ExprBinaryConcat); !ok {
		test.Errorf("expected ($a . $b) . $c, got %T on the left", outer.Left)
	}
}
