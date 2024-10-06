package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	inputs := make(chan []byte)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		a, _ := os.ReadFile("main.go")
		inputs <- a
		wg.Done()
	}()

	go func() {
		a, _ := os.ReadFile("go.mod")
		inputs <- a
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(inputs)
	}()

	for a := range inputs {
		fmt.Println(string(a))
	}
}
