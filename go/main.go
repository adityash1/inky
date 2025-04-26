package main

import (
	"blue/lexer"
	"fmt"
	"io"
	"os"
)

func die(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		die("Usage: go run main.go -- <filename>")
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
	fmt.Printf("LEXER: \n")

	tokens := lexer.NewLexer(source).Tokenize()

	for _, v := range tokens {
		fmt.Printf("%+v\n", v)
	}
}
