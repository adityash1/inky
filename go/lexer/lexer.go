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
		if ch == '+' {
			l.add_token(token.TOK_PLUS)
		}
		if ch == '-' {
			l.add_token(token.TOK_MINUS)
		}
		if ch == '*' {
			l.add_token(token.TOK_STAR)
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
	l.tokens = append(l.tokens, *token.NewToken(t, l.source[l.start:l.curr]))
}

// helper functions
// func (l *Lexer) Peek() {
// 	return
// }

// func (l *Lexer) Lookahead() {
// 	return
// }

// func (l *Lexer) Match() {
// 	return
// }
