package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync/atomic"
)

type Config struct {
	Router struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"router"`
	Name string `json:"name"`
	Type string `json:"type"`
	Extension string `json:"extension"`
}

var (
	config atomic.Value
)

func ParseConfg(configPath string) error{
	newConfig := Config{}

	configuration, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.New(fmt.Sprint("Couldn't load configuration file: ", err.Error()))
	}

	marshalErr := json.Unmarshal(configuration, &newConfig)
	if marshalErr != nil {
		return errors.New(fmt.Sprint("Couldn't unmarshal configuration: ", marshalErr))
	}

	config.Store(newConfig)

	return nil
}

func SetConfig(newConfig Config) {
	config.Store(newConfig)
}

func Global() Config{
	return config.Load().(Config)
}