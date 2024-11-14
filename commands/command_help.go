package commands 

import (
	"fmt"
	"strings"
)

func commandHelp(state *State) error {
	commands := BuildCommands()
	intro := "Welcome to AuxQuest!\n"
	introMsg := "The following commands are available to you: "
	fmt.Println(header.Render(intro + introMsg))
	commandDiv := ""
	for _, command := range commands {
		name := bold.Render(command.Name)
		desc := italic.Render(command.Description)
		commandDiv += fmt.Sprintf("%-12v <==> %v\n", name, desc)
		if command.Flags != nil {
			for k, v := range command.Flags {
				commandDiv += fmt.Sprintf("  *** %-6s - %s\n", k, v)
			}
		}
	}
	fmt.Println(commandBox.Render(strings.TrimRight(commandDiv, "\n")))
	return nil
}
