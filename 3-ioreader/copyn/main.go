package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	b := bytes.NewBuffer(nil)
	copyn(b, strings.NewReader("0123456789"), int64(5))
	fmt.Printf("b.String(): %v\n", b.String())
}

func copyn(dist io.Writer, src io.Reader, n int64) (int64, error) {
	reader := io.LimitReader(src, n)
	return io.Copy(dist, reader)
}
