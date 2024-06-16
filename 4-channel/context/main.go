package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("goroutine started!")
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("end")
}
