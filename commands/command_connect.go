package commands

import (
	"fmt"
)

func commandConnect(s *State) error {
	args := s.Args[1:]

	// Break out if user did not provide enough flags.
	if len(args) < 2 {
		fmt.Println("Places command requires at least two arguments!")
		connectHelp()
		return nil
	}

	return nil
}

func connectHelp() {
	intro := "Lore Connect Help\n"
	introTip := "Connect subcommand information"
	fmt.Println(header.Render(intro + introTip))
}
