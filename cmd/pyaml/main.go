package main

import (
	"fmt"
	"github.com/ghabxph/gabriel/blog-server/internal/app/pyaml"
	"os"
)

const usage string = `
Usage: pyaml file.yaml

Executes yaml file. Guide will be created soon.
`

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Errorf("Please pass the yaml file to execute.\n\n")
		fmt.Println(usage)
		return
	}

	pyaml.Execute(args[0])
}
