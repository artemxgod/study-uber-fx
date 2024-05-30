package main

import (
	"log"
	"flag"

	"github.com/artemxgod/study/study-uber-fx/configs"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "./configs/config.yaml", "Path to config file")
	flag.Parse()

	config, err := configs.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config.ServiceName)
}