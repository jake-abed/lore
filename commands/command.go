package commands

import (
	"github.com/jake-abed/auxquest/internal/config"
	"github.com/jake-abed/auxquest/internal/db"
)

type Command struct {
	Flags       map[string]string
	Name        string
	Description string
	Callback    func(*State) error
}

type State struct {
	Cfg  *config.Config
	Db   *db.Queries
	Args []string
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
				"-i": "Looks up info about a particular monster by name or id slug.",
				"-f": "Simulate a fight between two monsters. Name or id slug work.",
			},
			Callback: commandMonsters,
		},
		"npcs": {
			Name:        "npcs",
			Description: "Add, search, edit, and view info about NPCs.",
			Flags: map[string]string{
				"-i": "Inspect an NPC and view their info.",
				"-a": "Add an NPC to your local database for your campaign.",
				"-e": "Edit an NPC's info. [Not implemented yet.]",
				"-s": "Search your NPCs by name. Returns all possible matches.",
			},
			Callback: commandNpcs,
		},
	}
}
