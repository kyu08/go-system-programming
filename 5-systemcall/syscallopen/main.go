package main

import "os"

func main() {
	file, _ := os.Create("test.txt")
	defer file.Close()
}
