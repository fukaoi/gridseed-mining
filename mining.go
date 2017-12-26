package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

const (
	CGMINER = "cgminer --gridseed-options=baud=115200,freq=800,chips=40,modules=1,usefifo=0,btc=0 --scrypt "
	POOL = "--url=stratum+tcp://us2.litecoinpool.org:3333 --userpass=yarichin.1:1 --url=stratum+tcp://us.litecoinpool.org:3333 --userpass=yarichin.1:1 "
)

const (
	DELIMITER = ","
)

func main() {
	count := getUsbDeviceCount()
	devices := getUsbDevice(count)
	cmdStr := createMiningCmd(devices)
	_, err := exec.Command(cmdStr).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getUsbDevice(count int) (result []string) {
	comdstr := "lsusb | grep \"STM\" | cut -c5-7,15-18"
	out, err := exec.Command("sh", "-c", comdstr).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	re := regexp.MustCompile("[0-9]{3}\\s[0-9]{3}")
	return re.FindAllString(string(out), count)
}

func getUsbDeviceCount() (count int) {
	usbCount := "lsusb | grep STM | wc -l"
	out, err := exec.Command("sh", "-c", usbCount).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, _ := strconv.Atoi(string(out[0]))
	return result
}

func createMiningCmd(devices []string) (cmd string){
	concatStr := "--usb="
	for i := 0; i < len(devices); i++ {
		split := regexp.MustCompile("\\s").ReplaceAllString(devices[i], ":")
		concatStr += split
		if len(devices) - 1 > i {
			concatStr += DELIMITER
		}
	}

	return CGMINER + POOL + concatStr
}