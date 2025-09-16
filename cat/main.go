package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		return
	}

	for _, path := range os.Args[1:] {
		f, err := os.Open(path)
		if err != nil {

			fmt.Println("ERROR:", err)
			continue
		}

		_, copyErr := io.Copy(os.Stdout, f)

		closeErr := f.Close()

		if copyErr != nil {
			fmt.Println("ERROR:", copyErr)
		}

		if closeErr != nil {
			fmt.Println("ERROR:", closeErr)
		}
	}
}
