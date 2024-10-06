package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}
	smap.Store("key", "value")
	smap.Store(1, 2)
	smap.Delete("test")
	// anyで返ってくるので型アサーションが必要
	v, ok := smap.Load("key")
	fmt.Println(v, ok)
}
