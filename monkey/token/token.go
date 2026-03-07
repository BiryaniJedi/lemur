package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) PrintToken() {
	fmt.Printf("Token: { Type: %s, Literal: \"%s\" }\n", t.Type, t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, etc.
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "<-"
	PLUS   = "+"
	MINUS  = "-"

	// Comparison??
	LT = "<"
	LE = "<="
	EQ = "="
	GT = ">"
	GE = ">="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
