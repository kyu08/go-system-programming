package main

import (
	"fmt"
	"sync"
	"time"
)

var id int

func generateID(mutext *sync.Mutex) int {
	mutext.Lock()
	defer mutext.Unlock()
	id++
	return id
}

func main() {
	mutex := sync.Mutex{}
	for range 10000 {
		go func() {
			fmt.Println(generateID(&mutex))
		}()
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("id: %v\n", id)
}
