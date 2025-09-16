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

	anySuccess := false

	for _, path := range os.Args[1:] {
		file, err := os.Open(path)
		if err != nil {
			printString("ERROR: ")
			printString(err.Error())
			z01.PrintRune('\n')
			continue
		}

		_, copyErr := io.Copy(os.Stdout, file)
		closeErr := file.Close()

		if copyErr != nil {
			printString("ERROR: ")
			printString(copyErr.Error())
			z01.PrintRune('\n')
		} else {
			anySuccess = true
		}

		if closeErr != nil {
			printString("ERROR: ")
			printString(closeErr.Error())
			z01.PrintRune('\n')
		}
	}

	if !anySuccess {
		os.Exit(1)
	}
}
