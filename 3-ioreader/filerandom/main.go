package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}

	limitReader := io.LimitReader(rand.Reader, 1024)
	io.Copy(file, limitReader)
}
