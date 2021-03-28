package ast

import "PoopLang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementLiteral() string
}

type Expression interface {
	Node
	ExpressionLiteral() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Tok     token.Token
	Ident   *Identifier
	Value   *Expression
	Literal string
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Tok.Literal
}

func (ls *LetStatement) StatementLiteral() string {
	return ls.Literal
}

type ReturnStatement struct {
	Tok     token.Token
	Value   *Expression
	Literal string
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Tok.Literal
}

func (rs *ReturnStatement) StatementLiteral() string {
	return rs.Literal
}

type Identifier struct {
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Value
}
