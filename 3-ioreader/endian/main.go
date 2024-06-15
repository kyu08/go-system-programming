package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 10000を16進数表記すると0x2710
	data := []byte{0x0, 0x0, 0x27, 0x10} // 0x2710をビッグエンディアンで格納
	var i int32

	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("i: %v\n", i)
}
