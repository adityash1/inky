package repl

import (
	"bufio"
	"fmt"
	"inky/interpreter"
	"inky/lexer"
	"inky/parser"
	"inky/utils"
	"os"
	"strings"
)

type REPL struct {
	interpreter *interpreter.Interpreter
	reader      *bufio.Reader
	prompt      string
	isRunning   bool
}

func NewREPL() *REPL {
	return &REPL{
		interpreter: interpreter.NewInterpreter(),
		reader:      bufio.NewReader(os.Stdin),
		prompt:      "inky> ",
		isRunning:   false,
	}
}

func (r *REPL) Run() {
	r.isRunning = true
	r.printWelcomeMessage()

	for r.isRunning {
		fmt.Print(r.prompt)
		input, err := r.reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if r.handleSpecialCommands(input) {
			continue
		}

		r.evaluate([]byte(input))
	}
}

func (r *REPL) printWelcomeMessage() {
	fmt.Println()
	utils.ColorPrint(utils.CYAN, "  ╭────── ")
	utils.ColorPrint(utils.GREEN, "✧ Inky Language ")
	utils.ColorPrint(utils.YELLOW, "v0.1 ✧")
	utils.ColorPrint(utils.CYAN, " ──────╮\n")

	utils.ColorPrint(utils.BLUE, "  ▸ ")
	utils.ColorPrint(utils.WHITE, "Type ")
	utils.ColorPrint(utils.GREEN, ".help")
	utils.ColorPrint(utils.WHITE, " for commands, ")
	utils.ColorPrint(utils.RED, ".exit")
	utils.ColorPrint(utils.WHITE, " to quit\n")

	fmt.Println()
}

// handleSpecialCommands handles special REPL commands that start with a dot
func (r *REPL) handleSpecialCommands(input string) bool {
	if !strings.HasPrefix(input, ".") {
		return false
	}

	command := strings.TrimSpace(strings.TrimPrefix(input, "."))
	switch command {
	case "exit", "quit":
		r.isRunning = false
		fmt.Println("Goodbye!")
		return true
	case "help":
		r.printHelp()
		return true
	case "clear":
		// Clear screen - works on most terminals
		fmt.Print("\033[H\033[2J")
		return true
	default:
		fmt.Printf("Unknown command: %s\n", command)
		return true
	}
}

func (r *REPL) printHelp() {
	utils.ColorPrint(utils.YELLOW, "Inky REPL Commands:\n")
	fmt.Printf("  %s.help%s          Show this help message\n", utils.GREEN, utils.WHITE)
	fmt.Printf("  %s.exit%s, %s.quit%s   Exit the REPL\n", utils.GREEN, utils.WHITE, utils.GREEN, utils.WHITE)
	fmt.Printf("  %s.clear%s         Clear the screen\n", utils.GREEN, utils.WHITE)

	utils.ColorPrint(utils.YELLOW, "\nExample expressions:\n")
	fmt.Printf("  %s2 + 3 * 4%s\n", utils.BLUE, utils.WHITE)
	fmt.Printf("  %s\"hello\" + \" \" + \"world\"%s\n", utils.BLUE, utils.WHITE)
	fmt.Printf("  %s2 > 1 and 3 < 4%s\n", utils.BLUE, utils.WHITE)
}

// evaluate processes a single line of input
func (r *REPL) evaluate(source []byte) {
	// We need to wrap errors with recover to prevent the REPL from crashing
	defer func() {
		if err := recover(); err != nil {
			utils.ColorPrint(utils.RED, fmt.Sprintf("Runtime error: %v\n", err))
		}
	}()

	tokens := lexer.NewLexer(source).Tokenize()

	// Skip evaluation if there are no tokens (e.g., only comments)
	if len(tokens) == 0 {
		return
	}

	ast := parser.NewParser(tokens).Parse()
	if ast == nil {
		utils.ColorPrint(utils.RED, "Parse error: Could not parse expression\n")
		return
	}

	typ, result, err := r.interpreter.Interpret(ast)
	if err != nil {
		utils.ColorPrint(utils.RED, fmt.Sprintf("Interpreter error: %v\n", err))
		return
	}

	utils.ColorPrint(utils.WHITE, fmt.Sprintf("%v: %v\n", typ, result))
}
