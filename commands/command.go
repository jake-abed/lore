package commands

import "github.com/jake-abed/auxquest/internals/config"

type Command struct {
	Flags       map[string]string
	Name        string
	Description string
	Callback    func(*State) error
}

type State struct {
	Args []string
	Cfg  *config.Config
}

func BuildCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Get information about all available commands.",
			Flags:       nil,
			Callback:    commandHelp,
		},
		"monsters": {
			Name:        "monsters",
			Description: "Learn about a monster and potentially store the data.",
			Flags: map[string]string{
				"--cr": "Filter by Challenge Rating using an int or float from 0.25 to 20",
			},
			Callback: commandMonsters,
		},
	}
}
