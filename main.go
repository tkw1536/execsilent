package main

import (
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = nil

	err := cmd.Run()

	// exit code
	if e, ok := err.(*exec.ExitError); ok {
		os.Exit(e.ExitCode())
	}

	// other error
	if err != nil {
		panic(err)
	}

	// ok!
	os.Exit(0)
}
