package main

type Command struct {
	flags       map[string]string
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

func (c *Command) Execute(cfg *Config) {
}
