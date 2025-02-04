package commands

import (
	"fmt"
	"strings"
)

func commandHelp(s *State) error {
	commands := BuildCommands()

	if len(s.Args) > 1 {
		cmd := strings.ToLower(s.Args[1])
		switch cmd {
		case "monsters":
			monstersHelp()
			return nil
		case "npcs":
			npcHelp()
			return nil
		case "quests":
			questsHelp()
			return nil
		case "places":
			placesHelp()
			return nil
		case "dice":
			diceHelp()
		case "help":
			helpHelp()
			return nil
		}
	}

	intro := "Welcome to Lore!\n"
	introMsg := "The following commands are available to you: "
	fmt.Println(header.Render(intro + introMsg))
	commandDiv := ""
	for _, command := range commands {
		name := bold.Render(fmt.Sprintf("%-8s", command.Name))
		desc := italic.Render(command.Description)
		commandDiv += fmt.Sprintf("%s <==> %v\n", name, desc)
	}
	fmt.Println(commandBox.Render(strings.TrimRight(commandDiv, "\n")))
	return nil
}

func helpHelp() {
	intro := "Welcome to Lore!\n"
	introMsg := fmt.Sprintf("The %s command has the following flags: ", "help")
	fmt.Println(header.Render(intro + introMsg))
	commandDiv := "The help subcommand helps you. It's helping right now!"
	fmt.Println(commandBox.Render(strings.TrimRight(commandDiv, "\n")))
}
