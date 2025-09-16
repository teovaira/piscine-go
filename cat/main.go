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

func exitError(err error) {
	printString("ERROR: ")
	printString(err.Error())
	z01.PrintRune('\n')
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			exitError(err)
		}
		return
	}

	for _, path := range os.Args[1:] {
		f, err := os.Open(path)
		if err != nil {
			exitError(err)
		}

		if _, err := io.Copy(os.Stdout, f); err != nil {
			_ = f.Close()
			exitError(err)
		}

		if err := f.Close(); err != nil {
			exitError(err)
		}
	}
}
