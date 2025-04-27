package commands

import (
	"github.com/jake-abed/lore/internal/config"
	"github.com/jake-abed/lore/internal/db"
)

type Command struct {
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
			Callback:    commandHelp,
		},
		"dice": {
			Name:        "dice",
			Description: "Roll dice in the following format: {qty}d{Die}{+/-}{modifier}.",
			Callback:    commandDice,
		},
		"monsters": {
			Name:        "monsters",
			Description: "Get info about monsters and simulate fights.",
			Callback:    commandMonsters,
		},
		"npcs": {
			Name:        "npcs",
			Description: "Add, search, edit, & view info about NPCs.",
			Callback:    commandNpcs,
		},
		"places": {
			Name:        "places",
			Description: "Add, search, edit, & view worlds, areas, & locations.",
			Callback:    commandPlaces,
		},
		"quests": {
			Name:        "quests",
			Description: "Add, search, edit & view quests & quest steps.",
			Callback:    commandQuests,
		},
		"connect": {
			Name:        "connect",
			Description: "Connect quests, npcs, and places together.",
			Callback:    commandConnect,
		},
	}
}
