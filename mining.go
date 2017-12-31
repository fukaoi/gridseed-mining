package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

const (
	MINER     = "cgminer"
	DELIMITER = ","
)

func main() {
	count := getUsbDeviceCount()
	devices := getUsbDevice(count)
	cmdStr := createMiningCmd(devices)
	out, err := exec.Command("sh", "-c", cmdStr).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)
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

func createMiningCmd(devices []string) (cmd string) {
	concatStr := "--usb="
	for i := 0; i < len(devices); i++ {
		split := regexp.MustCompile("\\s").ReplaceAllString(devices[i], ":")
		concatStr += split
		if len(devices)-1 > i {
			concatStr += DELIMITER
		}
	}

	return MINER + getMinerInfo() + getPoolInfo() + concatStr
}

func getMinerInfo() (info string) {
	c := configure()
	return MINER + c.Algo + c.Option
}

func getPoolInfo() (info string) {
	c := configure()
	return c.Host1 + c.Host2 + c.UserPassWorker
}

func configure() (pool pool) {
	var tomlConfig tomlConfig

	_, err := toml.DecodeFile("config.toml", &tomlConfig)
	if err != nil {
		fmt.Println("config.toml error: ", err)
		os.Exit(9)
	}
	return tomlConfig.Pool
}
