package ast

import (
	"fmt"
	"inky/token"
)

// Node is the parent interface for all AST nodes
type Node interface {
	String() string
}

// Expr is the interface for expressions
type Expr interface {
	Node
}

// Stat is the interface for statements.
type Stmt interface {
	Node
}

// Integer represents an integer expression.
type Integer struct {
	Value int
	Line  int
}

func (i Integer) String() string {
	return fmt.Sprintf("Integer[%d]", i.Value)
}

// Float represents a float expression.
type Float struct {
	Value float64
	Line  int
}

func (f Float) String() string {
	return fmt.Sprintf("Float[%f]", f.Value)
}

// Bool represents a boolean expression.
type Bool struct {
	Value bool
	Line  int
}

func (b Bool) String() string {
	return fmt.Sprintf("Bool[%t]", b.Value)
}

// String represents a string expression.
type String struct {
	Value string
	Line  int
}

func (s String) String() string {
	return fmt.Sprintf("String[%s]", s.Value)
}

// BinOp represents a binary operation like x + y.
type BinOp struct {
	Op    token.Token
	Left  Expr
	Right Expr
	Line  int
}

func (b BinOp) String() string {
	return fmt.Sprintf("BinOp(%q, %s, %s)", b.Op.Lexeme, b.Left.String(), b.Right.String())
}

// UnOp represents a unary operation like -x.
type UnOp struct {
	Op      token.Token
	Operand Expr
	Line    int
}

func (u UnOp) String() string {
	return fmt.Sprintf("UnOp(%q, %s)", u.Op.Lexeme, u.Operand.String())
}

// Grouping represents a grouped expression like (x + y).
type Grouping struct {
	Value Expr
	Line  int
}

func (g Grouping) String() string {
	return fmt.Sprintf("Grouping(%s)", g.Value.String())
}

// LogicalOp represents a logical operation like x and y.
type LogicalOp struct {
	Op    token.Token
	Left  Expr
	Right Expr
	Line  int
}

func (l LogicalOp) String() string {
	return fmt.Sprintf("Logical(%q, %s, %s)", l.Op.Lexeme, l.Left.String(), l.Right.String())
}

// Stmts represents a list of statements.
type Stmts struct {
	Stmts []Stmt
	Line  int
}

func (s Stmts) String() string {
	return fmt.Sprintf("Stmts(%v)", s.Stmts)
}

type PrintStmt struct {
	Value Expr
	End   string
	Line  int
}

func (p PrintStmt) String() string {
	return fmt.Sprintf("PrintStmt(%s, end=%q)", p.Value.String(), p.End)
}

type IfStmt struct {
	Condition Expr
	ThenStmts *Stmts
	ElseStmts *Stmts
	Line      int
}

func (i IfStmt) String() string {
	var elseStr string
	if i.ElseStmts != nil {
		elseStr = i.ElseStmts.String()
	} else {
		elseStr = "nil"
	}
	return fmt.Sprintf("IfStmt(%s, then:%s, else:%s)", i.Condition.String(), i.ThenStmts.String(), elseStr)
}

type Identifier struct {
	Name string
	Line int
}

func (i Identifier) String() string {
	return fmt.Sprintf("Identifier(%q)", i.Name)
}

type AssignStmt struct {
	Left  Expr
	Right Expr
	Line  int
}

func (a AssignStmt) String() string {
	return fmt.Sprintf("AssignStmt(%s, %s)", a.Left.String(), a.Right.String())
}
