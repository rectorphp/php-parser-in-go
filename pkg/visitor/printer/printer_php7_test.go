package printer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/rectorphp/php-parser-in-go/internal/php7"
	"github.com/rectorphp/php-parser-in-go/internal/scanner"
	"github.com/rectorphp/php-parser-in-go/pkg/ast"
	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/printer"
)

func Example_php7() {
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

func parse(src string) ast.Vertex {
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

func print(node ast.Vertex) string {
	buffer := bytes.NewBufferString("")

	printerValue := printer.NewPrinter(buffer)
	node.Accept(printerValue)

	return buffer.String()
}

// test node

func TestParseAndPrintRoot(test *testing.T) {

	src := ` <div>Hello</div> 
	<?php
	$a;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIdentifier(test *testing.T) {
	src := `<? ;
	/* Foo */
	Foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintParameterTMP(test *testing.T) {

	src := `<?php
	function foo ( foo & ... $foo = null ) {}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintParameter(test *testing.T) {

	src := `<?php
	function & foo (
		? int $a , & $b = null
		, \ Foo ...$c
	) : namespace  \ Bar \  baz \ quuz{
		;
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNullable(test *testing.T) {

	src := `<?php
	function & foo ( ? int $a ) {
		/* do nothing */
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArgument(test *testing.T) {
	src := `<?php
	foo ( $a , $b
		, ... $c ,
	) ; `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test name

func TestParseAndPrintNames(test *testing.T) {
	src := `<?php
	foo ( ) ;
	\ foo ( ) ;
	namespace \ foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test scalar

func TestParseAndPrintMagicConstant(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNumber(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintString(test *testing.T) {
	src := `<?php
	'Hello' ;
	"Hello {$world } " ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHeredoc(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test assign

func TestParseAndPrintAssign(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test binary

func TestParseAndPrintBinary(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test cast

func TestParseAndPrintCast(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test expr

func TestParseAndPrintArrayDimFetch(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrayItem(test *testing.T) {
	src := `<?php
	$foo = [
		$world ,
		& $world ,
		'Hello' => $world ,
		... $unpack
	] ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArray(test *testing.T) {
	src := `<?php
	array ( /* empty array */ ) ;
	array ( 0 , 2 => 2 ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBitwiseNot(test *testing.T) {
	src := `<?php
	~ $var ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBooleanNot(test *testing.T) {
	src := `<?php
	! $var ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstFetch(test *testing.T) {
	src := `<?php
	$var :: CONST ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClone(test *testing.T) {
	src := `<?php
	clone $var ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosureUse(test *testing.T) {
	src := `<?php
	$a = function ( ) use ( $a , & $b ) {
		// do nothing
	} ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosure(test *testing.T) {
	src := `<?php
	$a  = static function & ( ) : void {
		// do nothing
	} ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrowFunction(test *testing.T) {
	src := `<?php
	$a = static fn & ( $b ) : void => $c ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintConstFetch(test *testing.T) {
	src := `<?php
	null ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEmpty(test *testing.T) {
	src := `<?php
	empty ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintErrorSuppress(test *testing.T) {
	src := `<?php
	@ foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEval(test *testing.T) {
	src := `<?php
	eval ( " " ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintExit(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionCall(test *testing.T) {
	src := `<?php
	foo ( ) ;
	$var ( $a , ... $b , $c ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInclude(test *testing.T) {

	src := `<?php
	include 'foo' ;
	include_once 'bar' ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInstanceOf(test *testing.T) {
	src := `<?php
	$a instanceof Foo ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIsset(test *testing.T) {
	src := `<?php
	isset ( $a , $b [ 2 ] , ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintList(test *testing.T) {
	src := `<?php
	list( , $var , ) = $b ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintMethodCall(test *testing.T) {
	src := `<?php
	$a -> bar ( $arg , ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNew(test *testing.T) {
	src := `<?php

	new Foo ;

	new Foo ( $a, $b ) ;

	new class ( $c ) extends Foo implements Bar , Baz {

	} ; `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIncDec(test *testing.T) {
	src := `<?php
	++ $a ;
	-- $a ;
	$a ++ ;
	$a -- ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPrint(test *testing.T) {
	src := `<?php
	print $a ;
	print ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyFetch(test *testing.T) {
	src := `<?php
	$a -> b ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintReference(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintRequire(test *testing.T) {

	src := `<?php
	require __DIR__ . '/folder' ;
	require_once $a ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShellExec(test *testing.T) {
	src := "<?php ` {$v} cmd ` ; "

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortArray(test *testing.T) {
	src := `<?php
	$a = [ ] ;
	$a = [ 0 ] ;
	$a = [
		1 => & $b , // one
		$c       , /* two */ 
	] ;
	$a = [0, 1, 2] ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortList(test *testing.T) {
	src := `<?php
	[ 
		/* skip */,
		$b 
		/* skip */,
	] = $a ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticCall(test *testing.T) {
	src := `<?php
	Foo :: bar ( $a , $b ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticPropertyFetch(test *testing.T) {
	src := `<?php
	Foo :: $bar ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTernary(test *testing.T) {
	src := `<?php
	$a ? $b : $c ;
	$a ? : $c ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnary(test *testing.T) {
	src := `<?php
	- $a ;
	+ $a ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintVariable(test *testing.T) {
	src := `<?php
	$ /* variable variable comment */ $var ; `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintYield(test *testing.T) {
	src := `<?php
	yield $a ;
	yield $k => $v ;
	yield from $a ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test stmt

func TestParseAndPrintAltIf(test *testing.T) {
	src := `<?php
	if ( 1 ) :
		// do nothing
	elseif ( 2 ) :
	elseif ( 3 ) :
		;
	else :
		;
	endif ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltFor(test *testing.T) {
	src := `<?php
	for ( $a ; $b ; $c ) :
	endfor ;
	
	for ( ; ; ) :
	endfor ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltForeach(test *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) :
		echo $v ;
	endforeach ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltSwitch(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltWhile(test *testing.T) {
	src := `<?php

	while ( $a ) :
		// do nothing
	endwhile ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBreak(test *testing.T) {
	src := `<?php

	break ;
	break 1 ;
	break ( 2 ) ;
`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassMethod(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClass(test *testing.T) {
	src := `<?php
	abstract final class Foo extends Bar implements Baz , Quuz {
		
	}

	new class ( $c, $a ) extends Foo implements Bar , Baz {

	} ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstList(test *testing.T) {
	src := `<?php
	class Foo {
		public const FOO = 'f' , BAR = 'b' ;
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintConstList(test *testing.T) {
	src := `<?php
	const FOO = 1 , BAR = 2 ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintContinue(test *testing.T) {
	src := `<?php

	continue ;
	continue 1 ;
	continue ( 2 ) ;
`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDeclare(test *testing.T) {
	src := `<?php
	declare ( FOO = 'bar' , BAR = "foo" ) ;
	declare ( FOO = 'bar' ) $a ;
	declare ( FOO = 'bar' ) { }

	declare ( FOO = 'bar' ) : enddeclare ;
	declare ( FOO = 'bar' ) :
		;
	enddeclare ;
`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDoWhile(test *testing.T) {
	src := `<?php
	do {
		;
	} while ( $a ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEcho(test *testing.T) {
	src := `<?php
	echo '' ;
	echo $a , ' ' , PHP_EOL;

	?>

	<?= $a, $b ?>
	<?= $c ; 

`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfExpression(test *testing.T) {
	src := `<?php
	$a ; `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFor(test *testing.T) {
	src := `<?php
	for ( $i = 0 ; $i < 3 ; $i ++ ) 
		echo $i . PHP_EOL;
	
	for ( ; ; ) {

	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintForeach(test *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) {
		;
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunction(test *testing.T) {

	src := `<?php
	function & foo ( ) : void {
		;
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGlobal(test *testing.T) {
	src := `<?php
	global $a , $b ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGoto(test *testing.T) {
	src := `<?php
	goto Foo ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGroupUse(test *testing.T) {
	src := `<?php
	use function Foo \ { Bar as Baz , Quuz , } ;
	use Foo \ { function Bar as Baz , const Quuz } ;
	use \ Foo \ { function Bar as Baz , const Quuz , } ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHaltCompiler(test *testing.T) {
	src := `<?php
	__halt_compiler ( ) ;
	this text is ignored by parser
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfElseIfElse(test *testing.T) {
	src := `<?php
	if ( 1 ) ;
	elseif ( 2 ) {
		;
	}
	else if ( 3 ) $a;
	else { }`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInlineHtml(test *testing.T) {
	src := `<?php
	$a;?>test<? `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShebang(test *testing.T) {
	src := `#!/usr/bin/env php
	<?php
	$a;?>test<? `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInterface(test *testing.T) {
	src := `<?php
	interface Foo extends Bar , Baz {

	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGotoLabel(test *testing.T) {
	src := `<?php
	Foo : $b ; `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNamespace(test *testing.T) {
	src := `<?php
	namespace Foo \ Bar ; 
	namespace Baz {

	}
	namespace {
		
	}
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNop(test *testing.T) {
	src := `<?php `

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyList(test *testing.T) {
	src := `<?php
	class Foo {
		var $a = '' , $b = null ;
		private $c ;
		public static Bar $d ;
		
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintReturn(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticVar(test *testing.T) {
	src := `<?php
	static $a , $b = foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStmtList(test *testing.T) {
	src := `<?php
	{
		;
	}
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintSwitch(test *testing.T) {
	src := `<?php

	switch ( $a ) {
		case 1 : ;
		default : ;
	}
	switch ( $a ) { ;
		case 1 ; ;
		default ; ;
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintThrow(test *testing.T) {
	src := `<?php
	throw new \ Exception ( "msg" ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTraitUse(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTrait(test *testing.T) {
	src := `<?php
	trait foo {
		function bar ( ) { }
	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTryCatchFinally(test *testing.T) {
	src := `<?php

	try {

	} catch ( \ Exception | \ Foo \ Bar $e) {

	} finally {

	}`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnset(test *testing.T) {
	src := `<?php
	unset ( $a ) ;
	unset ( $a , $b , ) ;`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUseList(test *testing.T) {
	src := `<?php
	use Foo ;
	use \ Foo as Bar ;
	use function \ Foo as Bar ;
	use const Foo as Bar, baz ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintWhile(test *testing.T) {
	src := `<?php
	while ( $a ) echo '' ;
	while ( $a ) { }
	while ( $a ) ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// other

func TestParseAndPrintParentheses(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString1(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString2(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString3(test *testing.T) {
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

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString4(test *testing.T) {
	src := `<?php
	"test {$foo }" ;
	"test {$foo [ ] }" ;
	"test {$foo [ 1 ] }" ;
	"test {$foo -> bar }" ;
	"test {$foo -> bar ( ) }" ;
	`

	actual := print(parse(src))

	if src != actual {
		test.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}
