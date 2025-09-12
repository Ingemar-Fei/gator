package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = "gator/.gatorconfig.json"

type Config struct {
	DBURL       string `json:"db_url"`
	CurUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path: %v", err)
	}
	file, err := os.Open(fullPath)
	if err != nil {

		return Config{}, fmt.Errorf("error reading congfig : %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding config : %v", err)
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurUserName = userName
	return write(*cfg)
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
