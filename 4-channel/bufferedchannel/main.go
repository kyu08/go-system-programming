package main

import "fmt"

func main() {
	tasks := make(chan int, 2)

	tasks <- 1
	tasks <- 2

	first := <-tasks
	second := <-tasks

	fmt.Printf("first: %v\n", first)
	fmt.Printf("second: %v\n", second)
}
