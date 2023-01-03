package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Set the target IP address
	ip := "XXX.XXX.XXX.XXX"

	// Open the SSH connection
	cmd := exec.Command("ssh", "-o", "StrictHostKeyChecking=no", fmt.Sprintf("root@%s", ip))

	// Set the output and error streams
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}
