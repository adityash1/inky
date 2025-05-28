package main

import (
	"fmt"
	"inky/interpreter"
	"inky/lexer"
	"inky/parser"
	"inky/repl"
	"inky/utils"
	"io"
	"os"
)

func die(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	// No arguments, launch REPL
	if len(os.Args) == 1 {
		repl := repl.NewREPL()
		repl.Run()
		return
	}

	// Check for correct arguments for file mode
	if len(os.Args) != 3 || os.Args[1] != "--" {
		fmt.Println("Usage:")
		fmt.Println("  inky                # Start REPL mode")
		fmt.Println("  inky -- <filename>  # Execute file")
		os.Exit(1)
	}

	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		die("Failed to open file: " + err.Error())
	}
	defer file.Close()

	source, err := io.ReadAll(file)
	if err != nil {
		die("Failed to read file: " + err.Error())
	}

	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")
	utils.ColorPrint(utils.GREEN, "Source:")
	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")
	fmt.Print(string(source))

	utils.ColorPrint(utils.GREEN, "\n\n---------------------------\n")
	utils.ColorPrint(utils.GREEN, "Lexer:")
	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")

	tokens := lexer.NewLexer(source).Tokenize()

	for _, v := range tokens {
		fmt.Printf("%v\n", v)
	}

	ast := parser.NewParser(tokens).Parse()

	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")
	utils.ColorPrint(utils.GREEN, "AST:")
	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")
	fmt.Printf("Original AST: \n%v\n\n", ast)
	fmt.Printf("Pretty AST: \n%s\n", utils.PrettyPrint(ast))

	interpreter := interpreter.NewInterpreter()
	typ, result, err := interpreter.Interpret(ast)
	if err != nil {
		die("Interpreter Error: " + err.Error())
	}

	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")
	utils.ColorPrint(utils.GREEN, "Interpreter:")
	utils.ColorPrint(utils.GREEN, "\n---------------------------\n")
	fmt.Printf("%v: %v\n\n", typ, result)
}
