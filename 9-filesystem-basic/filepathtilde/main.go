package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "~/hoge/fuga/main.go"
	cleanPath, err := cleanTilde(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(cleanPath)
}

func cleanTilde(path string) (string, error) {
	if 1 < len(path) && path[:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, path[2:]), nil
	}
	return path, nil
}
