package config

import (
	"encoding/json"
	"errors"
	"os"
)

func checkConfigDir() (bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}
	path := home + "/.config/auxquest"
	_, err = os.ReadDir(path)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func CreateDefaultConfig() error {
	found, err := checkConfigDir()
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if !found {
		err = os.Mkdir(home+"/.config/auxquest", 0766)
		if err != nil {
			return err
		}
	}
	cfg := &Config{
		Username: "default",
		DbPath:   "/.config/auxquest/sqlite.db",
	}
	buffer, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(home+"/.config/auxquest/config.json", buffer, 0766)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfig() (Config, error) {
	found, err := checkConfigDir()
	if err != nil {
		return Config{}, err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	if !found {
		return Config{}, errors.New("No config file found!")
	}
	buffer, err := os.ReadFile(home + "/.config/auxquest/config.json")
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	err = json.Unmarshal(buffer, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

type Config struct {
	Username string `json:"username"`
	DbPath   string `json:"dbPath"`
}
