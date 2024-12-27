package config

import (
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	home, _ := os.UserHomeDir()
	ok, err := checkConfigDir()
	if err != nil {
	}
	if ok {
		os.Remove(home + ".config/auxquest/config.json")
	}

	err = CreateDefaultConfig()
	if err != nil {
		t.Fatalf("An error occurred creating the config file: %!", err)
	}

	cfg, err := ReadConfig()
	if err != nil {
		t.Fatalf("An error occurred reading the config file: %!", err)
	}

	if cfg.DbPath != "/.config/auxquest/sqlite.db" {
		t.Errorf("Expected \"/.config/auxquest/sqlite.db\" as DbPath got=%s",
			cfg.DbPath)
	}

	if cfg.Username != "default" {
		t.Errorf("Expected \"default\" as Username got=%s", cfg.Username)
	}
}
