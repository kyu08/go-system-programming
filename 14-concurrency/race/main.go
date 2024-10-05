package main

import (
	"fmt"
	"time"
)

// go run -race ./... を実行するとrace conditionが検知できる
func main() {
	count := 0

	go func() {
		count++
	}()

	go func() {
		if count == 1 {
			count = 100
			return
		}
		count++
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("count: %v\n", count)
}
