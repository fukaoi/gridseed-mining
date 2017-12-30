package main

import (
	"fmt"
	"vendor/github.com/BurntSushi/toml"
)

type Config struct {
	CGMINEER CgminerConfig
}

type CgminerConfig struct {
	Port     uint
	Version  string
	Endpoint string
	Debug    bool
}

var config Config

_, err := toml.DecodeFile("config.toml", &config)
if err != nil {
	fmt.Println("port:", config.API.port)
}
