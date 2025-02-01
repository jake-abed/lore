package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jake-abed/lore/commands"
	"github.com/jake-abed/lore/internal/config"
	"github.com/jake-abed/lore/internal/db"
	"github.com/jake-abed/lore/internal/utils"
	_ "modernc.org/sqlite"
)

func main() {
	// Only pass in args after cli app name.
	args := utils.SanitizeArgs(os.Args[1:])
	cfg, err := config.ReadConfig()
	if err != nil {
		err = config.CreateDefaultConfig()
		if err != nil {
			msg := commands.ErrorMsg.Render(err.Error())
			fmt.Println(msg)
		}
		cfg, err = config.ReadConfig()
		if err != nil {
			msg := commands.ErrorMsg.Render(err.Error())
			fmt.Println(msg)
		}
	}

	sqliteDb, err := db.OpenDb(&cfg)
	if err != nil {
		msg := commands.ErrorMsg.Render(err.Error())
		fmt.Println(msg)
	}
	defer sqliteDb.Close()

	queries := db.New(sqliteDb)

	state := &commands.State{
		Args: args,
		Cfg:  &cfg,
		Db:   queries,
	}

	allCommands := commands.BuildCommands()
	if len(args) == 0 {
		allCommands["help"].Callback(state)
	} else {
		command, ok := allCommands[args[0]]
		if args[0] == "--help" {
			command = allCommands["help"]
			err := command.Callback(state)
			if err != nil {
				msg := commands.ErrorMsg.Render(err.Error())
				fmt.Println(msg)
			}
		} else if !ok {
			msg := fmt.Sprintf("Lore has no %s command!", args[0])
			fmt.Println(commands.ErrorMsg.Render(msg))
			time.Sleep(1000 * time.Millisecond)
			fmt.Print("Now showing the help command")
			time.Sleep(690 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(690 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(690 * time.Millisecond)
			fmt.Print(".\n")
			time.Sleep(690 * time.Millisecond)
			allCommands["help"].Callback(state)
		} else {
			err := command.Callback(state)
			if err != nil {
				msg := commands.ErrorMsg.Render(err.Error())
				fmt.Println(msg)
			}
		}
	}
}
