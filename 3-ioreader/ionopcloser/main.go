package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	read := strings.NewReader("Hello, World!")
	readCloser := io.NopCloser(read)
	fmt.Printf("readCloser: %v\n", readCloser)
}
