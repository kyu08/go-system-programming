package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	oldFileName := "old.txt"
	oldFile, err := os.Open(oldFileName)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFileName := flag.String("f", "new file name", "string flag")
	flag.Parse()
	if newFileName == nil {
		panic("new file name was not given!")
	}
	newFile, err := os.Create(*newFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, oldFile)
}
