package commands

import (
	"context"
	"fmt"
	"github.com/jake-abed/auxquest/internal/db"
)

func commandNpcs(s *State) error {
	npcArgs := s.Args[1:]
	if len(npcArgs) == 0 {
		fmt.Println("Npcs command expects at least one argument!")
	}

	testNpc := &db.NpcParams{
		Name:        "Tony Da Deer",
		Race:        "Deer",
		Class:       "Friend",
		Subclass:    "",
		Alignment:   "Chaotic Good",
		Level:       69,
		Hitpoints:   420,
		Sex:         "Yes",
		Description: "A godly and lovely Deer with a heart of gold.",
		Languages:   "Deer, English, & Tagalog",
	}

	npc, err := s.Db.AddNpc(context.Background(), testNpc)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(npc)

	return nil
}
