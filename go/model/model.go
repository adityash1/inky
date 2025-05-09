package model

import (
	"blue/token"
	"fmt"
)

// Expr is the interface for expressions
type Expr interface {
	String() string
}

// Stat is the interface for statements.
type Stat interface {
	String() string
}

// Integer represents an integer expression.
type Integer struct {
	Value int
}

func (i Integer) String() string {
	return fmt.Sprintf("Integer[%d]", i.Value)
}

// Float represents a float expression.
type Float struct {
	Value float64
}

func (f Float) String() string {
	return fmt.Sprintf("Float[%f]", f.Value)
}

// BinOp represents a binary operation like x + y.
type BinOp struct {
	Op    token.Token
	Left  Expr
	Right Expr
}

func (b BinOp) String() string {
	return fmt.Sprintf("BinOp[%q, %s, %s]", b.Op.Lexeme, b.Left.String(), b.Right.String())
}

// UnOp represents a unary operation like -x.
type UnOp struct {
	Op      token.Token
	Operand Expr
}

func (u UnOp) String() string {
	return fmt.Sprintf("UnOp[%q, %s]", u.Op.Lexeme, u.Operand.String())
}

// Grouping represents a grouped expression like (x + y).
type Grouping struct {
	Value Expr
}

func (g Grouping) String() string {
	return fmt.Sprintf("Grouping[%s]", g.Value.String())
}
