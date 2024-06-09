package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("test.txt")
	fmt.Fprintf(file, "%s's age: %d", "taro", 3)
}
