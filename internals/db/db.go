package db

import (
	"database/sql"
	"fmt"
	"github.com/jake-abed/auxquest/internals/config"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
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
		fmt.Println("Something went wrong get the user's home path!")
		return nil, err
	}

	_, err = os.ReadFile(homeDir + dbPath)
	if err != nil {
		fmt.Println(err)
		_, err = os.Create(homeDir + dbPath)
		if err != nil {
			fmt.Println(err)
		}
	}

	db, err := goose.OpenDBWithDriver("sqlite", homeDir+dbPath)
	if err != nil {
		fmt.Println("Goose had an issue opening the db!")
		fmt.Println(err)
		return nil, err
	}

	goose.SetLogger(goose.NopLogger())

	err = goose.Up(db, "./sql")
	if err != nil {
		fmt.Println("Goose had an error!")
		fmt.Println(err)
	}

	return db, nil
}
