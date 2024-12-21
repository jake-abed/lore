package main

import (
	"fmt"
	"github.com/jake-abed/auxquest/commands"
	"github.com/jake-abed/auxquest/internals/config"
	"github.com/jake-abed/auxquest/internals/db"
	"github.com/jake-abed/auxquest/internals/utils"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	args := utils.SanitizeArgs(os.Args[1:])
	cfg, err := config.ReadConfig()
	if err != nil {
		err = config.CreateDefaultConfig()
		if err != nil {
			fmt.Println(err)
		}
		cfg, err = config.ReadConfig()
		if err != nil {
			fmt.Println(err)
		}
	}

	sqliteDb, err := db.OpenDb(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer sqliteDb.Close()

	state := &commands.State{
		Args: args,
		Cfg:  &cfg,
		Db:   sqliteDb,
	}

	commands := commands.BuildCommands()
	if len(args) == 0 {
		commands["help"].Callback(state)
	} else {
		commands[args[0]].Callback(state)
	}
}
