package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("pid: %v\n", pid)

	ppid := os.Getppid()
	fmt.Printf("ppid: %v\n", ppid)
}
