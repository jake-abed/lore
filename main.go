package main

import (
	"os"
	"fmt"
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
	state := &State{
		args: args,
		cfg: &cfg,
	}
	commands := buildCommands()
	if len(args) == 0 {
		commands["help"].callback(state)
	} else {
		commands[args[0]].callback(state)
	}
}
