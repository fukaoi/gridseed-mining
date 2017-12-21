package main

import (
	"os/exec"
	"fmt"
)

func main() {
	//comdstr := "lsusb | grep \"STM\" | cut -c5-7,15-18"
	comdstr := "lsusb | grep \"ID\" | cut -c5-7,15-18"
	out, err := exec.Command("sh", "-c", comdstr).Output()

	//out2, _ := exec.Command("ls", "-la").Output()
	if err == nil {
		s := fmt.Sprintf("%s", out)
		fmt.Println(s)
	}
}

