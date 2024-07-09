package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var imageSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tiff": true,
	".eps":  true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [path to find]", os.Args[0])
		return
	}

	root := os.Args[1]

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(info.Name()))
		if imageSuffix[ext] {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				panic(err)
			}
			fmt.Println(rel)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
