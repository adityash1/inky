package token

import "fmt"

type TokenType string

const (
	// Single-character tokens
	TOK_LPAREN    TokenType = "TOK_LPAREN"    // (
	TOK_RPAREN    TokenType = "TOK_RPAREN"    // )
	TOK_LCURLY    TokenType = "TOK_LCURLY"    // {
	TOK_RCURLY    TokenType = "TOK_RCURLY"    // }
	TOK_LSQUAR    TokenType = "TOK_LSQUAR"    // [
	TOK_RSQUAR    TokenType = "TOK_RSQUAR"    // ]
	TOK_COMMA     TokenType = "TOK_COMMA"     // ,
	TOK_DOT       TokenType = "TOK_DOT"       // .
	TOK_PLUS      TokenType = "TOK_PLUS"      // +
	TOK_MINUS     TokenType = "TOK_MINUS"     // -
	TOK_STAR      TokenType = "TOK_STAR"      // *
	TOK_SLASH     TokenType = "TOK_SLASH"     // /
	TOK_CARET     TokenType = "TOK_CARET"     // ^
	TOK_MOD       TokenType = "TOK_MOD"       // %
	TOK_COLON     TokenType = "TOK_COLON"     // :
	TOK_SEMICOLON TokenType = "TOK_SEMICOLON" // ;
	TOK_QUESTION  TokenType = "TOK_QUESTION"  // ?
	TOK_NOT       TokenType = "TOK_NOT"       // ~
	TOK_GT        TokenType = "TOK_GT"        // >
	TOK_LT        TokenType = "TOK_LT"        // <

	// Two-character tokens
	TOK_GE     TokenType = "TOK_GE"     // >=
	TOK_LE     TokenType = "TOK_LE"     // <=
	TOK_NE     TokenType = "TOK_NE"     // ~=
	TOK_EQ     TokenType = "TOK_EQ"     // ==
	TOK_ASSIGN TokenType = "TOK_ASSIGN" // :=
	TOK_GTGT   TokenType = "TOK_GTGT"   // >>
	TOK_LTLT   TokenType = "TOK_LTLT"   // <<

	// Literals
	TOK_IDENTIFIER TokenType = "TOK_IDENTIFIER"
	TOK_STRING     TokenType = "TOK_STRING"
	TOK_INTEGER    TokenType = "TOK_INTEGER"
	TOK_FLOAT      TokenType = "TOK_FLOAT"

	// Keywords
	TOK_IF      TokenType = "TOK_IF"
	TOK_THEN    TokenType = "TOK_THEN"
	TOK_ELSE    TokenType = "TOK_ELSE"
	TOK_TRUE    TokenType = "TOK_TRUE"
	TOK_FALSE   TokenType = "TOK_FALSE"
	TOK_AND     TokenType = "TOK_AND"
	TOK_OR      TokenType = "TOK_OR"
	TOK_WHILE   TokenType = "TOK_WHILE"
	TOK_DO      TokenType = "TOK_DO"
	TOK_FOR     TokenType = "TOK_FOR"
	TOK_FUNC    TokenType = "TOK_FUNC"
	TOK_NULL    TokenType = "TOK_NULL"
	TOK_END     TokenType = "TOK_END"
	TOK_PRINT   TokenType = "TOK_PRINT"
	TOK_PRINTLN TokenType = "TOK_PRINTLN"
	TOK_RET     TokenType = "TOK_RET"
)

var Keywords = map[string]TokenType{
	"if":      TOK_IF,
	"then":    TOK_THEN,
	"else":    TOK_ELSE,
	"true":    TOK_TRUE,
	"false":   TOK_FALSE,
	"and":     TOK_AND,
	"or":      TOK_OR,
	"while":   TOK_WHILE,
	"do":      TOK_DO,
	"for":     TOK_FOR,
	"func":    TOK_FUNC,
	"null":    TOK_NULL,
	"end":     TOK_END,
	"print":   TOK_PRINT,
	"println": TOK_PRINTLN,
	"ret":     TOK_RET,
}

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
}

func NewToken(token_type TokenType, lexeme string, line int) *Token {
	return &Token{
		Type:   token_type,
		Lexeme: lexeme,
		Line:   line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("(%s, %q, %d)", t.Type, t.Lexeme, t.Line)
}
