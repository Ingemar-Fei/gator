package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl       string `json:"db_url"`
	CurUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"
const configDirName = "gator"

func (cfg *Config) SetUser(username string) error {
	cfg.CurUserName = username
	return write(*cfg)
}
func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting config filepath: %v", err)
	}
	fullPath := filepath.Join(homePath, configDirName, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return fmt.Errorf("error encoding config : %v", err)
	}
	return nil
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding file: %v", err)
	}
	return cfg, nil
}
