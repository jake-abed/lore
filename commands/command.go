package commands

import (
	"github.com/jake-abed/lore/internal/config"
	"github.com/jake-abed/lore/internal/db"
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
		"dice": {
			Name:        "dice",
			Description: "Roll dice in the following format: {qtyDice}d{Die}{+/-}{modifier}.",
			Flags: map[string]string{
				"-a":       "Rolls all the dice at once and returns the result.",
				"-i":       "Rolls individual dice, presenting the result, then returning the total.",
				"Examples": "1d6, 2d12+1, 1d4-1, 3d6+8, 5d20+12, 80d100-50, etc.",
			},
			Callback: commandDice,
		},
		"monsters": {
			Name:        "monsters",
			Description: "Get info about monsters and simulate fights.",
			Flags: map[string]string{
				"-i":  "Looks up info about a particular monster by name or id slug.",
				"-f":  "Simulate a fight between two monsters. Name or id slug work.",
				"-va": "View all monsters on the D&D 5e OpenAPI.",
			},
			Callback: commandMonsters,
		},
		"npcs": {
			Name:        "npcs",
			Description: "Add, search, edit, & view info about NPCs.",
			Flags: map[string]string{
				"-v": "Inspect an NPC and view their info.",
				"-a": "Add an NPC to your local database for your campaign.",
				"-e": "Edit an NPC's info by name. Case-insensitive.",
				"-s": "Search your NPCs by name. Returns all possible matches.",
				"-d": "Delete an NPC by name. Case-sensitive.",
			},
			Callback: commandNpcs,
		},
		"places": {
			Name:        "places",
			Description: "Add, search, edit, & view worlds, areas, & locations.",
			Flags: map[string]string{
				"-a":            "Add a place.",
				"-e":            "Edit a place.",
				"-s":            "Search places by name. Returns all possible matches",
				"-d":            "Delete a place by name. Case-sensitive.",
				"-v":            "Inspect a place and it's information by name.",
				"--world":       "Specify an operation on a world.",
				"--area":        "Specify an operation on an area.",
				"--location":    "Specify an operation on a location.",
				"--sublocation": "Specify an operation on a sublocation.",
			},
			Callback: commandPlaces,
		},
	}
}
