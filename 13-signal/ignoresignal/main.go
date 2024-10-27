package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Accept Ctrl + C for 2 sec")
	time.Sleep(2 * time.Second)

	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	fmt.Println("Ignore Ctrl + C for 10 sec")
	time.Sleep(10 * time.Second)
}
