package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
)

func main() {
	p, _ := process.NewProcess(int32(os.Getppid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	fmt.Printf("p.Pid: %v\n", p.Pid)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("cmd: %v\n", cmd)
	fmt.Println(p.Exe())
	fmt.Println(p.Times())
	fmt.Println(p.Username())
}
