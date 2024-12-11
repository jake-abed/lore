package db

import (
	"database/sql"
	"github.com/jake-abed/auxquest/internals/config"
	"os"
)

const DEFAULT_PATH = "/.config/auxquest/sqlite.db"

func OpenDb(cfg *config.Config) (*sql.DB, error) {
	var dbPath string
	if cfg.DbPath == "" {
		dbPath = DEFAULT_PATH
	} else {
		dbPath = cfg.DbPath
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", homeDir+dbPath)
	if err != nil {
		_, err := os.Create(homeDir + dbPath)
		if err != nil {
			return nil, err
		}
		db, err = sql.Open("sqlite3", homeDir+dbPath)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
