package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
