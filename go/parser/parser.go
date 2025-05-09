package parser

import (
	"blue/model"
	"blue/token"
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

func (p *Parser) Parse() model.Expr {
	ast := p.expr()
	return ast
}

func (p *Parser) expr() model.Expr {
	expr := p.term()
	for p.match(token.TOK_PLUS) || p.match(token.TOK_MINUS) {
		op := p.previousToken()
		right := p.term()
		expr = &model.BinOp{Op: op, Left: expr, Right: right}
	}
	return expr
}

func (p *Parser) unary() model.Expr {
	if p.match(token.TOK_NOT) || p.match(token.TOK_MINUS) || p.match(token.TOK_PLUS) {
		op := p.previousToken()
		operand := p.unary()
		return &model.UnOp{Op: op, Operand: operand}
	}
	return p.primary()
}

func (p *Parser) primary() model.Expr {
	if p.match(token.TOK_INTEGER) {
		val, _ := strconv.Atoi(p.previousToken().Lexeme)
		return &model.Integer{Value: val}
	}
	if p.match(token.TOK_FLOAT) {
		val, _ := strconv.ParseFloat(p.previousToken().Lexeme, 64)
		return &model.Float{Value: val}
	}
	if p.match(token.TOK_LPAREN) {
		expr := p.expr()
		if !p.match(token.TOK_RPAREN) {
			panic("Error: ')' expected.")
		}
		return &model.Grouping{Value: expr}
	}
	return nil
}

func (p *Parser) factor() model.Expr {
	return p.unary()
}

func (p *Parser) term() model.Expr {
	expr := p.factor()
	for p.match(token.TOK_STAR) || p.match(token.TOK_SLASH) {
		op := p.previousToken()
		right := p.factor()
		expr = &model.BinOp{Op: op, Left: expr, Right: right}
	}
	return expr
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

func (p *Parser) advance() token.Token {
	tok := p.tokens[p.curr]
	p.curr++
	return tok
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.curr]
}

func (p *Parser) previousToken() token.Token {
	return p.tokens[p.curr-1]
}
