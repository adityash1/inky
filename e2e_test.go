package main

import (
	"inky/interpreter"
	"inky/lexer"
	"inky/parser"
	"testing"
)

type TestCase struct {
	name     string
	source   string
	expected interface{}
}

func TestE2E(t *testing.T) {
	tests := []TestCase{
		// Basic arithmetic
		{
			name:     "Simple arithmetic",
			source:   "3 + 4",
			expected: float64(7),
		},
		{
			name:     "Subtraction",
			source:   "10 - 5",
			expected: float64(5),
		},
		{
			name:     "Multiplication",
			source:   "3 * 4",
			expected: float64(12),
		},
		{
			name:     "Division",
			source:   "10 / 2",
			expected: float64(5),
		},
		{
			name:     "Modulo",
			source:   "10 % 3",
			expected: float64(1),
		},
		{
			name:     "Operator precedence",
			source:   "2 + 3 * 4",
			expected: float64(14),
		},
		{
			name:     "Parenthesized expression",
			source:   "(2 + 3) * 4",
			expected: float64(20),
		},

		// Unary operators
		{
			name:     "Negation",
			source:   "~3",
			expected: float64(-3),
		},
		{
			name:     "Unary minus",
			source:   "-5",
			expected: float64(-5),
		},
		{
			name:     "Unary plus",
			source:   "+5",
			expected: float64(5),
		},
		{
			name:     "Double negation",
			source:   "~-5",
			expected: float64(5),
		},

		// Exponentiation
		{
			name:     "Exponentiation",
			source:   "2^3",
			expected: float64(8),
		},
		{
			name:     "Exponentiation with precedence",
			source:   "2^3^2", // Right associative: 2^(3^2) = 2^9 = 512
			expected: float64(512),
		},
		{
			name:     "Exponentiation with parentheses",
			source:   "(2^3)^2", // (2^3)^2 = 8^2 = 64
			expected: float64(64),
		},

		// Comparison operators
		{
			name:     "Greater than (true)",
			source:   "5 > 3",
			expected: true,
		},
		{
			name:     "Greater than (false)",
			source:   "3 > 5",
			expected: false,
		},
		{
			name:     "Less than (true)",
			source:   "3 < 5",
			expected: true,
		},
		{
			name:     "Less than (false)",
			source:   "5 < 3",
			expected: false,
		},
		{
			name:     "Greater than or equal (true, equal)",
			source:   "5 >= 5",
			expected: true,
		},
		{
			name:     "Greater than or equal (true, greater)",
			source:   "5 >= 3",
			expected: true,
		},
		{
			name:     "Greater than or equal (false)",
			source:   "3 >= 5",
			expected: false,
		},
		{
			name:     "Less than or equal (true, equal)",
			source:   "5 <= 5",
			expected: true,
		},
		{
			name:     "Less than or equal (true, less)",
			source:   "3 <= 5",
			expected: true,
		},
		{
			name:     "Less than or equal (false)",
			source:   "5 <= 3",
			expected: false,
		},

		// Equality operators
		{
			name:     "Equal to (true)",
			source:   "5 == 5",
			expected: true,
		},
		{
			name:     "Equal to (false)",
			source:   "5 == 3",
			expected: false,
		},
		{
			name:     "Not equal to (true)",
			source:   "5 ~= 3",
			expected: true,
		},
		{
			name:     "Not equal to (false)",
			source:   "5 ~= 5",
			expected: false,
		},

		// Boolean expressions
		{
			name:     "Boolean expression",
			source:   "2 > 1 and 5 > 1",
			expected: true,
		},
		{
			name:     "Complex expression",
			source:   "~(3 < 1 + -6)",
			expected: true,
		},
		{
			name:     "Logical AND (true)",
			source:   "true and true",
			expected: true,
		},
		{
			name:     "Logical AND (false)",
			source:   "true and false",
			expected: false,
		},
		{
			name:     "Logical OR (true, both true)",
			source:   "true or true",
			expected: true,
		},
		{
			name:     "Logical OR (true, one true)",
			source:   "true or false",
			expected: true,
		},
		{
			name:     "Logical OR (false)",
			source:   "false or false",
			expected: false,
		},
		{
			name:     "Logical AND with short circuit",
			source:   "true and false or true",
			expected: true,
		},
		{
			name:     "Logical OR with short circuit",
			source:   "false or true and false",
			expected: false,
		},
		{
			name:     "Complex boolean expression",
			source:   "(5 > 3 and 2 < 4) or (7 <= 7 and 9 >= 8)",
			expected: true,
		},
		{
			name:     "Nested boolean with mixed operators",
			source:   "~(false or true and false) and 3^2 > 8",
			expected: true,
		},

		// String literals
		{
			name:     "String comparison",
			source:   "\"hello\" == \"hello\"",
			expected: true,
		},

		// Mixed type operations
		{
			name:     "Boolean negation",
			source:   "~true",
			expected: false,
		},
		{
			name:     "Complex calculation with mixed types",
			source:   "2^3 + 1 > 8 and ~false",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokens := lexer.NewLexer([]byte(test.source)).Tokenize()
			ast := parser.NewParser(tokens).Parse()
			interpreter := interpreter.NewInterpreter()

			_, result, err := interpreter.Interpret(ast)
			if err != nil {
				t.Fatalf("Interpreter error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
