package parser

import (
	"blue/ast"
	"blue/token"
	"blue/utils"
	"fmt"
	"strconv"
)

type Parser struct {
	tokens []token.Token
	curr   int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens: tokens,
		curr:   0,
	}
}

func (p *Parser) Parse() ast.Expr {
	ast := p.expr()
	return ast
}

func (p *Parser) expr() ast.Expr {
	return p.addition()
}

// <addition> ::= <multiplication> ( ('+'|'-') <multiplication> )*
func (p *Parser) addition() ast.Expr {
	expr := p.multiplication()
	for p.match(token.TOK_PLUS) || p.match(token.TOK_MINUS) {
		op := p.previousToken()
		right := p.multiplication()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// <multiplication> ::= <unary> ( ('*'|'/') <unary> )*
func (p *Parser) multiplication() ast.Expr {
	expr := p.unary()
	for p.match(token.TOK_STAR) || p.match(token.TOK_SLASH) {
		op := p.previousToken()
		right := p.unary()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// ‹unary> ::= ('+'|'-'|'~') ‹unary› | <primary>
func (p *Parser) unary() ast.Expr {
	if p.match(token.TOK_NOT) || p.match(token.TOK_MINUS) || p.match(token.TOK_PLUS) {
		op := p.previousToken()
		operand := p.unary()
		return &ast.UnOp{Op: op, Operand: operand, Line: op.Line}
	}
	return p.primary()
}

// ‹primary> ::= <integer> | ‹float> | '(' ‹expr> ')'
func (p *Parser) primary() ast.Expr {
	if p.match(token.TOK_INTEGER) {
		val, _ := strconv.Atoi(p.previousToken().Lexeme)
		return &ast.Integer{Value: val, Line: p.previousToken().Line}
	}
	if p.match(token.TOK_FLOAT) {
		val, _ := strconv.ParseFloat(p.previousToken().Lexeme, 64)
		return &ast.Float{Value: val, Line: p.previousToken().Line}
	}
	if p.match(token.TOK_LPAREN) {
		expr := p.expr()
		if !p.match(token.TOK_RPAREN) {
			utils.ParseError("Error: ')' expected.", p.previousToken().Line)
		}
		return &ast.Grouping{Value: expr, Line: p.previousToken().Line}
	}
	return nil
}

// Utility methods
func (p *Parser) match(expectedType token.TokenType) bool {
	if p.curr >= len(p.tokens) {
		return false
	}
	if p.peek().Type != expectedType {
		return false
	}
	p.curr++
	return true
}

func (p *Parser) previousToken() token.Token {
	return p.tokens[p.curr-1]
}

func (p *Parser) expect(expectedType token.TokenType) token.Token {
	if p.curr >= len(p.tokens) {
		utils.ParseError(fmt.Sprintf("Found %v at the end of parsing", p.previousToken().Lexeme), p.previousToken().Line)
	} else if p.tokens[p.curr].Type == expectedType {
		return p.advance()
	} else {
		utils.ParseError(fmt.Sprintf("Expected %v, found %v.", expectedType, p.peek().Lexeme), p.peek().Line)
	}
	return token.Token{} // unreachable, but required
}

func (p *Parser) isNext(expectedType token.TokenType) bool {
	if p.curr >= len(p.tokens) {
		return false
	}
	return p.peek().Type == expectedType
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.curr]
}

func (p *Parser) advance() token.Token {
	tok := p.tokens[p.curr]
	p.curr++
	return tok
}
