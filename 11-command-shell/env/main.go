package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.ExpandEnv(\"HOME\"): %v\n", os.ExpandEnv("${HOME}"))
}
