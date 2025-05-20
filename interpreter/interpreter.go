package interpreter

import (
	"fmt"
	"inky/ast"
	"inky/token"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Interpret(expr ast.Expr) (float64, error) {
	switch node := expr.(type) {
	case *ast.BinOp:
		return i.visitBinOp(node)
	case *ast.UnOp:
		return i.visitUnOp(node)
	case *ast.Grouping:
		return i.Interpret(node.Value)
	case *ast.Integer:
		return float64(node.Value), nil
	case *ast.Float:
		return node.Value, nil
	default:
		return 0, fmt.Errorf("unknown expression type %T", node)
	}
}

func (i *Interpreter) visitBinOp(node *ast.BinOp) (float64, error) {
	left, err := i.Interpret(node.Left)
	if err != nil {
		return 0, err
	}
	right, err := i.Interpret(node.Right)
	if err != nil {
		return 0, err
	}

	switch node.Op.Type {
	case token.TOK_PLUS:
		return left + right, nil
	case token.TOK_MINUS:
		return left - right, nil
	case token.TOK_STAR:
		return left * right, nil
	case token.TOK_SLASH:
		return left / right, nil
	default:
		return 0, fmt.Errorf("unsupported binary operator %v", node.Op.Type)
	}
}

func (i *Interpreter) visitUnOp(node *ast.UnOp) (float64, error) {
	operand, err := i.Interpret(node.Operand)
	if err != nil {
		return 0, err
	}

	switch node.Op.Type {
	case token.TOK_MINUS:
		return -operand, nil
	case token.TOK_PLUS:
		return operand, nil
	default:
		return 0, fmt.Errorf("unsupported unary operator %v", node.Op.Type)
	}
}
