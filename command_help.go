package main

import "fmt"

func commandHelp(cfg *Config) error {
	commands := buildCommands()
	fmt.Println("The following commands are available to you: ")
	for _, command := range commands {
		fmt.Printf("%-8s <==> %s\n", command.name, command. description)
		if command.flags != nil {
			for k, v := range command.flags {
				fmt.Printf("    *** %-6s - %s\n", k, v)
			}
		}
	}
	return nil
}
