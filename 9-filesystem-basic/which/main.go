package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s [exec path]\n", os.Args[0])
		os.Exit(1)
	}

	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		execpath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(execpath)
		if !os.IsNotExist(err) {
			fmt.Println(execpath)
			return
		}
	}
	os.Exit(1)
}
