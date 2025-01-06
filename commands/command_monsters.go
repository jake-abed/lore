package commands

import (
	"fmt"
	"github.com/jake-abed/lore/internal/dndapi"
	"math/rand/v2"
	"slices"
	"strings"
	"time"
)

func commandMonsters(state *State) error {
	if len(state.Args) == 1 {
		monstersHelp()
		return nil
	}

	flag := state.Args[1]

	if len(state.Args) > 2 && (flag == "-i" || flag == "--inspect") {
		inspectMonster(state)
		return nil
	}

	if len(state.Args) > 3 && (flag == "-f" || flag == "--fight") {
		monsterFight(state)
		return nil
	}

	if len(state.Args) > 1 && (flag == "-va" || flag == "--viewall") {
		client := dndapi.NewClient(5 * time.Second)

		monsters, err := client.GetAllMonsters()
		if err != nil {
			return err
		}

		fmt.Println(bold.Render("Viewing all monsters!"))
		for _, monster := range monsters {
			fmt.Printf("Monster: %s - Index: %s\n", monster.Name, monster.Index)
		}
		return nil
	}
	monstersHelp()
	return nil
}

func monstersHelp() {
	intro := "Lore Monsters Help\n"
	introTip := "Monsters subcommands information"
	fmt.Println(header.Render(intro + introTip))
	inspect := bold.Render("  *** monsters -i <monster-name> | ")
	inspectMessage := "Inspect information about specific monster."
	fmt.Println(inspect + inspectMessage)
	fight := bold.Render("  *** monsters -f <monster-1> <monster-2> | ")
	fightMessage := "Simulate a fight between monsters."
	fmt.Println(fight + fightMessage)
	view := bold.Render("  *** monsters -v | ")
	viewMessage := "View all monsters. Pipe into grep for a good time!\n"
	fmt.Println(view + viewMessage)
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

	if monster1.Name == "" || monster2.Name == "" {
		return
	}

	simulateFight(monster1, monster2)
}

func simulateFight(monsterOne, monsterTwo dndapi.Monster) {
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

	hpOne := first.HitPoints
	hpTwo := second.HitPoints

	fmt.Printf("%s will now fight %s!\n", first.Name, second.Name)

	firstAttacks := first.ParseAttacks()
	secondAttacks := second.ParseAttacks()

	for hpOne > 0 && hpTwo > 0 {
		firstAttack := dndapi.UseRandomAttack(firstAttacks)
		secondAttack := dndapi.UseRandomAttack(secondAttacks)
		if firstAttack == nil || secondAttack == nil {
			fmt.Println("One of these monsters does not have valid attacks.")
			fmt.Println("This fight cannot be simulated.")
			return
		}
		firstDamage := firstAttack.Damage
		firstRoll := firstAttack.AttackBonus + rand.IntN(19) + 1
		var firstMessage string
		if firstRoll < second.ArmorClass[0].Value {
			fmt.Printf("%s tried to use %s, but it missed!\n",
				first.Name, firstAttack.Name)
		} else {
			firstMessage = parseDamage(
				&firstDamage,
				firstAttack.Name,
				*firstAttack,
				first,
				second,
			)
		}
		hpTwo -= firstDamage
		fmt.Print(firstMessage)
		time.Sleep(time.Millisecond * 800)

		secondDamage := secondAttack.Damage
		secondRoll := secondAttack.AttackBonus + rand.IntN(19) + 1
		var secondMessage string
		if secondRoll < first.ArmorClass[0].Value {
			fmt.Printf("%s tried to use %s, but it missed!\n",
				second.Name, secondAttack.Name)
		} else {
			secondMessage = parseDamage(
				&secondDamage,
				secondAttack.Name,
				*secondAttack,
				second,
				first,
			)
		}
		hpOne -= secondDamage
		fmt.Print(secondMessage)
		time.Sleep(time.Millisecond * 800)

		fmt.Printf("%s HP Remaining: %d, %s HP Remaining: %d\n",
			first.Name, hpOne, second.Name, hpTwo)
		time.Sleep(time.Millisecond * 1000)
	}
}

func parseDamage(
	damageVal *int,
	attackName string,
	damage dndapi.AttackDamage,
	attacker *dndapi.Monster,
	target *dndapi.Monster,
) (damageMessage string) {
	damageMessage = fmt.Sprintf("%s uses %s. ", attacker.Name, attackName)
	if slices.Contains(target.DamageResistances, strings.ToLower(damage.Type)) {
		*damageVal /= 2
		damageMessage += fmt.Sprintf("%s is resistant to %s. ", target.Name,
			damage.Type)
	}
	if slices.Contains(target.DamageVulnerabilities, strings.ToLower(damage.Type)) {
		*damageVal *= 2
		damageMessage += fmt.Sprintf("%s is weak to %s. ", target.Name, damage.Type)
	}
	if slices.Contains(target.DamageImmunities, strings.ToLower(damage.Type)) {
		*damageVal = 0
		damageMessage += fmt.Sprintf("%s is immune to %s. ", target.Name,
			damage.Type)
	}
	damageMessage += "\n"
	return
}
