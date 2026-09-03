package printer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php7"
	"github.com/rectorphp/php-parser-in-go/internal/php8"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

// parsePHP8Real parses with the real PHP 8 parser (the parsePHP8 helper above
// uses the PHP 7 parser), so it exercises PHP 8-only grammar.
func parsePHP8Real(src string) ast.Vertex {
	config := conf.Config{
		Version: &version.Version{
			Major: 8,
			Minor: 0,
		},
	}
	lexer := php8.NewLexer([]byte(src), config)
	php8parser := php8.NewParser(lexer, config)
	php8parser.Parse()

	return php8parser.GetRootNode()
}

func Example_php8() {
	src := `<?php

namespace Foo;

abstract class Bar extends Baz
{
    public function greet()
    {
        echo "Hello";
        // some comment
    }
}
	`

	// parse

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := scanner.NewLexer([]byte(src), config)
	php7parser := php7.NewParser(lexer, config)
	php7parser.Parse()

	rootNode := php7parser.GetRootNode()

	// change namespace

	parts := &rootNode.(*ast.Root).Stmts[0].(*ast.StmtNamespace).Name.(*ast.Name).Parts
	*parts = append(*parts, &ast.NamePart{Value: []byte("Quuz")})

	// print

	printerValue := printer.NewPrinter(os.Stdout)
	rootNode.Accept(printerValue)

	// Output:
	//<?php
	//
	// namespace Foo\Quuz;
	//
	// abstract class Bar extends Baz
	// {
	//     public function greet()
	//     {
	//         echo "Hello";
	//         // some comment
	//     }
	// }
}

func parsePHP8(src string) ast.Vertex {
	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := scanner.NewLexer([]byte(src), config)
	php7parser := php7.NewParser(lexer, config)
	php7parser.Parse()

	return php7parser.GetRootNode()
}

func printPHP8(node ast.Vertex) string {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer)
	node.Accept(printerValue)

	return buffer.String()
}

// test node

