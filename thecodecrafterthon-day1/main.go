package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("✦ Terminal Calculator ✦")
	fmt.Println("Type 'help' for instructions.")

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Try again.")
			continue
		}

		parts := strings.Fields(strings.TrimSpace(line))
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "quit":
			fmt.Println("Goodbye! ✦")
			return

		case "help":
			printHelp()

		case "add", "sub", "mul", "div":
			if len(parts) != 3 {
				fmt.Printf("Error: '%s' requires exactly 2 arguments.\n", cmd)
				continue
			}

			a, err1 := strconv.Atoi(parts[1])
			b, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error: arguments must be valid integers.")
				continue
			}

			switch cmd {
			case "add":
				fmt.Printf("✦ Result: %d\n", a+b)
			case "sub":
				fmt.Printf("✦ Result: %d\n", a-b)
			case "mul":
				fmt.Printf("✦ Result: %d\n", a*b)
			case "div":
				if b == 0 {
					fmt.Println("Error: division by zero is not allowed.")
				} else {
					fmt.Printf("✦ Result: %d\n", a/b)
				}
			}

		default:
			fmt.Printf("Unknown command: '%s'. Type 'help' for usage.\n", cmd)
		}
	}
}

func printHelp() {
	fmt.Println("Supported commands:")
	fmt.Println("  add <a> <b>   addition")
	fmt.Println("  sub <a> <b>   subtraction")
	fmt.Println("  mul <a> <b>   multiplication")
	fmt.Println("  div <a> <b>   division")
	fmt.Println("  help          show this message")
	fmt.Println("  quit          exit the calculator")
}
