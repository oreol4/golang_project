package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[2:]
	var SlicaArgs []string
	command := os.Args[1]
	for i := 0; i < len(args); i++ {
		SlicaArgs = append(SlicaArgs, args[i])
	}
	cmd := exec.Command(command, SlicaArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(output))
}
