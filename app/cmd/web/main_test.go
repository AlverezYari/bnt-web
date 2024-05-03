package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	// Call the main function
	main()

	cmd := exec.Command("curl", "localhost:4000")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))

}
