package main

import (
	"errors"
	"fmt"
	"github.com/jake-abed/auxquest/internals/dndapi"
	"time"
)

func commandMonsters(state *State) error {
	if len(state.args) == 1 {
		monstersHelp()
		return nil
	}

	client := dndapi.NewClient(5 * time.Second)

	monsters, err := client.GetAllMonsters()
	if err != nil {
		return err
	}

	fmt.Println(bold.Render("Showing all monsters!"))
	for _, monster := range monsters {
		fmt.Printf("Monster: %s - Index: %s\n", monster.Name, monster.Index)
	}
	return nil
}

func monstersHelp() {
	intro := "AuxQuest Monsters Help\n"
	introTip := "Monsters subcommands information"
	fmt.Println(header.Render(intro + introTip))
	inspect := bold.Render("  *** monsters -i <monster-name> | ")
	inspectMessage := "Inspect information about specific monster."
	fmt.Println(inspect + inspectMessage)
	fight := bold.Render("  *** monsters -f <monster-1> <monster-2> | ")
	fightMessage := "Simulate a fight between monsters.\n"
	fmt.Println(fight + fightMessage)
}

func inspectMonster(state *State) error {
	argsCount := len(state.args)
	if argsCount < 2 || argsCount > 3 {
		return errors.New("Incorrect number of args to inspect monster!")
	}

	return nil
}
