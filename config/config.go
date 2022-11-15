package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/a1exander256/todo/models"
)

func InitConfig(filePath string) (*models.Config, error) {
	configuration, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load configuration file : %v", err)
	}

	var config models.Config
	if err := json.Unmarshal(configuration, &config); err != nil {
		return nil, fmt.Errorf("could't unmarshal configuration file : %v", err)
	}
	return &config, nil
}
