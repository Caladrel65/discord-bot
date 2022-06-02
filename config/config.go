package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// Config holds the read-in values from the config file
var Config *configStruct

// Struct that matches config file data
type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	log.Info("Reading config file...")

	// Read the config file
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.WithError(err).Error("error reading config.json")
		return err
	}

	// Unmarshal the config from json
	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.WithError(err).Error("error unmarshalling the json values")
		return err
	}

	return nil

}
