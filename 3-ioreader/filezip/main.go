package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Create("test.zip")
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	w1, _ := zipWriter.Create("a.txt")
	io.Copy(w1, strings.NewReader("hoge1 from strings.Reader"))
	w2, _ := zipWriter.Create("b.txt")
	io.Copy(w2, strings.NewReader("hoge2 from strings.Reader"))
}
