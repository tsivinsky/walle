package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ImagePath string `json:"image-path"`
}

func GetConfigDir() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDir := fmt.Sprintf("%s/walle", userConfigDir)

	return configDir, nil
}

func getConfigPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	configPath := fmt.Sprintf("%s/config.json", configDir)

	return configPath, nil
}

func CreateConfigPathIfNotExist() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(configDir, 0777)
		if err != nil {
			return err
		}
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		f, err := os.Create(configPath)
		if err != nil {
			return err
		}
		defer f.Close()

		f.Write([]byte("{}"))
	}

	return nil
}

func GetConfig() (*Config, error) {
	config := new(Config)

	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Print("1")
		return nil, err
	}

	return config, nil
}

func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
