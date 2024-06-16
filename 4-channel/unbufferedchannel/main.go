package main

import "fmt"

func main() {
	tasks := make(chan int)
	go func() {
		tasks <- 1 // ここをgoroutineで書かないとL11に到達しないので注意（受信されるまでブロックされるため）
	}()

	first := <-tasks

	fmt.Printf("%v\n", first)
}
