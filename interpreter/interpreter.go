package interpreter

import (
	"fmt"
	"inky/ast"
	"inky/token"
	"inky/utils"
	"math"
)

// Constants for different runtime value types
const (
	TYPE_NUMBER = "TYPE_NUMBER"
	TYPE_STRING = "TYPE_STRING"
	TYPE_BOOL   = "TYPE_BOOL"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Interpret(node ast.Node) (string, any, error) {
	switch node := node.(type) {
	case *ast.BinOp:
		return i.visitBinOp(node)
	case *ast.UnOp:
		return i.visitUnOp(node)
	case *ast.LogicalOp:
		return i.visitLogical(node)
	case *ast.Grouping:
		return i.Interpret(node.Value)
	case *ast.Integer:
		return TYPE_NUMBER, float64(node.Value), nil
	case *ast.Float:
		return TYPE_NUMBER, float64(node.Value), nil
	case *ast.String:
		return TYPE_STRING, string(node.Value), nil
	case *ast.Bool:
		return TYPE_BOOL, node.Value, nil
	case *ast.Stmts:
		for _, stmt := range node.Stmts {
			i.Interpret(stmt)
		}
		return "", 0, nil
	case *ast.PrintStmt:
		_, exprVal, err := i.Interpret(node.Value)
		if err != nil {
			return "", 0, err
		}
		fmt.Print(exprVal, node.End)
		return "", 0, nil
	case *ast.IfStmt:
		condType, condVal, err := i.Interpret(node.Condition)
		if err != nil {
			return "", 0, err
		}
		if condType != TYPE_BOOL {
			return "", 0, fmt.Errorf("expected boolean expression, got %s at line %d", condType, node.Line)
		}
		if condVal.(bool) {
			i.Interpret(node.ThenStmts)
		} else if node.ElseStmts != nil {
			i.Interpret(node.ElseStmts)
		}
		return "", 0, nil
	default:
		return "", 0, fmt.Errorf("unknown expression type %T", node)
	}
}

func (i *Interpreter) visitBinOp(node *ast.BinOp) (string, any, error) {
	leftType, leftVal, err := i.Interpret(node.Left)
	if err != nil {
		return "", 0, err
	}
	rightType, rightVal, err := i.Interpret(node.Right)
	if err != nil {
		return "", 0, err
	}

	switch node.Op.Type {

	case token.TOK_PLUS:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_NUMBER, leftNum + rightNum, nil
		} else if leftType == TYPE_STRING || rightType == TYPE_STRING {
			leftStr := fmt.Sprintf("%v", leftVal)
			rightStr := fmt.Sprintf("%v", rightVal)
			return TYPE_STRING, leftStr + rightStr, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Lexeme, leftType, rightType), node.Op.Line)
		}

	case token.TOK_MINUS:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_NUMBER, leftNum - rightNum, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Lexeme, leftType, rightType), node.Op.Line)
		}

	case token.TOK_STAR:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_NUMBER, leftNum * rightNum, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Lexeme, leftType, rightType), node.Op.Line)
		}

	case token.TOK_SLASH:
		if rightType == TYPE_NUMBER && rightVal.(float64) == 0 {
			utils.RuntimeError("division by zero", node.Op.Line)
		}
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_NUMBER, leftNum / rightNum, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_MOD:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_NUMBER, math.Mod(leftNum, rightNum), nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_CARET:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_NUMBER, math.Pow(leftNum, rightNum), nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_GT:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_BOOL, leftNum > rightNum, nil
		} else if leftType == TYPE_STRING && rightType == TYPE_STRING {
			leftStr := leftVal.(string)
			rightStr := rightVal.(string)
			return TYPE_BOOL, leftStr > rightStr, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_LT:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_BOOL, leftNum < rightNum, nil
		} else if leftType == TYPE_STRING && rightType == TYPE_STRING {
			leftStr := leftVal.(string)
			rightStr := rightVal.(string)
			return TYPE_BOOL, leftStr < rightStr, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_GE:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_BOOL, leftNum >= rightNum, nil
		} else if leftType == TYPE_STRING && rightType == TYPE_STRING {
			leftStr := leftVal.(string)
			rightStr := rightVal.(string)
			return TYPE_BOOL, leftStr >= rightStr, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_LE:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_BOOL, leftNum <= rightNum, nil
		} else if leftType == TYPE_STRING && rightType == TYPE_STRING {
			leftStr := leftVal.(string)
			rightStr := rightVal.(string)
			return TYPE_BOOL, leftStr <= rightStr, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_EQEQ:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_BOOL, leftNum == rightNum, nil
		} else if leftType == TYPE_STRING && rightType == TYPE_STRING {
			leftStr := leftVal.(string)
			rightStr := rightVal.(string)
			return TYPE_BOOL, leftStr == rightStr, nil
		} else if leftType == TYPE_BOOL && rightType == TYPE_BOOL {
			leftBool := leftVal.(bool)
			rightBool := rightVal.(bool)
			return TYPE_BOOL, leftBool == rightBool, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	case token.TOK_NE:
		if leftType == TYPE_NUMBER && rightType == TYPE_NUMBER {
			leftNum := leftVal.(float64)
			rightNum := rightVal.(float64)
			return TYPE_BOOL, leftNum != rightNum, nil
		} else if leftType == TYPE_STRING && rightType == TYPE_STRING {
			leftStr := leftVal.(string)
			rightStr := rightVal.(string)
			return TYPE_BOOL, leftStr != rightStr, nil
		} else if leftType == TYPE_BOOL && rightType == TYPE_BOOL {
			leftBool := leftVal.(bool)
			rightBool := rightVal.(bool)
			return TYPE_BOOL, leftBool != rightBool, nil
		} else {
			utils.RuntimeError(fmt.Sprintf("unsupported operator %v between %v and %v", node.Op.Type, leftType, rightType), node.Op.Line)
		}

	default:
		return "", 0, fmt.Errorf("unsupported binary operator %v", node.Op.Type)
	}
	return "", 0, nil
}

