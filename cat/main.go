package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printString("ERROR: ")
			printString(err.Error())
			z01.PrintRune('\n')
			os.Exit(1)
		}
		return
	}

	anyError := false

	for _, path := range os.Args[1:] {
		file, err := os.Open(path)
		if err != nil {
			printString("ERROR: ")
			printString(err.Error())
			z01.PrintRune('\n')
			anyError = true
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			printString("ERROR: ")
			printString(err.Error())
			z01.PrintRune('\n')
			anyError = true
		}

		if err := file.Close(); err != nil {
			printString("ERROR: ")
			printString(err.Error())
			z01.PrintRune('\n')
			anyError = true
		}
	}

	if anyError {
		os.Exit(1)
	}
}
