package main

import (
	"os/exec"
	"fmt"
)

func main() {
	out, err := exec.Command("ls", "-la").Output()
	if err == nil {
		fmt.Print(out)
	}
}

