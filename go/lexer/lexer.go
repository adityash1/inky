package lexer

import "blue/token"

type Lexer struct {
	tokens []token.Token
	source []byte
	start  int
	curr   int
	line   int
}

func NewLexer(source []byte) *Lexer {
	return &Lexer{
		tokens: []token.Token{},
		source: source,
		curr:   0,
		start:  0,
		line:   1,
	}
}

func (l *Lexer) Tokenize() []token.Token {
	for l.curr < len(l.source) {
		l.start = l.curr
		ch := l.advance()
		if ch == '\n' {
			l.line = l.line + 1
		} else if ch == ' ' || ch == '\t' || ch == '\r' {
			continue
		} else if ch == '#' {
			for l.peek() != '\n' {
				l.advance()
			}
		} else if ch == '(' {
			l.add_token(token.TOK_LPAREN)
		} else if ch == ')' {
			l.add_token(token.TOK_RPAREN)
		} else if ch == '{' {
			l.add_token(token.TOK_LCURLY)
		} else if ch == '}' {
			l.add_token(token.TOK_RCURLY)
		} else if ch == '[' {
			l.add_token(token.TOK_LSQUAR)
		} else if ch == ']' {
			l.add_token(token.TOK_RSQUAR)
		} else if ch == ',' {
			l.add_token(token.TOK_COMMA)
		} else if ch == '.' {
			l.add_token(token.TOK_DOT)
		} else if ch == '+' {
			l.add_token(token.TOK_PLUS)
		} else if ch == '-' {
			l.add_token(token.TOK_MINUS)
		} else if ch == '*' {
			l.add_token(token.TOK_STAR)
		} else if ch == '/' {
			l.add_token(token.TOK_SLASH)
		} else if ch == '^' {
			l.add_token(token.TOK_CARET)
		} else if ch == '%' {
			l.add_token(token.TOK_MOD)
		} else if ch == ':' {
			l.add_token(token.TOK_COLON)
		} else if ch == ';' {
			l.add_token(token.TOK_SEMICOLON)
		} else if ch == '?' {
			l.add_token(token.TOK_QUESTION)
		} else if ch == '~' {
			l.add_token(token.TOK_NOT)
		} else if ch == '>' {
			l.add_token(token.TOK_GT)
		} else if ch == '<' {
			l.add_token(token.TOK_LT)
		}
	}
	return l.tokens
}

func (l *Lexer) advance() byte {
	ch := l.source[l.curr]
	l.curr++
	return ch
}

func (l *Lexer) add_token(t token.TokenType) {
	l.tokens = append(l.tokens, *token.NewToken(t, string(l.source[l.start:l.curr]), l.line))
}

func (l *Lexer) peek() byte {
	return l.source[l.curr]
}

func (l *Lexer) lookahead(n ...int) byte {
	count := 1
	if len(n) > 0 {
		count = n[0]
	}
	return l.source[l.curr+count]
}

func (l *Lexer) match(expected byte) bool {
	if l.curr >= len(l.source) {
		return false
	}
	if l.source[l.curr] != expected {
		return false
	}
	l.curr++
	return true
}
