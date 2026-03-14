package lexer

import (
	"lemur/token"
)

type Lexer struct {
	input        string
	prevPosition int  // previous position
	position     int  // current position in input (points to current char)
	nextPosition int  // current reading position in input (after current char)
	ch           byte // curent char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextChar()
	l.prevPosition = -1
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '<':
		tok = l.lessThanSwitch()
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.FSLASH, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '=':
		tok = l.equalsSwitch()
	case '!':
		tok = l.bangSwitch()
	case '>':
		tok = l.greaterThanSwitch()
	case '&':
		if l.peekChar() == '&' {
			l.nextChar()
			tok = token.Token{Type: token.AND, Literal: "&&"}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			l.nextChar()
			tok = token.Token{Type: token.OR, Literal: "||"}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '"':
		tok = l.getStringToken()
	case 0:
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readIdentifier(isDigit)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.nextChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenStr(tokenType token.TokenType, strLiter string) token.Token {
	return token.Token{Type: tokenType, Literal: strLiter}
}

func (l *Lexer) nextChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.prevPosition = l.position
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return l.input[l.nextPosition]
}

func (l *Lexer) prevChar() byte {
	if l.prevPosition < 0 {
		return 0
	}
	return l.input[l.prevPosition]
}

func (l *Lexer) lessThanSwitch() token.Token {
	nextChar := l.peekChar()
	switch nextChar {
	case 0:
		return newToken(token.ILLEGAL, l.ch)
	case '-':
		l.nextChar()
		return newTokenStr(token.ASSIGN, "<-")
	case '=':
		l.nextChar()
		return newTokenStr(token.LE, "<=")
	default:
		return newToken(token.LT, l.ch)
	}
}

func (l *Lexer) greaterThanSwitch() token.Token {
	nextChar := l.peekChar()
	switch nextChar {
	case 0:
		return newToken(token.ILLEGAL, l.ch)
	case '=':
		l.nextChar()
		return newTokenStr(token.GE, ">=")
	default:
		return newToken(token.GT, l.ch)
	}
}

func (l *Lexer) bangSwitch() token.Token {
	nextChar := l.peekChar()
	switch nextChar {
	case 0:
		return newToken(token.ILLEGAL, l.ch)
	case '=':
		l.nextChar()
		return newTokenStr(token.NEQ, "!=")
	default:
		return newToken(token.BANG, l.ch)
	}
}

func (l *Lexer) equalsSwitch() token.Token {
	nextChar := l.peekChar()
	switch nextChar {
	case 0:
		return newToken(token.ILLEGAL, l.ch)
	case '>':
		l.nextChar()
		return newTokenStr(token.ARROW, "=>")
	default:
		return newToken(token.EQ, l.ch)
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextChar()
	}
}

func (l *Lexer) getStringToken() token.Token {
	// current state '"' character
	position := l.position + 1
	l.nextChar()
	for l.ch != '"' {
		l.nextChar()
		if l.ch == 0 {
			return newToken(token.ILLEGAL, l.ch)
		}
	}
	tok := newTokenStr(token.STRING, l.input[position:l.position])
	return tok
}

func (l *Lexer) readIdentifier(charTypeChecker func(c byte) bool) string {
	position := l.position
	for charTypeChecker(l.ch) {
		l.nextChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
