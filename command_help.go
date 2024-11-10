package main

import (
	"fmt"
	"strings"
)

func commandHelp(state *State) error {
	commands := buildCommands()
	intro := "Welcome to AuxQuest!\n"
	introMsg := "The following commands are available to you: "
	fmt.Println(header.Render(intro + introMsg))
	commandDiv := ""
	for _, command := range commands {
		name := bold.Render(command.name)
		desc := italic.Render(command.description)
		commandDiv += fmt.Sprintf("%-8s <==> %s\n", name, desc) 
		if command.flags != nil {
			for k, v := range command.flags {
				commandDiv += fmt.Sprintf("*** %-6s - %s\n", k, v)
			}
		}
	}
	fmt.Println(commandBox.Render(strings.TrimRight(commandDiv, "\n")))
	return nil
}
