package utils

import (
	"fmt"
	"inky/ast"
	"os"
	"strings"
)

// -------- v1 --------- //
// func PrettyPrint(expr ast.Expr) string {
// 	return prettyPrintWithIndent(expr, 0)
// }

// func prettyPrintWithIndent(expr ast.Expr, indent int) string {
// 	indentStr := strings.Repeat("  ", indent)

// 	switch node := expr.(type) {
// 	case *Integer:
// 		return fmt.Sprintf("%sInteger: %d", indentStr, node.Value)

// 	case *Float:
// 		return fmt.Sprintf("%sFloat: %f", indentStr, node.Value)

// 	case *BinOp:
// 		var sb strings.Builder
// 		sb.WriteString(fmt.Sprintf("%sBinaryOp: %q\n", indentStr, node.Op.Lexeme))
// 		sb.WriteString(fmt.Sprintf("%s  Left:\n", indentStr))
// 		sb.WriteString(prettyPrintWithIndent(node.Left, indent+2))
// 		sb.WriteString(fmt.Sprintf("\n%s  Right:\n", indentStr))
// 		sb.WriteString(prettyPrintWithIndent(node.Right, indent+2))
// 		return sb.String()

// 	case *UnOp:
// 		var sb strings.Builder
// 		sb.WriteString(fmt.Sprintf("%sUnaryOp: %q\n", indentStr, node.Op.Lexeme))
// 		sb.WriteString(fmt.Sprintf("%s  Operand:\n", indentStr))
// 		sb.WriteString(prettyPrintWithIndent(node.Operand, indent+2))
// 		return sb.String()

// 	case *Grouping:
// 		var sb strings.Builder
// 		sb.WriteString(fmt.Sprintf("%sGrouping:\n", indentStr))
// 		sb.WriteString(prettyPrintWithIndent(node.Value, indent+1))
// 		return sb.String()

// 	default:
// 		return fmt.Sprintf("%sUnknown node type", indentStr)
// 	}
// }

// -------- v2 --------- //
// func PrettyPrint(expr ast.Expr) string {
// 	return formatParenStyle(expr, 0)
// }

// func formatParenStyle(expr ast.Expr, indent int) string {
// 	indentStr := strings.Repeat("  ", indent)
// 	nextIndentStr := strings.Repeat("  ", indent+1)

// 	switch node := expr.(type) {
// 	case *ast.Integer:
// 		return fmt.Sprintf("%sInteger(%d)", indentStr, node.Value)
// 	case *ast.Float:
// 		return fmt.Sprintf("%sFloat(%f)", indentStr, node.Value)
// 	case *ast.BinOp:
// 		var sb strings.Builder
// 		sb.WriteString(fmt.Sprintf("%sBinOp(\n", indentStr))
// 		sb.WriteString(fmt.Sprintf("%s'%s',\n", nextIndentStr, node.Op.Lexeme))
// 		sb.WriteString(formatParenStyle(node.Left, indent+1) + ",\n")
// 		sb.WriteString(formatParenStyle(node.Right, indent+1) + "\n")
// 		sb.WriteString(fmt.Sprintf("%s)", indentStr))
// 		return sb.String()
// 	case *ast.UnOp:
// 		var sb strings.Builder
// 		sb.WriteString(fmt.Sprintf("%sUnOp(\n", indentStr))
// 		sb.WriteString(fmt.Sprintf("%s'%s',\n", nextIndentStr, node.Op.Lexeme))
// 		sb.WriteString(formatParenStyle(node.Operand, indent+1) + "\n")
// 		sb.WriteString(fmt.Sprintf("%s)", indentStr))
// 		return sb.String()
// 	case *ast.Grouping:
// 		var sb strings.Builder
// 		sb.WriteString(fmt.Sprintf("%sGrouping(\n", indentStr))
// 		sb.WriteString(formatParenStyle(node.Value, indent+1) + "\n")
// 		sb.WriteString(fmt.Sprintf("%s)", indentStr))
// 		return sb.String()
// 	default:
// 		return fmt.Sprintf("%sUnknown", indentStr)
// 	}
// }

// -------- v3 --------- //
func PrettyPrint(expr ast.Expr) string {
	lines := []string{}
	buildTreeLines(expr, "", "", &lines)
	return strings.Join(lines, "\n")
}

func buildTreeLines(expr ast.Expr, prefix string, childrenPrefix string, lines *[]string) {
	var nodeDesc string
	var children []ast.Expr

	switch node := expr.(type) {
	case *ast.Integer:
		nodeDesc = fmt.Sprintf("● Integer: %d", node.Value)
	case *ast.Float:
		nodeDesc = fmt.Sprintf("● Float: %f", node.Value)
	case *ast.BinOp:
		nodeDesc = fmt.Sprintf("● BinOp: %q", node.Op.Lexeme)
		children = []ast.Expr{node.Left, node.Right}
	case *ast.UnOp:
		nodeDesc = fmt.Sprintf("● UnOp: %q", node.Op.Lexeme)
		children = []ast.Expr{node.Operand}
	case *ast.Grouping:
		nodeDesc = "● Grouping"
		children = []ast.Expr{node.Value}
	default:
		nodeDesc = "● Unknown"
	}

	// Add the current node to the output lines
	*lines = append(*lines, prefix+nodeDesc)

	// Add all children except the last one
	for i := 0; i < len(children)-1; i++ {
		// For each child except the last, use "├── " as the connector
		// and add "│   " to the prefix for its children
		buildTreeLines(children[i], childrenPrefix+"├── ", childrenPrefix+"│   ", lines)
	}

	// Add the last child
	if len(children) > 0 {
		// For the last child, use "└── " as the connector
		// and add "    " to the prefix for its children
		buildTreeLines(children[len(children)-1], childrenPrefix+"└── ", childrenPrefix+"    ", lines)
	}
}

const (
	WHITE  = "\033[0m"
	BLUE   = "\033[94m"
	CYAN   = "\033[96m"
	GREEN  = "\033[92m"
	YELLOW = "\033[93m"
	RED    = "\033[91m"
)

func ColorPrint(color, msg string) {
	fmt.Printf("%s%s%s", color, msg, WHITE)
}

func ParseError(msg string, lineno int) {
	fmt.Printf("%v [Line %d]: %s %s", RED, lineno, msg, WHITE)
	os.Exit(1)
}

func LexingError(msg string, lineno int) {
	fmt.Printf("%v [Line %d]: %s %s", RED, lineno, msg, WHITE)
	os.Exit(1)
}
