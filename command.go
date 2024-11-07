package main

import "github.com/jake-abed/auxquest/internals/config"

type Command struct {
	flags       map[string]string
	name        string
	description string
	callback    func(*State) error
}

type State struct {
	args []string
	cfg  *config.Config
}

func buildCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Get information about all available commands.",
			flags:       nil,
			callback:    commandHelp,
		},
		"monsters": {
			name:        "monsters",
			description: "Learn about a monster and potentially store the data.",
			flags: map[string]string{
				"--cr": "Filter by Challenge Rating using an int or float from 0.25 to 20",
			},
			callback: commandMonsters,
		},
	}
}
