package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	writer.Write([]byte("Hello, World!"))
}
