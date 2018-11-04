package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Base baseInfo
}

type baseInfo struct {
	Url string
}

func initConfig(filename string, conf *Config) {
	_, err := toml.DecodeFile(filename, &conf)
	if err != nil {
		log.Fatal(err)
	}
}
