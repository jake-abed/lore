package commands

import (
	"fmt"
	"github.com/jake-abed/auxquest/internals/dndapi"
	"math/rand/v2"
	"time"
)

func commandMonsters(state *State) error {
	if len(state.Args) == 1 {
		monstersHelp()
		return nil
	}

	if len(state.Args) > 2 && state.Args[1] == "-i" {
		inspectMonster(state)
		return nil
	}

	if len(state.Args) > 3 && state.Args[1] == "-f" {
		monsterFight(state)
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
	argsCount := len(state.Args)
	if argsCount < 2 || argsCount > 3 {
		fmt.Println("Incorrect number of args to inspect monster!")
		return
	}

	client := dndapi.NewClient(5 * time.Second)

	monster, err := client.GetMonster(state.Args[2])
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

func monsterFight(state *State) {
	argsCount := len(state.Args)
	if argsCount < 3 || argsCount < 4 {
		fmt.Println("Incorrect number of args for a monster fight!")
		return
	}

	client := dndapi.NewClient(5 * time.Second)

	ch := make(chan dndapi.Monster, 2)

	for _, monster_name := range state.Args[2:4] {
		go func() {
			res, err := client.GetMonster(monster_name)
			if err != nil {
				fmt.Println(err)
			}
			ch <- res
		}()
	}

	monster1 := <-ch
	monster2 := <-ch

	close(ch)

	simulateFight(monster1, monster2)
}

func simulateFight(monsterOne, monsterTwo dndapi.Monster) {
	fmt.Printf("%s will now fight %s!\n", monsterOne.Name, monsterTwo.Name)
	hpOne := monsterOne.HitPoints
	hpTwo := monsterTwo.HitPoints

	initOne := rand.IntN(19) + (monsterOne.Dexterity-10)/2
	initTwo := rand.IntN(19) + (monsterTwo.Dexterity-10)/2

	var first *dndapi.Monster
	var second *dndapi.Monster

	if initOne >= initTwo {
		first = &monsterOne
		second = &monsterTwo
	} else {
		first = &monsterTwo
		second = &monsterOne
	}

	firstAttacks := first.ParseAttacks()
	secondAttacks := second.ParseAttacks()

	for hpOne > 0 && hpTwo > 0 {
		fmt.Println(&firstAttacks)
		fmt.Println(&secondAttacks)
		hpOne = 0
	}
}	
