package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	ar := io.NewSectionReader(programming, 5, 1)
	sr := io.NewSectionReader(system, 0, 1)
	cr := io.NewSectionReader(computer, 0, 1)
	ir := io.NewSectionReader(programming, 8, 1)
	ir2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(ar, sr, cr, ir, ir2)
	io.Copy(os.Stdout, stream)
}
