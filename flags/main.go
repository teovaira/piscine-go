package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || hasArg(args, "--help") || hasArg(args, "-h") {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	input := ""

	for _, arg := range args {
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else if len(arg) > 0 && arg[0] != '-' {
			input = arg
		}
	}

	result := input + insertStr

	if orderFlag {
		runes := []rune(result)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[j] < runes[i] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		result = string(runes)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func hasArg(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("         This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("         This flag will behave like a boolean, if it is called it will order the argument.")
}
