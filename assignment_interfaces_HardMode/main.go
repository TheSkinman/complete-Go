package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ERROR: File name not provided.\n    USE: go run main.go <fileName>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("ERROR: File [%v] does not exist.\n", fileName)
		os.Exit(1)
	}

	file, _ := os.Open(fileName)
	io.Copy(os.Stdout, file)
}
