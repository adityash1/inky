package main

import (
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

	data, err := io.ReadAll(file)
	if err != nil {
		die("Failed to read file: " + err.Error())
	}
	fmt.Print(string(data))

}
