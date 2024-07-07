package main

import (
	"fmt"
	"os"
)

func main() {
	// pwd, _ := os.Getwd()
	// dir, err := os.Open(path.Join(pwd, "emptydir"))
	dir, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1) // dir.Readdirは引数に0以下の数値が渡るとすべての内容を返す
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("[Dir] %s\n", fileInfo.Name())
		} else {
			fmt.Printf("[File] %s\n", fileInfo.Name())
		}
	}
}
