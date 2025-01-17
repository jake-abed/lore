package commands

import (
	"fmt"
	"strings"
)

func commandHelp(s *State) error {
	commands := BuildCommands()

	if len(s.Args) > 1 {
		cmd := s.Args[1]
		err := SingleCommandHelp(s, commands, cmd)
		if err != nil {
			return err
		}
		return nil
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

func SingleCommandHelp(
	s *State,
	commands map[string]Command,
	cmdName string,
) error {
	intro := "Welcome to Lore!\n"
	introMsg := fmt.Sprintf("The %s command has the following flags: ", cmdName)
	fmt.Println(header.Render(intro + introMsg))
	commandDiv := ""
	cmd, ok := commands[cmdName]
	if !ok {
		return fmt.Errorf("That command does not exist in Lore!")
	}
	for key, val := range cmd.Flags {
		name := bold.Render(fmt.Sprintf("%-13s", key))
		desc := italic.Render(val)
		commandDiv += fmt.Sprintf("%s <==> %v\n", name, desc)
	}
	fmt.Println(commandBox.Render(strings.TrimRight(commandDiv, "\n")))
	return nil
}
