package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	b := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	io.Copy(os.Stdout, b)
}
