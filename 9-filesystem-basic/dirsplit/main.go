package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	sourceFilePath := path.Join(cwd, "main.go")
	dir, name := filepath.Split(sourceFilePath)
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
	fmt.Printf("strings.Split(sourceFilePath, string(filepath.Separator)): %v\n", strings.Split(sourceFilePath, string(filepath.Separator)))
	fmt.Printf("filepath.Base(gopath): %v\n", filepath.Base(sourceFilePath))
	fmt.Printf("filepath.Dir(gopath): %v\n", filepath.Dir(sourceFilePath))
	fmt.Printf("filepath.Ext(gopath): %v\n", filepath.Ext(sourceFilePath))
}
