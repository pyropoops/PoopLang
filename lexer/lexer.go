package lexer

import (
	"PoopLang/token"
	"strconv"
	"strings"
)

type Lexer struct {
	position int
	ch       byte
	input    string
	errors   []string
}

func (l *Lexer) advancePosition() {
	l.ch = l.peekChar()
	l.position++
}

func (l *Lexer) peekChar() byte {
	readPosition := l.position + 1
	if readPosition < len(l.input) {
		return l.input[readPosition]
	}
	return 0
}

func NewLexer(input string) Lexer {
	l := Lexer{
		position: -1,
		input:    input,
	}
	l.advancePosition()
	return l
}

func (l *Lexer) NextToken() token.Token {
	for isWhitespace(l.ch) {
		l.advancePosition()
	}

	tok := token.Token{Type: token.UNDEFINED, Literal: string(l.ch)}
	switch l.ch {
	case 0:
		tok.Type = token.EOF
	case '=':
		if l.peekChar() == '=' {
			tok.Type = token.EQ
			tok.Literal = string(l.ch) + string(l.peekChar())
			l.advancePosition()
		} else {
			tok.Type = token.ASSIGN
		}
	case '!':
		if l.peekChar() == '=' {
			tok.Type = token.NOT_EQ
			tok.Literal = string(l.ch) + string(l.peekChar())
			l.advancePosition()
		} else {
			tok.Type = token.BANG
		}
	default:
		tok := l.identifyToken()
		return tok
	}
	l.advancePosition()
	return tok
}

func (l *Lexer) identifyToken() token.Token {
	// Evaluate against l.ch
	if t, ok := token.Tokens[string(l.ch)]; ok {
		tok := token.Token{Type: t, Literal: string(l.ch)}
		l.advancePosition()
		return tok
	}

	tok := token.Token{Type: token.UNDEFINED}

	// If not, evaluate for identifier/keyword
	pos := l.position

	// String detections
	switch l.ch {
	case '\'', '`', '"':
		return l.identifyString(l.ch)
	}
	for {
		if isLetter(l.ch) {
			tok.Type = token.IDENTIFIER
		} else if tok.Type == token.UNDEFINED {
			if isDigit(l.ch) {
				tok.Type = token.INT
			} else {
				tok.Type = token.ILLEGAL
			}
		}

		if isLetter(l.ch) || isDigit(l.ch) {
			l.advancePosition()
		} else {
			if tok.Type == token.ILLEGAL {
				if l.position == pos {
					tok.Literal = string(l.input[pos])
				} else {
					tok.Literal = l.input[pos:l.position]
				}
				l.advancePosition()
				return tok
			}
			break
		}
	}

	if tok.Type == token.INT && l.ch == '.' {
		l.advancePosition()
		if isDigit(l.ch) {
			tok.Type = token.FLOAT
			for isDigit(l.ch) {
				l.advancePosition()
			}
		} else {
			tok.Type = token.ILLEGAL
		}
	}

	literal := l.input[pos:l.position]
	tok.Literal = literal

	if t, ok := token.Tokens[literal]; tok.Type == token.IDENTIFIER && ok {
		tok.Type = t
	}
	return tok
}

func (l *Lexer) identifyString(char byte) token.Token {
	l.advancePosition()
	pos := l.position
	for l.ch != char {
		// EOF check
		if l.peekChar() == 0 {
			l.throwLexerError("Incomplete string literal. Expected = " + string(char) + ", got = EOF")
			return token.Token{Type: token.ILLEGAL, Literal: string(char)}
		}
		l.advancePosition()
	}
	tok := token.Token{Type: token.STRING, Literal: l.input[pos:l.position]}
	l.advancePosition()
	return tok
}

func (l *Lexer) throwLexerError(err string) {
	line, column := getLineColumnNumber(l.input, l.position)
	l.errors = append(l.errors, "line "+line+", column "+column+": "+err)
	l.ch = 0
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// @return line, collumn
func getLineColumnNumber(input string, index int) (string, string) {
	arr := strings.Split(input[:index], "\n")
	line := len(arr)
	column := len(arr[len(arr)-1])
	return strconv.Itoa(line), strconv.Itoa(column)
}

func (l *Lexer) GetPosition() int {
	return l.position
}
