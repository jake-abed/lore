package main

import (
	"os"
	"github.com/jake-abed/auxquest/internals/utils"
)

func main() {
	args := utils.SanitizeArgs(os.Args[1:])
	cfg := &Config{
		args: args,
	}
	commands := buildCommands()
	if len(args) == 0 {
		commands["help"].callback(cfg)
	} else {
		commands[args[0]].callback(cfg)
	}
}