func TestParseAndPrintPHP8Root(test *testing.T) {

	src := ` <div>Hello</div> 
	<?php
	$a;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Identifier(test *testing.T) {
	src := `<? ;
	/* Foo */
	Foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ParameterTMP(test *testing.T) {

	src := `<?php
	function foo ( foo & ... $foo = null ) {}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Parameter(test *testing.T) {

	src := `<?php
	function & foo (
		? int $a , & $b = null
		, \ Foo ...$c
	) : namespace  \ Bar \  baz \ quuz{
		;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Nullable(test *testing.T) {

	src := `<?php
	function & foo ( ? int $a ) {
		/* do nothing */
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Argument(test *testing.T) {
	src := `<?php
	foo ( $a , $b
		, ... $c ,
	) ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Attributes(test *testing.T) {
	src := `<?php
	#[ FooAttribute , BarAttribute ( $arg ) , ]
	function foo ( ) { } `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test name

func TestParseAndPrintPHP8Names(test *testing.T) {
	src := `<?php
	foo ( ) ;
	\ foo ( ) ;
	namespace \ foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test scalar

func TestParseAndPrintPHP8MagicConstant(test *testing.T) {
	src := `<?php
	__CLASS__     ;
	__DIR__       ;
	__FILE__      ;
	__FUNCTION__  ;
	__LINE__      ;
	__NAMESPACE__ ;
	__METHOD__    ;
	__TRAIT__     ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Number(test *testing.T) {
	src := `<?php
	// LNumber
	1234567890123456789 ;

	// DNumber
	12345678901234567890 ;
	0. ;
	.2 ;
	0.2 ;

	// binary LNumber
	0b0111111111111111111111111111111111111111111111111111111111111111 ;

	// binary DNumber
	0b1111111111111111111111111111111111111111111111111111111111111111 ;

	// HLNumber
	0x007111111111111111 ;

	// HDNumber
	0x8111111111111111 ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8String(test *testing.T) {
	src := `<?php
	'Hello' ;
	"Hello {$world } " ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Heredoc(test *testing.T) {
	src := `<?php
	foo(<<<EAP
test
EAP
, 'test'
);

<<<EAP
test
EAP;

<<<'EAP'
test
EAP;

<<<"EAP"
test
EAP;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test assign

func TestParseAndPrintPHP8Assign(test *testing.T) {
	src := `<?php
	$a = $b ;
	$a = & $b ;
	$a &= $b ;
	$a |= $b ;
	$a ^= $b ;
	$a ??= $b ;
	$a .= $b ;
	$a /= $b ;
	$a -= $b ;
	$a %= $b ;
	$a *= $b ;
	$a += $b ;
	$a **= $b ;
	$a <<= $b ;
	$a >>= $b ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test binary

func TestParseAndPrintPHP8Binary(test *testing.T) {
	src := `<?php
	$a & $b ;
	$a | $b ;
	$a ^ $b ;
	$a && $b ;
	$a || $b ;
	$a ?? $b ;
	$a . $b ;
	$a / $b ;
	$a == $b ;
	$a >= $b ;
	$a > $b ;
	$a === $b ;
	$a and $b ;
	$a or $b ;
	$a xor $b ;
	$a - $b ;
	$a % $b ;
	$a * $b ;
	$a != $b ;
	$a <> $b ;
	$a !== $b ;
	$a + $b ;
	$a ** $b ;
	$a << $b ;
	$a >> $b ;
	$a <= $b ;
	$a < $b ;
	$a <=> $b ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test cast

func TestParseAndPrintPHP8Cast(test *testing.T) {
	src := `<?php
	(  array     ) $a ;
	(  bool      ) $a ;
	(  boolean   ) $a ;
	(  real      ) $a ;
	(  double    ) $a ;
	(  float     ) $a ;
	(  int       ) $a ;
	(  integer   ) $a ;
	(  object    ) $a ;
	(  string    ) $a ;
	(  binary    ) $a ;
	(  unset     ) $a ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test expr

func TestParseAndPrintPHP8ArrayDimFetch(test *testing.T) {
	src := `<?php
	FOO [ ] ;
	FOO [ 1 ] ;
	$a [ ] ;
	$a [ 1 ] ;
	$a { 1 } ;
	new $a [ ] ;
	new $a [ 1 ] ;
	new $a { 1 } ;
	"$a[1]test" ;
	"${ a [ 1 ] }test" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ArrayItem(test *testing.T) {
	src := `<?php
	$foo = [
		$world ,
		& $world ,
		'Hello' => $world ,
		... $unpack
	] ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Array(test *testing.T) {
	src := `<?php
	array ( /* empty array */ ) ;
	array ( 0 , 2 => 2 ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8BitwiseNot(test *testing.T) {
	src := `<?php
	~ $var ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8BooleanNot(test *testing.T) {
	src := `<?php
	! $var ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ClassConstFetch(test *testing.T) {
	src := `<?php
	$var :: CONST ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Clone(test *testing.T) {
	src := `<?php
	clone $var ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ClosureUse(test *testing.T) {
	src := `<?php
	$a = function ( ) use ( $a , & $b ) {
		// do nothing
	} ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Closure(test *testing.T) {
	src := `<?php
	$a  = static function & ( ) : void {
		// do nothing
	} ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ArrowFunction(test *testing.T) {
	src := `<?php
	$a = static fn & ( $b ) : void => $c ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ConstFetch(test *testing.T) {
	src := `<?php
	null ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Empty(test *testing.T) {
	src := `<?php
	empty ( $a ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ErrorSuppress(test *testing.T) {
	src := `<?php
	@ foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Eval(test *testing.T) {
	src := `<?php
	eval ( " " ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Exit(test *testing.T) {
	src := `<?php
	exit ;
	exit ( ) ;
	exit (1) ;
	exit ( 1 );
	die ;
	die ( ) ;
	die (1) ;
	die ( 1 );
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8FunctionCall(test *testing.T) {
	src := `<?php
	foo ( ) ;
	$var ( $a , ... $b , $c ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Include(test *testing.T) {

	src := `<?php
	include 'foo' ;
	include_once 'bar' ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8InstanceOf(test *testing.T) {
	src := `<?php
	$a instanceof Foo ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Isset(test *testing.T) {
	src := `<?php
	isset ( $a , $b [ 2 ] , ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8List(test *testing.T) {
	src := `<?php
	list( , $var , ) = $b ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8MethodCall(test *testing.T) {
	src := `<?php
	$a -> bar ( $arg , ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8New(test *testing.T) {
	src := `<?php

	new Foo ;

	new Foo ( $a, $b ) ;

	new class ( $c ) extends Foo implements Bar , Baz {

	} ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8NewReadonlyAnonymous(test *testing.T) {
	src := `<?php

	new readonly class ( $c ) extends Foo implements Bar , Baz {

	} ; `

	actual := printPHP8(parsePHP8Real(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8IncDec(test *testing.T) {
	src := `<?php
	++ $a ;
	-- $a ;
	$a ++ ;
	$a -- ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Print(test *testing.T) {
	src := `<?php
	print $a ;
	print ( $a ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8PropertyFetch(test *testing.T) {
	src := `<?php
	$a -> b ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Reference(test *testing.T) {
	src := `<?php
	$a = & $b ;
	$a = [ & $b ] ;
	$a = [ $b => & $c ] ;

	$a = function ( ) use ( & $b ) {
		// do nothing
	} ;

	foreach ( $a as & $b ) {
		// do nothing
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Require(test *testing.T) {

	src := `<?php
	require __DIR__ . '/folder' ;
	require_once $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ShellExec(test *testing.T) {
	src := "<?php ` {$v} cmd ` ; "

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ShortArray(test *testing.T) {
	src := `<?php
	$a = [ ] ;
	$a = [ 0 ] ;
	$a = [
		1 => & $b , // one
		$c       , /* two */ 
	] ;
	$a = [0, 1, 2] ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ShortList(test *testing.T) {
	src := `<?php
	[ 
		/* skip */,
		$b 
		/* skip */,
	] = $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8StaticCall(test *testing.T) {
	src := `<?php
	Foo :: bar ( $a , $b ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8StaticPropertyFetch(test *testing.T) {
	src := `<?php
	Foo :: $bar ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Ternary(test *testing.T) {
	src := `<?php
	$a ? $b : $c ;
	$a ? : $c ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Unary(test *testing.T) {
	src := `<?php
	- $a ;
	+ $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Variable(test *testing.T) {
	src := `<?php
	$ /* variable variable comment */ $var ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Yield(test *testing.T) {
	src := `<?php
	yield $a ;
	yield $k => $v ;
	yield from $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test stmt

func TestParseAndPrintPHP8AltIf(test *testing.T) {
	src := `<?php
	if ( 1 ) :
		// do nothing
	elseif ( 2 ) :
	elseif ( 3 ) :
		;
	else :
		;
	endif ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8AltFor(test *testing.T) {
	src := `<?php
	for ( $a ; $b ; $c ) :
	endfor ;
	
	for ( ; ; ) :
	endfor ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8AltForeach(test *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) :
		echo $v ;
	endforeach ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8AltSwitch(test *testing.T) {
	src := `<?php

	switch ( $a ) : 
	case 1 : 
		;
	case 2 : ;
	case 3 :
		 ;
	default :
		;
	endswitch ;
	

	switch ( $a ) : ;
	case 1 ; ;
	default ; ;
	endswitch ;

	switch ( $a ) :
	endswitch ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8AltWhile(test *testing.T) {
	src := `<?php

	while ( $a ) :
		// do nothing
	endwhile ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Break(test *testing.T) {
	src := `<?php

	break ;
	break 1 ;
	break ( 2 ) ;
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ClassMethod(test *testing.T) {
	src := `<?php
	class Foo {
		/**
		 * abstract method
		 */
		public static function & greet ( ? Foo $a ) : void ;
		
		function greet ( string $a )
		{
			return 'hello' ;
		}
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Class(test *testing.T) {
	src := `<?php
	abstract final class Foo extends Bar implements Baz , Quuz {
		
	}

	new class ( $c, $a ) extends Foo implements Bar , Baz {

	} ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ClassConstList(test *testing.T) {
	src := `<?php
	class Foo {
		public const FOO = 'f' , BAR = 'b' ;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ConstList(test *testing.T) {
	src := `<?php
	const FOO = 1 , BAR = 2 ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Continue(test *testing.T) {
	src := `<?php

	continue ;
	continue 1 ;
	continue ( 2 ) ;
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Declare(test *testing.T) {
	src := `<?php
	declare ( FOO = 'bar' , BAR = "foo" ) ;
	declare ( FOO = 'bar' ) $a ;
	declare ( FOO = 'bar' ) { }

	declare ( FOO = 'bar' ) : enddeclare ;
	declare ( FOO = 'bar' ) :
		;
	enddeclare ;
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8DoWhile(test *testing.T) {
	src := `<?php
	do {
		;
	} while ( $a ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Echo(test *testing.T) {
	src := `<?php
	echo '' ;
	echo $a , ' ' , PHP_EOL;

	?>

	<?= $a, $b ?>
	<?= $c ; 

`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8IfExpression(test *testing.T) {
	src := `<?php
	$a ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8For(test *testing.T) {
	src := `<?php
	for ( $i = 0 ; $i < 3 ; $i ++ ) 
		echo $i . PHP_EOL;
	
	for ( ; ; ) {

	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Foreach(test *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) {
		;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Function(test *testing.T) {

	src := `<?php
	function & foo ( ) : void {
		;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Global(test *testing.T) {
	src := `<?php
	global $a , $b ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Goto(test *testing.T) {
	src := `<?php
	goto Foo ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8GroupUse(test *testing.T) {
	src := `<?php
	use function Foo \ { Bar as Baz , Quuz , } ;
	use Foo \ { function Bar as Baz , const Quuz } ;
	use \ Foo \ { function Bar as Baz , const Quuz , } ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8HaltCompiler(test *testing.T) {
	src := `<?php
	__halt_compiler ( ) ;
	this text is ignored by parser
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8IfElseIfElse(test *testing.T) {
	src := `<?php
	if ( 1 ) ;
	elseif ( 2 ) {
		;
	}
	else if ( 3 ) $a;
	else { }`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8InlineHtml(test *testing.T) {
	src := `<?php
	$a;?>test<? `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Shebang(test *testing.T) {
	src := `#!/usr/bin/env php
	<?php
	$a;?>test<? `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Interface(test *testing.T) {
	src := `<?php
	interface Foo extends Bar , Baz {

	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8GotoLabel(test *testing.T) {
	src := `<?php
	Foo : $b ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Namespace(test *testing.T) {
	src := `<?php
	namespace Foo \ Bar ; 
	namespace Baz {

	}
	namespace {
		
	}
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Nop(test *testing.T) {
	src := `<?php `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8PropertyList(test *testing.T) {
	src := `<?php
	class Foo {
		var $a = '' , $b = null ;
		private $c ;
		public static Bar $d ;
		
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Return(test *testing.T) {
	src := `<?php
	class Foo {
		function bar ( )
		{
			return null ;
		}
	}

	function foo ( )
	{
		return $a ;
	}

	function bar ( )
	{
		return ;
	}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8StaticVar(test *testing.T) {
	src := `<?php
	static $a , $b = foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8StmtList(test *testing.T) {
	src := `<?php
	{
		;
	}
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Switch(test *testing.T) {
	src := `<?php

	switch ( $a ) {
		case 1 : ;
		default : ;
	}
	switch ( $a ) { ;
		case 1 ; ;
		default ; ;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Throw(test *testing.T) {
	src := `<?php
	throw new \ Exception ( "msg" ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8TraitUse(test *testing.T) {
	src := `<?php
	class foo {
		use \ foo , bar ;
		use foo , \ bar { }
		use \ foo , \ bar {
			foo :: a as b ;
			bar :: a insteadof foo ;
			foo :: c as public ;
			foo :: d as public e;
			f as g ;
		}
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Trait(test *testing.T) {
	src := `<?php
	trait foo {
		function bar ( ) { }
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8TryCatchFinally(test *testing.T) {
	src := `<?php

	try {

	} catch ( \ Exception | \ Foo \ Bar $e) {

	} finally {

	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Unset(test *testing.T) {
	src := `<?php
	unset ( $a ) ;
	unset ( $a , $b , ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8UseList(test *testing.T) {
	src := `<?php
	use Foo ;
	use \ Foo as Bar ;
	use function \ Foo as Bar ;
	use const Foo as Bar, baz ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8While(test *testing.T) {
	src := `<?php
	while ( $a ) echo '' ;
	while ( $a ) { }
	while ( $a ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// other

func TestParseAndPrintPHP8Parentheses(test *testing.T) {
	src := `<?php
	$b = (($a));
	$b = ( ($a) );
	$b = ( ( $a ) );
	$b = ( ($a ));
	$b = (( $a) );

	( $a + $b ) * 2 ;
	( $ $foo . 'foo' ) :: { $bar . 'bar' } ( ) ;
	( $ $foo ) [ 'bar' ] ;
	$ { $a . 'b' } -> call ( ) ;
	$a -> { $b . 'b' } ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ComplexString1(test *testing.T) {
	src := `<?php
	// "test $foo" ;
	"test $foo[1]" ;
	"test $foo[-1]" ;
	"test $foo[112345678901234567890] " ;
	"test $foo[-112345678901234567890] " ;
	"test $foo[a]" ;
	"test $foo[$bar]" ;
	"test $foo->bar" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ComplexString2(test *testing.T) {
	src := `<?php
	"test ${ foo }" ;
	"test ${ foo . 'bar' }" ;
	"test ${ foo [ ] }" ;
	"test ${ foo [ $b ] }" ;
	"test ${ foo [ 1 ] }" ;
	"test ${ foo [ 'expr' . $bar ] }" ;
	"test ${ $foo }" ;
	"test ${ $foo -> bar }" ;
	"test ${ $foo -> bar ( ) }" ;
	"test ${ $a . '' }" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ComplexString3(test *testing.T) {
	src := `<?php
	"test ${foo}" ;
	"test ${foo[0]}";
	"test ${foo::$bar}";
	"test ${foo }" ;
	"test ${foo . 'bar' }" ;
	"test ${foo [ ] }" ;
	"test ${foo [ $b ] }" ;
	"test ${foo [ 1 ] }" ;
	"test ${foo [ 'expr' . $bar ] }" ;
	"test ${$foo }" ;
	"test ${$foo -> bar }" ;
	"test ${$foo -> bar ( ) }" ;
	"test ${$a . '' }" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ComplexString4(test *testing.T) {
	src := `<?php
	"test {$foo }" ;
	"test {$foo [ ] }" ;
	"test {$foo [ 1 ] }" ;
	"test {$foo -> bar }" ;
	"test {$foo -> bar ( ) }" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8Match(test *testing.T) {
	src := `<?php
	$a = match ( $x ) {
		1 , 2 => 'low' ,
		3 => 'mid' ,
		default => 'high' ,
	} ;
	match ( true ) { } ;
	`

	actual := printPHP8(parsePHP8Real(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8ConstructorPromotion(test *testing.T) {
	src := `<?php
	class Point {
		public function __construct (
			public int $x = 0 ,
			protected string $name = 'p' ,
			private float ...$rest
		) { }
	}`

	actual := printPHP8(parsePHP8Real(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPHP8NamedArguments(test *testing.T) {
	src := `<?php
	foo ( name : $a , value : 1 ) ;
	new Point ( x : 1 , y : 2 ) ;
	`

	actual := printPHP8(parsePHP8Real(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}
