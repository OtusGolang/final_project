package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Parameter map[string]int `json:"value"`
}

func Parse() (N int, M int, K int) {
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Разбор JSON
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config.Parameter["N"], config.Parameter["M"], config.Parameter["K"]
}
