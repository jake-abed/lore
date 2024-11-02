package main

type Command struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	args []string
}

func buildCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Get information about available commands.",
			callback:    commandHelp,
		},
		"monster": {
			name: "monster",
			description: "Learn about a monster and potential store the data.",
			callback: commandMonster,
		},
	}
}

func (c *Command) Execute(cfg *Config) {
}
