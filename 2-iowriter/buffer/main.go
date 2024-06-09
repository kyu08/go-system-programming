package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "hi\n") // os.Stdoutもio.Writerなので当然書き込める

	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("first message\n")
	buffer.Flush()
	fmt.Println("👀")
	buffer.WriteString("second message\n")
	buffer.Flush()
}
