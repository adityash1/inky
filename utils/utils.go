package utils

import (
	"fmt"
	"inky/ast"
	"os"
	"strings"
)

func PrettyPrint(node ast.Node) string {
	lines := []string{}
	buildTreeLines(node, "", "", &lines)
	return strings.Join(lines, "\n")
}

func buildTreeLines(node ast.Node, prefix string, childrenPrefix string, lines *[]string) {
	var nodeDesc string
	var children []ast.Node

	switch n := node.(type) {
	case *ast.Integer:
		nodeDesc = fmt.Sprintf("● Integer: %d", n.Value)
	case *ast.Float:
		nodeDesc = fmt.Sprintf("● Float: %f", n.Value)
	case *ast.BinOp:
		nodeDesc = fmt.Sprintf("● BinOp: %q", n.Op.Lexeme)
		children = []ast.Node{n.Left, n.Right}
	case *ast.UnOp:
		nodeDesc = fmt.Sprintf("● UnOp: %q", n.Op.Lexeme)
		children = []ast.Node{n.Operand}
	case *ast.Grouping:
		nodeDesc = "● Grouping"
		children = []ast.Node{n.Value}
	case *ast.String:
		nodeDesc = fmt.Sprintf("● String: %s", n.Value)
	case *ast.Bool:
		nodeDesc = fmt.Sprintf("● Bool: %t", n.Value)
	case *ast.LogicalOp:
		nodeDesc = fmt.Sprintf("● LogicalOp: %q", n.Op.Lexeme)
		children = []ast.Node{n.Left, n.Right}
	case *ast.PrintStmt:
		nodeDesc = fmt.Sprintf("● PrintStmt: %q", n.End)
		children = []ast.Node{n.Value}
	case *ast.Stmts:
		nodeDesc = "● Stmts"
		children = []ast.Node{}
		for _, stmt := range n.Stmts {
			children = append(children, stmt)
		}
	default:
		nodeDesc = fmt.Sprintf("● Unknown: %T", n)
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

func RuntimeError(msg string, lineno int) {
	fmt.Printf("%v [Line %d]: %s %s", RED, lineno, msg, WHITE)
	os.Exit(1)
}
