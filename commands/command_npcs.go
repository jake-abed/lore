package commands

import (
	"fmt"
)

func commandNpcs(s *State) error {
	npcArgs := s.Args[1:]
	if len(npcArgs) == 0 {
		fmt.Println("Npcs command expects at least one argument!")
	}

	return nil
}
