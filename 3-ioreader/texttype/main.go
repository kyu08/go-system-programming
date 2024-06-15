package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)

	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s) // fmt.Fscanはデータがスペース区切りであることを前提としている
	// fmt.Fscanf(reader, "%v, %v, %v, %v", &i, &f, &g, &s) // fmt.Fscanfを使うと任意のフォーマットのデータをパースできる
	// var source = "123, 1.234, 1.0e4, test"
	fmt.Printf("%#v, %#v, %#v, %#v\n", i, f, g, s)
}
