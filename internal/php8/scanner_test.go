package php8

import (
	"gotest.tools/assert"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/errors"
	"github.com/rectorphp/php-parser-in-go/pkg/position"
	"github.com/rectorphp/php-parser-in-go/pkg/token"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func TestTokens(test *testing.T) {
	src := `inline html - 
		<? ?>
		<?= ?>
		<?php

		abstract
		array
		as
		break
		callable
		case
		catch
		class
		clone
		const
		continue
		declare
		default
		do
		echo
		else
		elseif
		empty
		enddeclare
		endfor
		endforeach
		endif
		endswitch
		endwhile
		eval
		exit
		extends
		final
		finally
		for
		foreach
		function
		cfunction
		global
		goto
		if
		isset
		implements
		instanceof
		insteadof
		interface
		list
		namespace
		private
		public
		print
		protected
		return
		static
		switch
		throw
		trait
		try
		unset
		use
		var
		while
		yield ` + "\t\r\n" + ` from
		yield
		include
		include_once
		require
		require_once

		__CLASS__
		__DIR__
		__FILE__
		__FUNCTION__
		__LINE__
		__NAMESPACE__
		__METHOD__
		__TRAIT__
		__halt_compiler

		new
		and
		or
		xor

		\
		...
		::
		&&
		||
		&=
		|=
		.=
		*=
		**=
		/=
		+=
		-=
		^=
		%=
		--
		++
		=>
		<=>
		!=
		<>
		!==
		==
		===
		<<=
		>>=
		>=
		<=
		**
		<<
		>>
		??

		#  inline comment
		// inline comment

		/*
			multiline comment
		*/

		/**
		 * PHP Doc comment
		 */

		;
		:
		,
		.
		[
		]
		(
		)
		|
		/
		^
		&
		+
		-
		*
		=
		%
		!
		~
		$
		<
		>
		?
		@
		{
		}

		$var
		str

		-> ` + "\t\r\n" + ` ->prop

		( array )
		( bool )
		( boolean )
		( double )
		( float )
		( int )
		( integer )
		( object )
		( string )
		( binary )
		( unset )

	`

	expected := []string{
		token.T_INLINE_HTML.String(),
		token.ID(int(';')).String(),
		token.T_INLINE_HTML.String(),
		token.T_ECHO.String(),
		token.ID(int(';')).String(),
		token.T_INLINE_HTML.String(),

		token.T_ABSTRACT.String(),
		token.T_ARRAY.String(),
		token.T_AS.String(),
		token.T_BREAK.String(),
		token.T_CALLABLE.String(),
		token.T_CASE.String(),
		token.T_CATCH.String(),
		token.T_CLASS.String(),
		token.T_CLONE.String(),
		token.T_CONST.String(),
		token.T_CONTINUE.String(),
		token.T_DECLARE.String(),
		token.T_DEFAULT.String(),
		token.T_DO.String(),
		token.T_ECHO.String(),
		token.T_ELSE.String(),
		token.T_ELSEIF.String(),
		token.T_EMPTY.String(),
		token.T_ENDDECLARE.String(),
		token.T_ENDFOR.String(),
		token.T_ENDFOREACH.String(),
		token.T_ENDIF.String(),
		token.T_ENDSWITCH.String(),
		token.T_ENDWHILE.String(),
		token.T_EVAL.String(),
		token.T_EXIT.String(),
		token.T_EXTENDS.String(),
		token.T_FINAL.String(),
		token.T_FINALLY.String(),
		token.T_FOR.String(),
		token.T_FOREACH.String(),
		token.T_FUNCTION.String(),
		token.T_FUNCTION.String(),
		token.T_GLOBAL.String(),
		token.T_GOTO.String(),
		token.T_IF.String(),
		token.T_ISSET.String(),
		token.T_IMPLEMENTS.String(),
		token.T_INSTANCEOF.String(),
		token.T_INSTEADOF.String(),
		token.T_INTERFACE.String(),
		token.T_LIST.String(),
		token.T_NAMESPACE.String(),
		token.T_PRIVATE.String(),
		token.T_PUBLIC.String(),
		token.T_PRINT.String(),
		token.T_PROTECTED.String(),
		token.T_RETURN.String(),
		token.T_STATIC.String(),
		token.T_SWITCH.String(),
		token.T_THROW.String(),
		token.T_TRAIT.String(),
		token.T_TRY.String(),
		token.T_UNSET.String(),
		token.T_USE.String(),
		token.T_VAR.String(),
		token.T_WHILE.String(),
		token.T_YIELD_FROM.String(),
		token.T_YIELD.String(),
		token.T_INCLUDE.String(),
		token.T_INCLUDE_ONCE.String(),
		token.T_REQUIRE.String(),
		token.T_REQUIRE_ONCE.String(),

		token.T_CLASS_C.String(),
		token.T_DIR.String(),
		token.T_FILE.String(),
		token.T_FUNC_C.String(),
		token.T_LINE.String(),
		token.T_NS_C.String(),
		token.T_METHOD_C.String(),
		token.T_TRAIT_C.String(),
		token.T_HALT_COMPILER.String(),

		token.T_NEW.String(),
		token.T_LOGICAL_AND.String(),
		token.T_LOGICAL_OR.String(),
		token.T_LOGICAL_XOR.String(),

		token.T_NS_SEPARATOR.String(),
		token.T_ELLIPSIS.String(),
		token.T_PAAMAYIM_NEKUDOTAYIM.String(),
		token.T_BOOLEAN_AND.String(),
		token.T_BOOLEAN_OR.String(),
		token.T_AND_EQUAL.String(),
		token.T_OR_EQUAL.String(),
		token.T_CONCAT_EQUAL.String(),
		token.T_MUL_EQUAL.String(),
		token.T_POW_EQUAL.String(),
		token.T_DIV_EQUAL.String(),
		token.T_PLUS_EQUAL.String(),
		token.T_MINUS_EQUAL.String(),
		token.T_XOR_EQUAL.String(),
		token.T_MOD_EQUAL.String(),
		token.T_DEC.String(),
		token.T_INC.String(),
		token.T_DOUBLE_ARROW.String(),
		token.T_SPACESHIP.String(),
		token.T_IS_NOT_EQUAL.String(),
		token.T_IS_NOT_EQUAL.String(),
		token.T_IS_NOT_IDENTICAL.String(),
		token.T_IS_EQUAL.String(),
		token.T_IS_IDENTICAL.String(),
		token.T_SL_EQUAL.String(),
		token.T_SR_EQUAL.String(),
		token.T_IS_GREATER_OR_EQUAL.String(),
		token.T_IS_SMALLER_OR_EQUAL.String(),
		token.T_POW.String(),
		token.T_SL.String(),
		token.T_SR.String(),
		token.T_COALESCE.String(),

		token.ID(int(';')).String(),
		token.ID(int(':')).String(),
		token.ID(int(',')).String(),
		token.ID(int('.')).String(),
		token.ID(int('[')).String(),
		token.ID(int(']')).String(),
		token.ID(int('(')).String(),
		token.ID(int(')')).String(),
		token.ID(int('|')).String(),
		token.ID(int('/')).String(),
		token.ID(int('^')).String(),
		token.ID(int('&')).String(),
		token.ID(int('+')).String(),
		token.ID(int('-')).String(),
		token.ID(int('*')).String(),
		token.ID(int('=')).String(),
		token.ID(int('%')).String(),
		token.ID(int('!')).String(),
		token.ID(int('~')).String(),
		token.ID(int('$')).String(),
		token.ID(int('<')).String(),
		token.ID(int('>')).String(),
		token.ID(int('?')).String(),
		token.ID(int('@')).String(),
		token.ID(int('{')).String(),
		token.ID(int('}')).String(),

		token.T_VARIABLE.String(),
		token.T_STRING.String(),

		token.T_OBJECT_OPERATOR.String(),
		token.T_OBJECT_OPERATOR.String(),
		token.T_STRING.String(),

		token.T_ARRAY_CAST.String(),
		token.T_BOOL_CAST.String(),
		token.T_BOOL_CAST.String(),
		token.T_DOUBLE_CAST.String(),
		token.T_DOUBLE_CAST.String(),
		token.T_INT_CAST.String(),
		token.T_INT_CAST.String(),
		token.T_OBJECT_CAST.String(),
		token.T_STRING_CAST.String(),
		token.T_STRING_CAST.String(),
		token.T_UNSET_CAST.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestShebang(test *testing.T) {
	src := `#!/usr/bin/env php
<?php
0.1
`

	expected := []string{
		"#!/usr/bin/env php\n",
		"<?php",
		"\n",
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	tkn := lexer.Lex()
	assert.Equal(test, tkn.ID, token.T_DNUMBER)

	for _, tokenValue := range tkn.FreeFloating {
		actual = append(actual, string(tokenValue.Value))
	}

	assert.DeepEqual(test, expected, actual)
}

func TestShebangHtml(test *testing.T) {
	src := `#!/usr/bin/env php
<br/><?php
0.1
`

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()
	assert.Equal(test, tkn.ID, token.T_INLINE_HTML)
	assert.Equal(test, string(tkn.FreeFloating[0].Value), "#!/usr/bin/env php\n")

	tkn = lexer.Lex()
	assert.Equal(test, tkn.ID, token.T_DNUMBER)
}

func TestNumberTokens(test *testing.T) {
	src := `<?php
		0.1
		.1
		1e10
		.1e10

		0b01111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111
		0b10111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111

		0x0_7FFF_FFFF_FFFF_FFFF
		0x8111_1111_1111_1111

		92233_72036_85477_5807
		0_77777_77777_77777_77777_7

		92233_72036_85477_5808
		0_77777_77777_77777_77777_70
	`

	expected := []string{
		token.T_DNUMBER.String(),
		token.T_DNUMBER.String(),
		token.T_DNUMBER.String(),
		token.T_DNUMBER.String(),

		token.T_LNUMBER.String(),
		token.T_DNUMBER.String(),

		token.T_LNUMBER.String(),
		token.T_DNUMBER.String(),

		token.T_LNUMBER.String(),
		token.T_LNUMBER.String(),

		token.T_DNUMBER.String(),
		token.T_DNUMBER.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestConstantStrings(test *testing.T) {
	src := `<?
		'str'
		'\''
		'\\'

		b"str"
		"\""
		"\\"

		"\$var"
		"$4"
		"$"
		"$\\"

		"{"
		"{a"
		"\{$"
		"{\""
	`

	expected := []string{
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),

		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),

		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),

		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestSingleQuoteStringTokens(test *testing.T) {
	src := `<?php
		'str $var str'
		
		'\''
		
		'\'
		'
		
		'\
		\''
		
		'\\'
		
		'\\
		'

		'\	
		\''
	`

	expected := []string{
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestTeplateStringTokens(test *testing.T) {
	src := `<?php
		"foo $a"

		"foo $a{$b}"

		"test $var {$var} ${var_name} {s $ \$a "
		
		"{$var}"
		
		"$foo/"
		"$foo/100;"

		"$/$foo"
		"$0$foo"

		"$foo$"
	`

	expected := []string{
		token.ID(int('"')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_DOLLAR_OPEN_CURLY_BRACES.String(),
		token.T_STRING_VARNAME.String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('"')).String(),

		token.ID(int('"')).String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('"')).String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestBackquoteStringTokens(test *testing.T) {
	src := `<?php
		` + "`foo $a`" + `
		` + "`foo $a{$b}`" + `

		` + "`test $var {$var} ${var_name} {s $ \\$a `" + `
		
		` + "`{$var}`" + `
		` + "`$foo/`" + `
		` + "`$foo/100`" + `
		` + "`$/$foo`" + `
		` + "`$0$foo`" + `
		` + "`$foo$`" + `
	`

	expected := []string{
		token.ID(int('`')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_DOLLAR_OPEN_CURLY_BRACES.String(),
		token.T_STRING_VARNAME.String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('`')).String(),

		token.ID(int('`')).String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('`')).String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestHereDocTokens(test *testing.T) {
	src := `<?php
	<<<CAT
	test
CAT;

	<<<'CAT'
	test
CAT;

	<<<"CAT"
	$var->prop
	$var[1]
	$var[0x1]
	$var[0b1]
	$var[var_name]
	$var[$var]

	{$var}
	${var_name}
	{s $ \$a 
CAT;
	`

	expected := []string{
		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_OBJECT_OPERATOR.String(),
		token.T_STRING.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('[')).String(),
		token.T_NUM_STRING.String(),
		token.ID(int(']')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('[')).String(),
		token.T_NUM_STRING.String(),
		token.ID(int(']')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('[')).String(),
		token.T_NUM_STRING.String(),
		token.ID(int(']')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('[')).String(),
		token.T_STRING.String(),
		token.ID(int(']')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.ID(int('[')).String(),
		token.T_VARIABLE.String(),
		token.ID(int(']')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_DOLLAR_OPEN_CURLY_BRACES.String(),
		token.T_STRING_VARNAME.String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestHereDocTokens2(test *testing.T) {
	src := `<?php
	<<<CAT
$foo/
CAT;

	<<<CAT
$foo/100
CAT;

	<<<CAT
$/$foo
CAT;

	<<<CAT
$0$foo
CAT;

	<<<CAT
$foo$bar\
CAT
`

	expected := []string{
		token.T_START_HEREDOC.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),

		token.T_START_HEREDOC.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),

		token.T_START_HEREDOC.String(),
		token.T_VARIABLE.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestHereDocTokens3(test *testing.T) {
	src := `<?php

	<<<"CAT"
\\{$a['b']}
CAT;
	`

	expected := []string{

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_CURLY_OPEN.String(),
		token.T_VARIABLE.String(),
		token.ID(int('[')).String(),
		token.T_CONSTANT_ENCAPSED_STRING.String(),
		token.ID(int(']')).String(),
		token.ID(int('}')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestHereDocTokens73(test *testing.T) {
	src := `<?php
	<<<"CAT"
		text
	CAT, $b`

	expected := []string{

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(',')).String(),
		token.T_VARIABLE.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestHereDocTokensBefore73(test *testing.T) {
	src := `<?php
	<<<"CAT"
	CAT
CAT;`

	expected := []string{

		token.T_START_HEREDOC.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_END_HEREDOC.String(),
		token.ID(int(';')).String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 2,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestHereDocUnclosed(test *testing.T) {
	src := "<?<<<'S'\n"

	expected := []string{
		token.T_START_HEREDOC.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestInlineHtmlNopTokens(test *testing.T) {
	src := `<?php
		$a; ?> test <?php
		$a ?> test
	`

	expected := []string{
		token.T_VARIABLE.String(),
		token.ID(int(';')).String(),
		token.T_INLINE_HTML.String(),

		token.T_VARIABLE.String(),
		token.ID(int(';')).String(),
		token.T_INLINE_HTML.String(),
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
}

func TestStringTokensAfterVariable(test *testing.T) {
	src := `<?php "test \"$var\""`

	expected := []string{
		token.ID(int('"')).String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.T_VARIABLE.String(),
		token.T_ENCAPSED_AND_WHITESPACE.String(),
		token.ID(int('"')).String(),
	}

	expectedTokens := []string{
		"\"",
		"test \\\"",
		"$var",
		"\\\"",
		"\"",
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}
	actualTokens := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actualTokens = append(actualTokens, string(tkn.Value))
		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
	assert.DeepEqual(test, expectedTokens, actualTokens)
}

func TestSlashAfterVariable(test *testing.T) {
	src := `<?php $foo/3`

	expected := []string{
		token.T_VARIABLE.String(),
		token.ID(int('/')).String(),
		token.T_LNUMBER.String(),
	}

	expectedTokens := []string{
		"$foo",
		"/",
		"3",
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)
	actual := []string{}
	actualTokens := []string{}

	for {
		tkn := lexer.Lex()
		if tkn.ID == 0 {
			break
		}

		actualTokens = append(actualTokens, string(tkn.Value))
		actual = append(actual, tkn.ID.String())
	}

	assert.DeepEqual(test, expected, actual)
	assert.DeepEqual(test, expectedTokens, actualTokens)
}

func TestCommentEnd(test *testing.T) {
	src := `<?php //test`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("//test"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestCommentEnd2(test *testing.T) {
	src := `<?php //`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("//"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestCommentEnd3(test *testing.T) {
	src := `<?php #`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("#"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestAttribute(test *testing.T) {
	src := `<?php #[ FooAttribute]`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
		ErrorHandlerFunc: func(error *errors.Error) {
			println(error.Msg)
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)

	assert.DeepEqual(test, token.T_ATTRIBUTE, tkn.ID)
	assert.DeepEqual(test, "#[", string(tkn.Value))
}

func TestCommentNewLine(test *testing.T) {
	src := "<?php //test\n$a"

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("//test\n"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestCommentNewLine1(test *testing.T) {
	src := "<?php //test\r$a"

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("//test\r"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestCommentNewLine2(test *testing.T) {
	src := "<?php #test\r\n$a"

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("#test\r\n"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestCommentWithPhpEndTag(test *testing.T) {
	src := `<?php
	//test?> test`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("//test"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestInlineComment(test *testing.T) {
	src := `<?php
	/*test*/`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("/*test*/"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestInlineComment2(test *testing.T) {
	src := `<?php
	/*/*/`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("/*/*/"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestEmptyInlineComment(test *testing.T) {
	src := `<?php
	/**/ `

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
		{
			ID:    token.T_COMMENT,
			Value: []byte("/**/"),
		},
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestEmptyInlineComment2(test *testing.T) {
	src := `<?php
	/***/`

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},

		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
		{
			ID:    token.T_DOC_COMMENT,
			Value: []byte("/***/"),
		},
	}

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()

	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}

	assert.DeepEqual(test, expected, actual)
}

func TestMethodCallTokens(test *testing.T) {
	src := `<?php
	$a -> bar ( '' ) ;`

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},
		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
	}
	tkn := lexer.Lex()
	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token2 := range actual {
		token2.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token3 := range actual {
		token3.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token4 := range actual {
		token4.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token5 := range actual {
		token5.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token6 := range actual {
		token6.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token7 := range actual {
		token7.Position = nil
	}
	assert.DeepEqual(test, expected, actual)
}

func TestYieldFromTokens(test *testing.T) {
	src := `<?php
	yield from $a`

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	expected := []*token.Token{
		{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php"),
		},
		{
			ID:    token.T_WHITESPACE,
			Value: []byte("\n\t"),
		},
	}
	tkn := lexer.Lex()
	actual := tkn.FreeFloating
	for _, tokenValue := range actual {
		tokenValue.Position = nil
	}
	assert.DeepEqual(test, expected, actual)

	expected = []*token.Token{
		{
			ID:    token.T_WHITESPACE,
			Value: []byte(" "),
		},
	}
	tkn = lexer.Lex()
	actual = tkn.FreeFloating
	for _, token2 := range actual {
		token2.Position = nil
	}
	assert.DeepEqual(test, expected, actual)
}

func TestVarNameByteChars(test *testing.T) {
	src := "<?php $\x80 $\xff"

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()
	assert.Equal(test, "$\x80", string(tkn.Value))

	tkn = lexer.Lex()
	assert.Equal(test, "$\xff", string(tkn.Value))
}

func TestStringVarNameByteChars(test *testing.T) {
	src := "<?php \"$\x80 $\xff\""

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	tkn := lexer.Lex()
	assert.Equal(test, "\"", string(tkn.Value))

	tkn = lexer.Lex()
	assert.Equal(test, "$\x80", string(tkn.Value))

	tkn = lexer.Lex()
	assert.Equal(test, " ", string(tkn.Value))

	tkn = lexer.Lex()
	assert.Equal(test, "$\xff", string(tkn.Value))

	tkn = lexer.Lex()
	assert.Equal(test, "\"", string(tkn.Value))
}

func TestIgnoreControllCharacters(test *testing.T) {
	src := "<?php \004 echo $b;"

	var actualErr *errors.Error
	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
		ErrorHandlerFunc: func(error *errors.Error) {
			actualErr = error
		},
	}
	lexer := NewLexer([]byte(src), config)

	expected := "echo"
	tkn := lexer.Lex()
	actual := string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "$b"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expectedErr := &errors.Error{
		Msg: "WARNING: Unexpected character in input: '\x04' (ASCII=4)",
		Pos: &position.Position{StartLine: 1, EndLine: 1, StartPos: 6, EndPos: 7},
	}
	assert.DeepEqual(test, expectedErr, actualErr)
}

func TestIgnoreControllCharactersAtStringVarOffset(test *testing.T) {
	src := "<?php \"$a[test\004]\";"

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	expected := "\""
	tkn := lexer.Lex()
	actual := string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "$a"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "["
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "test"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "]"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)
}

func TestDoubleDollar(test *testing.T) {
	src := `<?php "$$a";`

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	expected := "\""
	tkn := lexer.Lex()
	actual := string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "$"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "$a"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)
}

func TestTripleDollar(test *testing.T) {
	src := `<?php "$$$a";`

	config := conf.Config{
		Version: &version.Version{
			Major: 7,
			Minor: 4,
		},
	}
	lexer := NewLexer([]byte(src), config)

	expected := "\""
	tkn := lexer.Lex()
	actual := string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "$$"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)

	expected = "$a"
	tkn = lexer.Lex()
	actual = string(tkn.Value)
	assert.DeepEqual(test, expected, actual)
}
