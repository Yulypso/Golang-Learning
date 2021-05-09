package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
	}

	io.Copy(os.Stdout, file)
}
