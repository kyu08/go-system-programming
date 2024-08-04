package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Executable()

	fmt.Printf("path: %v\n", path)
}
