package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	fmt.Println(filepath.Clean("~/path/filepath//path.go"))

	// filepath.Abs(path)はpathが絶対パスではない場合にcwd/pathを返す
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Printf("abspath: %v\n", abspath)

	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/filepath/path.go")
	fmt.Printf("relpath: %v\n", relpath)
}
