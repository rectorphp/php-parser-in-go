package printer_test

import (
	"bytes"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php5"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func parsePhp5(src string) ast.Vertex {
	config := conf.Config{
		Version: &version.Version{
			Major: 5,
			Minor: 6,
		},
	}
	lexer := scanner.NewLexer([]byte(src), config)
	php5parser := php5.NewParser(lexer, config)
	php5parser.Parse()

	return php5parser.GetRootNode()
}

func printPhp5(node ast.Vertex) string {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer)
	node.Accept(printerValue)

	return buffer.String()
}

// test node

func TestParseAndPrintPhp5Root(test *testing.T) {

	src := ` <div>Hello</div> 
	<?php
	$a;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Identifier(test *testing.T) {
	src := `<? ;
	/* Foo */
	Foo ( ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Parameter(test *testing.T) {

	src := `<?php
	function & foo ( $a , & $b = null , ... $c ) {
		;
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Argument(test *testing.T) {
	src := `<?php
	foo ( $a , $b
		, $c
	) ; `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test name

func TestParseAndPrintPhp5Names(test *testing.T) {
	src := `<?php
	foo ( ) ;
	\ foo ( ) ;
	namespace \ foo ( ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test scalar

func TestParseAndPrintPhp5MagicConstant(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Number(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5String(test *testing.T) {
	src := `<?php
	'Hello' ;
	"Hello {$world } " ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Heredoc(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test assign

func TestParseAndPrintPhp5Assign(test *testing.T) {
	src := `<?php
	$a = $b ;
	$a = & $b ;
	$a &= $b ;
	$a |= $b ;
	$a ^= $b ;
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test binary

func TestParseAndPrintPhp5Binary(test *testing.T) {
	src := `<?php
	$a & $b ;
	$a | $b ;
	$a ^ $b ;
	$a && $b ;
	$a || $b ;
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
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test cast

func TestParseAndPrintPhp5Cast(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test expr

func TestParseAndPrintPhp5ArrayDimFetch(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ArrayItem(test *testing.T) {
	src := `<?php
	$foo = [
		$world ,
		& $world ,
		'Hello' => $world ,
	] ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Array(test *testing.T) {
	src := `<?php
	array ( /* empty array */ ) ;
	array ( 0 , 2 => 2 ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5BitwiseNot(test *testing.T) {
	src := `<?php
	~ $var ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5BooleanNot(test *testing.T) {
	src := `<?php
	! $var ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ClassConstFetch(test *testing.T) {
	src := `<?php
	$var :: CONSTANT ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Clone(test *testing.T) {
	src := `<?php
	clone $var ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ClosureUse(test *testing.T) {
	src := `<?php
	$a = function ( ) use ( $a , & $b ) {
		// do nothing
	} ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Closure(test *testing.T) {
	src := `<?php
	$a  = static function & ( ) {
		// do nothing
	} ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ConstFetch(test *testing.T) {
	src := `<?php
	null ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Empty(test *testing.T) {
	src := `<?php
	empty ( $a ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ErrorSuppress(test *testing.T) {
	src := `<?php
	@ foo ( ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Eval(test *testing.T) {
	src := `<?php
	eval ( " " ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Exit(test *testing.T) {
	src := `<?php
	exit ;
	exit ( ) ;
	exit (1) ;
	exit ( 1 ) ;
	die ;
	die ( ) ;
	die (1) ;
	die ( 1 ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5FunctionCall(test *testing.T) {
	src := `<?php
	foo ( ) ;
	$var ( $a , ... $b , $c ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Include(test *testing.T) {

	src := `<?php
	include 'foo' ;
	include_once 'bar' ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5InstanceOf(test *testing.T) {
	src := `<?php
	$a instanceof Foo ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Isset(test *testing.T) {
	src := `<?php
	isset ( $a , $b [ 2 ] ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5List(test *testing.T) {
	src := `<?php
	list( , $var , ) = $b ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5MethodCall(test *testing.T) {
	src := `<?php
	$a -> bar ( $arg ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5New(test *testing.T) {
	src := `<?php

	new Foo ;

	new Foo ( $a, $b ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5IncDec(test *testing.T) {
	src := `<?php
	++ $a ;
	-- $a ;
	$a ++ ;
	$a -- ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Print(test *testing.T) {
	src := `<?php
	print $a ;
	print ( $a ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5PropertyFetch(test *testing.T) {
	src := `<?php
	$a -> b ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Reference(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Require(test *testing.T) {

	src := `<?php
	require __DIR__ . '/folder' ;
	require_once $a ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ShellExec(test *testing.T) {
	src := "<?php ` {$v} cmd ` ; "

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ShortArray(test *testing.T) {
	src := `<?php
	$a = [ ] ;
	$a = [ 0 ] ;
	$a = [
		1 => & $b , // one
		$c       , /* two */ 
	] ;
	$a = [0, 1, 2] ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5StaticCall(test *testing.T) {
	src := `<?php
	Foo :: bar ( $a , $b ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5StaticPropertyFetch(test *testing.T) {
	src := `<?php
	Foo :: $bar ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Ternary(test *testing.T) {
	src := `<?php
	$a ? $b : $c ;
	$a ? : $c ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Unary(test *testing.T) {
	src := `<?php
	- $a ;
	+ $a ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Variable(test *testing.T) {
	src := `<?php
	$ /* variable variable comment */ $var ; `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Yield(test *testing.T) {
	src := `<?php
	yield $a ;
	yield $k => $v ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test stmt

func TestParseAndPrintPhp5AltIf(test *testing.T) {
	src := `<?php
	if ( 1 ) :
		// do nothing
	elseif ( 2 ) :
	elseif ( 3 ) :
		;
	else :
		;
	endif ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5AltFor(test *testing.T) {
	src := `<?php
	for ( $a ; $b ; $c ) :
	endfor ;
	
	for ( ; ; ) :
	endfor ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5AltForeach(test *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) :
		echo $v ;
	endforeach ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5AltSwitch(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5AltWhile(test *testing.T) {
	src := `<?php

	while ( $a ) :
		// do nothing
	endwhile ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Break(test *testing.T) {
	src := `<?php

	break ;
	break 1 ;
	break ( 2 ) ;
`

	root := parsePhp5(src)
	actual := printPhp5(root)

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ClassMethod(test *testing.T) {
	src := `<?php
	class Foo {
		/**
		 * abstract method
		 */
		public static function & greet ( $a ) ;
		
		function greet ( $a )
		{
			return 'hello' ;
		}
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Class(test *testing.T) {
	src := `<?php
	final class Foo extends Bar implements Baz , Quuz {
		
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ClassConstList(test *testing.T) {
	src := `<?php
	class Foo {
		const FOO = 'f' , BAR = 'b' ;
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ConstList(test *testing.T) {
	src := `<?php
	const FOO = 1 , BAR = 2 ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Continue(test *testing.T) {
	src := `<?php

	continue ;
	continue 1 ;
	continue ( 2 ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Declare(test *testing.T) {
	src := `<?php
	declare ( FOO = 'bar' , BAR = 'foo' ) ;
	declare ( FOO = 'bar' ) $a ;
	declare ( FOO = 'bar' ) { }

	declare ( FOO = 'bar' ) : enddeclare ;
	declare ( FOO = 'bar' ) :
		;
	enddeclare ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5DoWhile(test *testing.T) {
	src := `<?php
	do {
		;
	} while ( $a ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Echo(test *testing.T) {
	src := `<?php
	echo '' ;
	echo $a , ' ' , PHP_EOL;

	?>

	<?= $a, $b ?>
	<?= $c ; `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5IfExpression(test *testing.T) {
	src := `<?php
	$a ; `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5For(test *testing.T) {
	src := `<?php
	for ( $i = 0 ; $i < 3 ; $i ++ ) 
		echo $i . PHP_EOL;
	
	for ( ; ; ) {

	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Foreach(test *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) {
		;
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Function(test *testing.T) {

	src := `<?php
	function & foo ( ) {
		;
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Global(test *testing.T) {
	src := `<?php
	global $a , $b ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Goto(test *testing.T) {
	src := `<?php
	goto Foo ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5HaltCompiler(test *testing.T) {
	src := `<?php
	__halt_compiler ( ) ;
	this text is ignored by parser
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5IfElseIfElse(test *testing.T) {
	src := `<?php
	if ( 1 ) ;
	elseif ( 2 ) {
		;
	}
	else if ( 3 ) $a;
	else { }`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5InlineHtml(test *testing.T) {
	src := `<?php
	$a;?>test<? `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Shebang(test *testing.T) {
	src := `#!/usr/bin/env php
	<?php
	$a;?>test<? `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Interface(test *testing.T) {
	src := `<?php
	interface Foo extends Bar , Baz {

	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5GotoLabel(test *testing.T) {
	src := `<?php
	Foo : $b ; `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Namespace(test *testing.T) {
	src := `<?php
	namespace Foo \ Bar ; 
	namespace Baz {

	}
	namespace {
		
	}
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Nop(test *testing.T) {
	src := `<?php `

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5PropertyList(test *testing.T) {
	src := `<?php
	class Foo {
		var $a = '' , $b = null ;
		private $c ;
		public static $d ;
		
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Return(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5StaticVar(test *testing.T) {
	src := `<?php
	static $a , $b = ' ' ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5StmtList(test *testing.T) {
	src := `<?php
	{
		;
	}
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Switch(test *testing.T) {
	src := `<?php

	switch ( $a ) {
		case 1 : ;
		default : ;
	}
	switch ( $a ) { ;
		case 1 ; ;
		default ; ;
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Throw(test *testing.T) {
	src := `<?php
	throw new \ Exception ( "msg" ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5TraitUse(test *testing.T) {
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

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Trait(test *testing.T) {
	src := `<?php
	trait foo {
		function bar ( ) { }
	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5TryCatchFinally(test *testing.T) {
	src := `<?php

	try {

	} catch ( \ Exception $e) {

	} finally {

	}`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5Unset(test *testing.T) {
	src := `<?php
	unset ( $a ) ;
	unset ( $a , $b ) ;`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5UseList(test *testing.T) {
	src := `<?php
	use Foo ;
	use \ Foo as Bar ;
	use function \ Foo as Bar ;
	use const Foo as Bar, baz ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5While(test *testing.T) {
	src := `<?php
	while ( $a ) echo '' ;
	while ( $a ) { }
	while ( $a ) ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// other

func TestParseAndPrintPhp5Parentheses(test *testing.T) {
	src := `<?php
	global $ { $b } ;
	$b = (($a));
	$b = ( ($a) );
	$b = ( ( $a ) );
	$b = ( ($a ));
	$b = (( $a) );

	( $a + $b ) * 2 ;
	$ { $a . 'b' } -> call ( ) ;
	$a -> { $b . 'b' } ;
	$a -> $b ( ) -> { $c . 'c' } ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ComplexString1(test *testing.T) {
	src := `<?php
	// "test $foo" ;
	"test $foo[1]" ;
	"test $foo[112345678901234567890] " ;
	"test $foo[a]" ;
	"test $foo[$bar]" ;
	"test $foo->bar" ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ComplexString2(test *testing.T) {
	src := `<?php
	"test ${ foo }" ;
	"test ${ foo . 'bar' }" ;
	"test ${ foo [ ] }" ;
	"test ${ foo [ 1 ] }" ;
	"test ${ foo [ 'expr' . $bar ] }" ;
	"test ${ $foo }" ;
	"test ${ $foo -> bar }" ;
	"test ${ $foo -> bar ( ) }" ;
	"test ${ $a . '' }" ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ComplexString3(test *testing.T) {
	src := `<?php
	"test ${foo }" ;
	"test ${foo . 'bar' }" ;
	"test ${foo [ ] }" ;
	"test ${foo [ 1 ] }" ;
	"test ${foo [ 'expr' . $bar ] }" ;
	"test ${$foo }" ;
	"test ${$foo -> bar }" ;
	"test ${$foo -> bar ( ) }" ;
	"test ${$a . '' }" ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPhp5ComplexString4(test *testing.T) {
	src := `<?php
	"test {$foo }" ;
	"test {$foo [ ] }" ;
	"test {$foo [ 1 ] }" ;
	"test {$foo -> bar }" ;
	"test {$foo -> bar ( ) }" ;
	`

	actual := printPhp5(parsePhp5(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}
