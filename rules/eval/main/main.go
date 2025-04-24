package main

import (
	"bufio"
	"fmt"
	"github.com/mindiply/mindcarbon/rules/eval"
	"os"
)

func main() {
	// Create a new interpreter instance
	interpreter := eval.NewInterpreter()

	// Print welcome message
	fmt.Println("Welcome to the mindcarbon REPL!")
	fmt.Println("Type your commands below. Type 'exit' to quit.")

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display prompt
		fmt.Print(">> ")

		// Read input
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		// Exit condition
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		// Evaluate the input using the interpreter
		result, err := interpreter.EvaluateDML(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		// Print the result
		fmt.Printf("Result: %v\n", result)
	}

	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %s\n", err)
	}
}
