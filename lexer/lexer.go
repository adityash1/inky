package lexer

import (
	"blue/token"
	"blue/utils"
	"fmt"
	"unicode"
)

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
			for l.peek() != '\n' && !(l.curr >= len(l.source)) {
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
			if l.match('-') {
				for l.peek() != '\n' && !(l.curr >= len(l.source)) {
					l.advance()
				}
			} else {
				l.add_token(token.TOK_MINUS)
			}
		} else if ch == '*' {
			l.add_token(token.TOK_STAR)
		} else if ch == '/' {
			l.add_token(token.TOK_SLASH)
		} else if ch == '^' {
			l.add_token(token.TOK_CARET)
		} else if ch == '%' {
			l.add_token(token.TOK_MOD)
		} else if ch == ':' {
			l.match('=')
			if l.match('=') {
				l.add_token(token.TOK_ASSIGN)
			} else {
				l.add_token(token.TOK_COLON)
			}
		} else if ch == ';' {
			l.add_token(token.TOK_SEMICOLON)
		} else if ch == '?' {
			l.add_token(token.TOK_QUESTION)
		} else if ch == '>' {
			if l.match('=') {
				l.add_token(token.TOK_GE)
			} else if l.match('>') {
				l.add_token(token.TOK_GTGT)
			} else {
				l.add_token(token.TOK_GT)
			}
		} else if ch == '<' {
			if l.match('=') {
				l.add_token(token.TOK_LE)
			} else if l.match('<') {
				l.add_token(token.TOK_LTLT)
			} else {
				l.add_token(token.TOK_LT)
			}
		} else if ch == '=' {
			if l.match('=') {
				l.add_token(token.TOK_EQ)
			}
		} else if ch == '~' {
			if l.match('=') {
				l.add_token(token.TOK_NE)
			} else {
				l.add_token(token.TOK_NOT)
			}
		} else if isDigit(ch) {
			l.handleNumber()
		} else if ch == '"' || ch == '\'' {
			l.handleString(ch)
		} else if unicode.IsLetter(rune(ch)) || ch == '_' {
			l.handleIdentifier()
		} else {
			utils.LexingError(fmt.Sprintf("Error at %s: Unexpected character.", string(ch)), l.line)
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
	if l.curr >= len(l.source) {
		return '\x00' // null byte, safe end-of-stream marker
	}
	return l.source[l.curr]
}

func (l *Lexer) lookahead(n ...int) byte {
	count := 1
	if len(n) > 0 {
		count = n[0]
	}
	if l.curr+count >= len(l.source) {
		return '\x00'
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

func (l *Lexer) handleNumber() {
	for isDigit(l.peek()) {
		l.advance()
	}
	if l.peek() == '.' && isDigit(l.lookahead()) {
		l.advance()
		for isDigit(l.peek()) {
			l.advance()
		}
		l.add_token(token.TOK_FLOAT)
	} else {
		l.add_token(token.TOK_INTEGER)
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) handleString(start_quote byte) {
	for l.peek() != start_quote && !(l.curr >= len(l.source)) {
		l.advance()
	}
	if l.curr >= len(l.source) {
		utils.LexingError("Unterminated string.", l.line)
	}
	l.advance()
	l.add_token(token.TOK_STRING)
}

func (l *Lexer) handleIdentifier() {
	for (unicode.IsLetter(rune(l.peek())) || unicode.IsDigit(rune(l.peek()))) || l.peek() == '_' {
		l.advance()
	}
	text := l.source[l.start:l.curr]
	keyword_type := token.Keywords[string(text)]
	if keyword_type == "" {
		l.add_token(token.TOK_IDENTIFIER)
	} else {
		l.add_token(keyword_type)
	}
}
