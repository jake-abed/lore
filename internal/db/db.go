package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/jake-abed/lore/internal/config"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed sql/*.sql
var embedSQL embed.FS

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	Db DBTX
}

func New(db DBTX) *Queries {
	return &Queries{Db: db}
}

const DEFAULT_PATH = "/.config/lorecli/sqlite.db"

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

	goose.SetBaseFS(embedSQL)

	db, err := goose.OpenDBWithDriver("sqlite", homeDir+dbPath)
	if err != nil {
		fmt.Println("Goose had an issue opening the db!")
		fmt.Println(err)
		return nil, err
	}

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	goose.SetLogger(goose.NopLogger())

	if err := goose.Up(db, "sql"); err != nil {
		panic(err)
	}

	return db, nil
}
