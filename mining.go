package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	l "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

const (
	MINER     = "cgminer"
	DELIMITER = ","
	LOGFILE = "./output.log"
)

var (
	log = l.New()
)

func main() {
	log, err := logSetting()
	if err != nil {
		l.Error(err)
	}
	count := getUsbDeviceCount()
	devices := getUsbDevice(count)
	cmdStr := createMiningCmd(devices)
	out, err := exec.Command("sh", "-c", cmdStr).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)
	l.Info(out)

	defer log.Close()
}

func logSetting() (*os.File, error) {
	log.Formatter = new(l.TextFormatter)
	log.Level = l.InfoLevel
	if err := os.Remove(LOGFILE); err != nil {
		l.Error(err)
	}

	logfile, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Info(err)
	}

	return logfile, err
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
