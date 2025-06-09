package parser

import (
	"inky/ast"
	"inky/token"
	"inky/utils"
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

func (p *Parser) Parse() ast.Node {
	ast := p.program()
	return ast
}

// program  ::= stmts
func (p *Parser) program() ast.Node {
	ast := p.stmts()
	return ast
}

// stmts ::= stmt+
func (p *Parser) stmts() *ast.Stmts {
	stmts := []ast.Stmt{}
	//TODO: when to stop parsing a statement?
	for p.curr < len(p.tokens) {
		stmt := p.stmt()
		stmts = append(stmts, stmt)
	}
	return &ast.Stmts{Stmts: stmts, Line: p.previousToken().Line}
}

// stmt ::= expr_stmt | print_stmt | assign | local_assign | println_stmt |
// if_stmt | while_stmt | for_stmt | func_decl | func_call | ret_stmt
func (p *Parser) stmt() ast.Stmt {
	// TODO: predictive parsing, where the next token predicts what is the next statement
	// TODO: parse print, if, while, for, assignment, function call, etc.
	if p.peek().Type == token.TOK_PRINT {
		return p.print_stmt("")
	} else if p.peek().Type == token.TOK_PRINTLN {
		return p.print_stmt("\n")
	}
	// else if p.peek().Type == token.TOK_IF {
	// 	return p.if_stmt()
	// } else if p.peek().Type == token.TOK_WHILE {
	// 	return p.while_stmt()
	// } else if p.peek().Type == token.TOK_FOR {
	// 	return p.for_stmt()
	// } else if p.peek().Type == token.TOK_FUNC {
	// 	return p.func_decl()
	// } else {
	// 	// TODO
	// }
	return nil
}

// print_stmt ::= 'print' expr
func (p *Parser) print_stmt(end string) ast.Stmt {
	if p.match(token.TOK_PRINT) || p.match(token.TOK_PRINTLN) {
		val := p.expr()
		return &ast.PrintStmt{Value: val, End: end, Line: p.previousToken().Line}
	}
	return nil
}

// expr ::= or_logical
func (p *Parser) expr() ast.Expr {
	return p.or_logical()
}

// or_logical ::= and_logical ( 'or' and_logical )*
func (p *Parser) or_logical() ast.Expr {
	expr := p.and_logical()
	for p.match(token.TOK_OR) {
		op := p.previousToken()
		right := p.and_logical()
		expr = &ast.LogicalOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// and_logical ::= equality ( 'and' equality )*
func (p *Parser) and_logical() ast.Expr {
	expr := p.equality()
	for p.match(token.TOK_AND) {
		op := p.previousToken()
		right := p.equality()
		expr = &ast.LogicalOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// equality ::= comparison ( ( '~=' | '==' ) comparison )*
func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.TOK_EQEQ) || p.match(token.TOK_NE) {
		op := p.previousToken()
		right := p.comparison()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// comparison ::= addition ( ( '>' | '>=' | '<' | '<=' ) addition )*
func (p *Parser) comparison() ast.Expr {
	expr := p.addition()
	for p.match(token.TOK_GT) || p.match(token.TOK_GE) || p.match(token.TOK_LT) || p.match(token.TOK_LE) {
		op := p.previousToken()
		right := p.addition()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// addition ::= multiplication ( ( '+' | '-' ) multiplication )*
func (p *Parser) addition() ast.Expr {
	expr := p.multiplication()
	for p.match(token.TOK_PLUS) || p.match(token.TOK_MINUS) {
		op := p.previousToken()
		right := p.multiplication()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// multiplication ::= modulo ( ( '*' | '/' ) modulo )*
func (p *Parser) multiplication() ast.Expr {
	expr := p.modulo()
	for p.match(token.TOK_STAR) || p.match(token.TOK_SLASH) {
		op := p.previousToken()
		right := p.modulo()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// modulo ::= unary ( '%' unary )*
func (p *Parser) modulo() ast.Expr {
	expr := p.unary()
	for p.match(token.TOK_MOD) {
		op := p.previousToken()
		right := p.unary()
		expr = &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// unary ::= ( '~' | '-' | '+' )* exponent
func (p *Parser) unary() ast.Expr {
	if p.match(token.TOK_NOT) || p.match(token.TOK_MINUS) || p.match(token.TOK_PLUS) {
		op := p.previousToken()
		operand := p.unary()
		return &ast.UnOp{Op: op, Operand: operand, Line: op.Line}
	}
	return p.exponent()
}

// exponent ::= primary ( '^' exponent )*
func (p *Parser) exponent() ast.Expr {
	expr := p.primary()
	if p.match(token.TOK_CARET) {
		op := p.previousToken()
		right := p.exponent() // Recursively parse the right side for right-associativity
		return &ast.BinOp{Op: op, Left: expr, Right: right, Line: op.Line}
	}
	return expr
}

// ‹primary> ::= <integer> | ‹float> | '(' ‹expr> ')' | <bool> | <string>
func (p *Parser) primary() ast.Expr {
	if p.match(token.TOK_INTEGER) {
		val, _ := strconv.Atoi(p.previousToken().Lexeme)
		return &ast.Integer{Value: val, Line: p.previousToken().Line}
	}
	if p.match(token.TOK_FLOAT) {
		val, _ := strconv.ParseFloat(p.previousToken().Lexeme, 64)
		return &ast.Float{Value: val, Line: p.previousToken().Line}
	}
	if p.match(token.TOK_TRUE) {
		return &ast.Bool{Value: true, Line: p.previousToken().Line}
	}
	if p.match(token.TOK_FALSE) {
		return &ast.Bool{Value: false, Line: p.previousToken().Line}
	}
	if p.match(token.TOK_STRING) {
		return &ast.String{Value: p.previousToken().Lexeme[1 : len(p.previousToken().Lexeme)-1], Line: p.previousToken().Line} // Remove the quotes from the string
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

func (p *Parser) peek() token.Token {
	return p.tokens[p.curr]
}

// func (p *Parser) expect(expectedType token.TokenType) token.Token {
// 	if p.curr >= len(p.tokens) {
// 		utils.ParseError(fmt.Sprintf("Found %v at the end of parsing", p.previousToken().Lexeme), p.previousToken().Line)
// 	} else if p.tokens[p.curr].Type == expectedType {
// 		return p.advance()
// 	} else {
// 		utils.ParseError(fmt.Sprintf("Expected %v, found %v.", expectedType, p.peek().Lexeme), p.peek().Line)
// 	}
// 	return token.Token{} // unreachable, but required
// }

// func (p *Parser) isNext(expectedType token.TokenType) bool {
// 	if p.curr >= len(p.tokens) {
// 		return false
// 	}
// 	return p.peek().Type == expectedType
// }

// func (p *Parser) advance() token.Token {
// 	tok := p.tokens[p.curr]
// 	p.curr++
// 	return tok
// }
