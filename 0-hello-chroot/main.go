package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Chroot directory
	newRoot := "/tmp/chroot"

	// Script bash to execute
	scriptPath := "/hello.sh"

	// Change root directory
	if err := syscall.Chroot(newRoot); err != nil {
		fmt.Fprintf(os.Stderr, "chroot failed: %v\n", err)
		os.Exit(1)
	}

	// Change current directory after chroot
	if err := os.Chdir("/"); err != nil {
		fmt.Fprintf(os.Stderr, "chdir failed: %v\n", err)
		os.Exit(1)
	}

	// Executing bash script
	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "script execution failed: %v\n", err)
		os.Exit(1)
	}
}
