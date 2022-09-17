package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	PxLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: `EOE`, Pattern: `;`},
		{Name: `EOK`, Pattern: `=`},
		{Name: `String`, Pattern: `"(?:\\.|[^"])*"`},
		{Name: `Ident`, Pattern: `[a-zA-Z][a-zA-Z-_\d]*`},
		{Name: `Integer`, Pattern: `\d+`},
		{Name: `Decimal`, Pattern: `\d*\.\d+`},
		{Name: `Punct`, Pattern: `[][\-(),"]`},
		{Name: `EOL`, Pattern: `[\n\r]+`},
		{Name: `whitespace`, Pattern: `\s+`},
	})
	PxParser = participle.MustBuild[PxFileHeader](
		participle.Lexer(PxLexer),
		participle.Unquote("String"),
		participle.Elide("whitespace", "EOL"),
	)
)

type PxFileHeader struct {
	Rows []PxRow `parser:"( @@ )* 'DATA' '=' "`
}

type PxRow struct {
	Keyword PxKeyword `parser:" @@ '=' "`
	Value   PxValue   `parser:" @@ ';' "`
}

type PxKeyword struct {
	Keyword    string    `parser:" ( (?! 'DATA' ) @Ident )! "`
	Language   *string   `parser:"( '[' @Ident ']' )?"`
	Specifiers *[]string `parser:"( '(' @String ( ',' @String )* ')' )?"`
}

type PxValue struct {
	Integer *int           `parser:"   @Integer"`
	Times   *[]PxTimeVal   `parser:"  | @@ (',' @@)* "`
	String  *string        `parser:"  | @Ident "`
	List    *[]PxStringVal `parser:"  | @@ (',' @@)* "`
}

type PxTimeVal struct {
	Units string    `parser:" 'TLIST' '(' @( 'A1' | 'H1' | 'Q1' | 'M1' | 'W1' ) "`
	Range *[]string `parser:" ( ',' @String '-' @String )? ')' "`
	Times *[]string `parser:" ( ',' @String )* "`
}

type PxStringVal struct {
	Strings []string `parser:" ( @String )* "`
}
