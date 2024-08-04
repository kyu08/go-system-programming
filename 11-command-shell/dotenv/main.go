package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

// go run main.go --e .myenv ./my-printenv のように呼び出す
func main() {
	filename := flag.String("e", ".env", ".env file name to read")
	flag.Parse()
	cmdName := flag.Arg(0)
	args := flag.Args()[1:]
	flag.Args()

	cmd := exec.Command(cmdName, args...)

	envs := make([]string, 0, 10)
	dotenvs, _ := godotenv.Read(*filename)
	for key, value := range dotenvs {
		envs = append(envs, key+"="+value)
	}

	cmd.Env = envs
	o, err := cmd.CombinedOutput()
	fmt.Print(string(o))
	if err != nil {
		panic(err)
	}
}
