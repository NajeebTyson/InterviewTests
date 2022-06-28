package config

import (
	"encoding/json"
	"os"
)

// CONFIG is a globale variable where configurations will be load
var CONFIG Config

// Config is the data type to store configurations
type Config struct {
	TeamApi string
	Teams   []string
}

// LoadConfig takes a config file path as a parameter and return configurtations
// if there are any erros in loading configs, then error will be return
// this function also saves the configs in CONFIG variable
func LoadConfig(configFile string) (Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return CONFIG, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&CONFIG)
	if err != nil {
		return CONFIG, err
	}

	return CONFIG, nil
}
