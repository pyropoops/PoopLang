package parser

import (
	"PoopLang/ast"
	"PoopLang/lexer"
	"PoopLang/token"
)

type Parser struct {
	input string
	currentToken token.Token
	l lexer.Lexer
	errors []string
}

func (p *Parser) ParseProgram() ast.Program{
	var statements []ast.Statement
	for p.currentToken.Type != token.EOF {
		if statement := p.parseStatement(); statement != nil {
			statements = append(statements, statement)
		}
		p.advanceToken()
	}
	return ast.Program{Statements: statements}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}
	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	pos := p.l.GetPosition()
	tok := p.currentToken
	p.advanceToken()
	ident := p.parseIdentifier()
	if ident == nil {
		p.throwParserError("Invalid 'let' statement - expected=identifier, got=" + p.currentToken.Type)
	}
	p.advanceToken()
	if p.currentToken.Type != token.ASSIGN {
		p.throwParserError("Invalid 'let' statement - expected: '=', got: " + p.currentToken.Type)
	}
	p.advanceToken()
	expr := p.parseExpression()
	return &ast.LetStatement{
		Tok:   tok,
		Ident: ident,
		Value: expr,
		Literal: p.input[pos:p.l.GetPosition()],
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	pos := p.l.GetPosition()
	tok := p.currentToken
	p.advanceToken()
	expr := p.parseExpression()
	return &ast.ReturnStatement{
		Tok:     tok,
		Value:   expr,
		Literal: p.input[pos:p.l.GetPosition()],
	}
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	if p.currentToken.Type != token.IDENTIFIER {
		p.advanceToken()
		return nil
	}
	ident := &ast.Identifier{Value: p.currentToken.Literal}
	p.advanceToken()
	return ident
}

func (p *Parser) advanceToken() {
	p.currentToken = p.l.NextToken()
}

func (p *Parser) throwParserError(err string) {
	p.errors = append(p.errors, "Parser Error: " + err)
}

func (p *Parser) parseExpression() *ast.Expression {
	return nil
}