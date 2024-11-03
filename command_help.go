package main

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	commands := buildCommands()
	intro := "Welcome to AuxQuest!\n"
	introMsg := "The following commands are available to you: "
	fmt.Println(header.Render(intro + introMsg))
	commandDiv := ""
	for _, command := range commands {
		commandDiv += fmt.Sprintf("%-8s <==> %s\n", command.name, command. description)
		if command.flags != nil {
			for k, v := range command.flags {
				commandDiv += fmt.Sprintf("    *** %-6s - %s\n", k, v)
			}
		}
	}
	fmt.Println(commandBox.Render(commandDiv))
	return nil
}
