package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type config struct {
	Port  uint `json:"port"`
	Duration int `json:"duration"`
	Mainnet []string `json:"mainnet"`
	Testnet []string `json:"testnet"`
}

var cfg config

func GetConfig() config {
	return cfg
}

func init() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		log.Println(err)
		return
	}
	b, _ := json.Marshal(cfg)
	fmt.Println(string(b))
}
