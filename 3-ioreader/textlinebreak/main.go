package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目 a`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanLines) // `scanner.Split`で区切りを変更できる
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

	// or...

	// reader := bufio.NewReader(strings.NewReader(source))
	//
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	fmt.Printf("%#v\n", line)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
}
