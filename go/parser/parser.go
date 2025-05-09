package parser

import (
	"blue/ast"
	"blue/token"
	"fmt"
	"os"
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

// ‹expr> ::= ‹term> ( ('+'|'-') ‹term› )*
func (p *Parser) expr() ast.Expr {
	expr := p.term()
	for p.match(token.TOK_PLUS) || p.match(token.TOK_MINUS) {
		op := p.previousToken()
		right := p.term()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right}
	}
	return expr
}

// <term> ::= ‹factor> ( ('*'|'/') ‹factor> )*
func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(token.TOK_STAR) || p.match(token.TOK_SLASH) {
		op := p.previousToken()
		right := p.factor()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right}
	}
	return expr
}

// <factor> ::= <unary>
func (p *Parser) factor() ast.Expr {
	return p.unary()
}

// ‹unary> ::= ('*'|'-'|'~') ‹unary› | <primary>
func (p *Parser) unary() ast.Expr {
	if p.match(token.TOK_NOT) || p.match(token.TOK_MINUS) || p.match(token.TOK_PLUS) {
		op := p.previousToken()
		operand := p.unary()
		return &ast.UnOp{Op: op, Operand: operand}
	}
	return p.primary()
}

// ‹primary> ::= <integer> | ‹float> | '(' ‹expr> ')'
func (p *Parser) primary() ast.Expr {
	if p.match(token.TOK_INTEGER) {
		val, _ := strconv.Atoi(p.previousToken().Lexeme)
		return &ast.Integer{Value: val}
	}
	if p.match(token.TOK_FLOAT) {
		val, _ := strconv.ParseFloat(p.previousToken().Lexeme, 64)
		return &ast.Float{Value: val}
	}
	if p.match(token.TOK_LPAREN) {
		expr := p.expr()
		if !p.match(token.TOK_RPAREN) {
			die("Error: ')' expected.")
		}
		return &ast.Grouping{Value: expr}
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
		die(fmt.Sprintf("Found %v at the end of parsing", p.previousToken().Lexeme))
	} else if p.tokens[p.curr].Type == expectedType {
		return p.advance()
	} else {
		die(fmt.Sprintf("Expected %v, found %v.", expectedType, p.peek().Lexeme))
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

func die(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
