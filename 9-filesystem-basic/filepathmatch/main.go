package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 利用できるパターンはfilepath.Mathのコードコメントに書かれているので使う時は参照するとよさそう
	fmt.Println(filepath.Match("*sampl[^a]*", "main.sample.php"))

	files, err := filepath.Glob("*le")
	if err != nil {
		panic(err)
	}
	fmt.Printf("files: %v\n", files)
}
