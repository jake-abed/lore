package main

import (
	"os"
	"fmt"
	"github.com/jake-abed/auxquest/commands"
	"github.com/jake-abed/auxquest/internals/utils"
	"github.com/jake-abed/auxquest/internals/config"
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
	state := &commands.State{
		Args: args,
		Cfg: &cfg,
	}
	commands := commands.BuildCommands()
	if len(args) == 0 {
		commands["help"].Callback(state)
	} else {
		commands[args[0]].Callback(state)
	}
}
