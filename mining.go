package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

func main() {
	comdstr := "lsusb | grep \"ID\" | cut -c5-7,15-18"
	usbCount := "lsusb | grep ID | wc -l"

	out, err := exec.Command("sh", "-c", comdstr).Output()
	countOut, errOut := exec.Command("sh", "-c", usbCount).Output()

	if err != nil || errOut != nil {
		fmt.Println(err, errOut)
		os.Exit(1)
	}

	count, _ := strconv.Atoi(string(countOut[0]))
	re := regexp.MustCompile("[0-9]{3}\\s[0-9]{3}\\n")
	var result []string = re.FindAllString(string(out), count)
	fmt.Println(result[0])
}