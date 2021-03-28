package token

type Token struct {
	Type    string
	Literal string
}

const (
	EOF        = "EOF"
	ILLEGAL    = "ILLEGAL"
	INT        = "INT"
	FLOAT      = "FLOAT"
	STRING     = "STRING"
	UNDEFINED  = "UNDEFINED"
	IDENTIFIER = "IDENTIFIER"
	ASSIGN     = "ASSIGN"
	PLUS       = "PLUS"
	MINUS      = "MINUS"
	MUL        = "MUL"
	DIV        = "DIV"
	SEMICOLON  = "SEMICOLON"
	COMMA      = "COMMA"
	BANG       = "BANG"
	EQ         = "EQ"
	NOT_EQ     = "NOT_EQ"
	LET        = "LET"
	FOR        = "FOR"
	IF         = "IF"
	ELSE       = "ELSE"
	LPAREN     = "LPAREN"
	RPAREN     = "RPAREN"
	LBRACE     = "LBRACE"
	RBRACE     = "RBRACE"
	TRUE       = "TRUE"
	FALSE      = "FALSE"
	RETURN     = "RETURN"
)

var Tokens = map[string]string{
	"+":     PLUS,
	"-":     MINUS,
	"*":     MUL,
	"/":     DIV,
	";":     SEMICOLON,
	",":     COMMA,
	"!":     BANG,
	"=":     ASSIGN,
	"let":   LET,
	"for":   FOR,
	"if":    IF,
	"else":  ELSE,
	"(":     LPAREN,
	")":     RPAREN,
	"{":     LBRACE,
	"}":     RBRACE,
	"true":  TRUE,
	"false": FALSE,
	"return": RETURN,
}
