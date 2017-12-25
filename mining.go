package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

func main() {
	count := getUsbDeviceCount()
	usb := getUsbDevice(count)
	fmt.Println(usb)
	for i := 0; i < count; i++ {
		fmt.Println(usb[i])
	}
}

func getUsbDevice(count int) (result []string) {
	comdstr := "lsusb | grep \"ID\" | cut -c5-7,15-18"
	out, err := exec.Command("sh", "-c", comdstr).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	re := regexp.MustCompile("[0-9]{3}\\s[0-9]{3}")
	return re.FindAllString(string(out), count)
}

func getUsbDeviceCount() (count int) {
	usbCount := "lsusb | grep ID | wc -l"
	out, err := exec.Command("sh", "-c", usbCount).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, _ := strconv.Atoi(string(out[0]))
	return result
}