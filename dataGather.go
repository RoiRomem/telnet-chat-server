package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Port string `json:"port"`
}

func getPort() string {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("unable to open config file: %s", err.Error())
		fmt.Scan()
		return ("err")
	}
	defer file.Close()

	var config Data
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("unable to decode config file: %s", err.Error())
		fmt.Scan()
		return ("err")
	}
	return (config.Port)
}
