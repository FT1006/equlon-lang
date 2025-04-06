package lexer

import (
	"github.com/spaceship/equlon/token"
)

// Lexer represents a lexical scanner.
type Lexer struct {
	input 		 string
	position 	 int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch 			 byte // current char under examination
}

// New returns a new Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l

}

// readChar initializes the Lexer's current character.
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0 // indicates end of input
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

// NextToken returns the next token.
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
	case ':':
		l.readChar()
		if l.input[l.position] == '=' {
			tok = newToken(token.EQULON, ':')
		} else {
			tok = newToken(token.ILLEGAL, ':')
		}
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
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
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }

    l.readChar()
    return tok
}

// newToken is a helper function to create a new token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	if ch == ':' {
		return token.Token{Type: tokenType, Literal: ":="}
	}
    return token.Token{Type: tokenType, Literal: string(ch)}
}