func (i *Interpreter) visitUnOp(node *ast.UnOp) (string, any, error) {
	operandType, operand, err := i.Interpret(node.Operand)
	if err != nil {
		return "", 0, err
	}

	switch node.Op.Type {

	case token.TOK_MINUS:
		if operandType == TYPE_NUMBER {
			operandNum := operand.(float64)
			return operandType, -operandNum, nil
		} else {
			return "", 0, fmt.Errorf("unsupported unary operator %v on type %v", node.Op.Type, operandType)
		}

	case token.TOK_PLUS:
		if operandType == TYPE_NUMBER {
			operandNum := operand.(float64)
			return TYPE_NUMBER, operandNum, nil
		} else {
			return "", 0, fmt.Errorf("unsupported unary operator %v on type %v", node.Op.Type, operandType)
		}

	case token.TOK_NOT:
		if operandType == TYPE_BOOL {
			operandBool := operand.(bool)
			return operandType, !operandBool, nil
		} else if operandType == TYPE_NUMBER {
			operandNum := operand.(float64)
			return TYPE_NUMBER, -operandNum, nil
		} else {
			return "", 0, fmt.Errorf("unsupported unary operator %v on type %v", node.Op.Type, operandType)
		}

	default:
		return "", 0, fmt.Errorf("unsupported unary operator %v", node.Op.Type)
	}
}

func (i *Interpreter) visitLogical(node *ast.LogicalOp) (string, any, error) {
	leftType, leftVal, err := i.Interpret(node.Left)
	if err != nil {
		return "", 0, err
	}

	// short circuit evaluation
	if node.Op.Type == token.TOK_AND {
		if leftType == TYPE_BOOL && leftVal.(bool) {
			return i.Interpret(node.Right)
		} else {
			return TYPE_BOOL, false, nil
		}
	} else if node.Op.Type == token.TOK_OR {
		if leftType == TYPE_BOOL && leftVal.(bool) {
			return TYPE_BOOL, true, nil
		} else {
			return i.Interpret(node.Right)
		}
	}
	return "", 0, nil
}
