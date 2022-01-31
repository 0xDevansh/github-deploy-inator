package config

import (
	"encoding/json"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"io/ioutil"
)

func ReadConfig() (structs.Config, error) {
	configText, err := ioutil.ReadFile("config.json")
	if err != nil {
		return structs.Config{}, err
	}

	var config structs.Config
	err = json.Unmarshal(configText, &config)
	if err != nil {
		return structs.Config{}, err
	}

	err = ValidateConfig(&config)
	if err != nil {
		return structs.Config{}, err
	}

	return config, nil
}
