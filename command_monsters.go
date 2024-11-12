package main

import (
	"fmt"
	"github.com/jake-abed/auxquest/internals/dndapi"
	"time"
)

func commandMonsters(state *State) error {
	if len(state.args) == 1 {
		monstersHelp()
		return nil
	}

	if len(state.args) > 2 && state.args[1] == "-i" {
		inspectMonster(state)
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

func inspectMonster(state *State) {
	argsCount := len(state.args)
	if argsCount < 2 || argsCount > 3 {
		fmt.Println("Incorrect number of args to inspect monster!")
	}

	client := dndapi.NewClient(5 * time.Second)

	monster, err := client.GetMonster(state.args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	
	intro := fmt.Sprintf("%s Info", monster.Name)
	fmt.Println(header.Render(intro))
	
	fmt.Printf(" HP - %d\n", monster.HitPoints)
	fmt.Printf(" Armor Class - %d\n", monster.ArmorClass[0].Value)
	fmt.Printf(" Type - %s\n", monster.Type) 
	fmt.Printf(" Hit Dice - %s\n", monster.HitDice)
	fmt.Printf(" Size - %s\n", monster.Size)
	fmt.Printf(" Alignment - %s\n", monster.Alignment)
	fmt.Printf(" XP Value - %d\n", monster.Xp)
	return
}
