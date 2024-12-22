package commands

import (
	"context"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/jake-abed/auxquest/internal/db"
	"os"
	"strconv"
)

func commandNpcs(s *State) error {
	npcArgs := s.Args[1:]
	if len(npcArgs) < 1 {
		fmt.Println("Npcs command expects at least one argument!")
		os.Exit(0)
	}

	flag := npcArgs[0]

	if len(npcArgs) == 1 && (flag == "-a" || flag == "-add") {
		err := addNpc(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if len(npcArgs) == 2 && (flag == "-v" || flag == "--view") {
		name := npcArgs[1]
		npc, err := s.Db.ViewNpcByName(context.Background(), name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		viewNpc(npc)
	}

	if len(npcArgs) == 2 && (flag == "-e" || flag == "--edit") {
		name := npcArgs[1]
		npc, err := s.Db.ViewNpcByName(context.Background(), name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		editNpc(npc, s)
	}
	return nil
}

func npcHelp() {
	intro := "AuxQuest Npc Help\n"
	introTip := "Monsters subcommands information"
	fmt.Println(header.Render(intro + introTip))
	add := bold.Render("  *** npc -a <npc-name> | ")
	addMessage := "Add a new npc by name."
	fmt.Println(add + addMessage)
	fight := bold.Render("  *** monsters -f <monster-1> <monster-2> | ")
	fightMessage := "Simulate a fight between monsters.\n"
	fmt.Println(fight + fightMessage)

}

func addNpc(s *State) error {
	var name string
	var race string
	var class string
	var subclass string
	var alignment string
	var sex string
	var desc string
	var languages string
	var level string
	var hitpoints string

	npcForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Value(&name).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("You have to enter a name!")
					}
					return nil
				}),

			huh.NewInput().
				Title("Race").
				Value(&race),

			huh.NewInput().
				Title("Class").
				Value(&class),

			huh.NewInput().
				Title("Subclass").
				Value(&subclass),

			huh.NewInput().
				Title("Alignment").
				Value(&alignment),

			huh.NewSelect[string]().
				Title("Sex").
				Options(
					huh.NewOption("Male", "male"),
					huh.NewOption("Female", "female"),
					huh.NewOption("Intersex", "intersex"),
					huh.NewOption("?", "?"),
					huh.NewOption("Other", "other"),
				).
				Value(&sex),

			huh.NewInput().
				Title("List known languages:").
				Value(&languages),
		),

		huh.NewGroup(
			huh.NewText().
				Title("Describe your NPC").
				Value(&desc),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("Level").
				Value(&level).
				Validate(func(s string) error {
					_, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return fmt.Errorf("%s is not a number!", s)
					}
					return nil
				}),

			huh.NewInput().
				Title("Hitpoints").
				Value(&hitpoints).
				Validate(func(s string) error {
					_, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return fmt.Errorf("%s is not a number!", s)
					}
					return nil
				}),
		),
	).WithTheme(huh.ThemeBase16())

	err := npcForm.Run()
	if err != nil {
		return err
	}

	parsedLevel, _ := strconv.ParseInt(level, 10, 64)
	parsedHP, _ := strconv.ParseInt(hitpoints, 10, 64)

	npc := &db.NpcParams{
		Name:        name,
		Race:        race,
		Class:       class,
		Subclass:    subclass,
		Alignment:   alignment,
		Sex:         sex,
		Description: desc,
		Languages:   languages,
		Level:       int(parsedLevel),
		Hitpoints:   int(parsedHP),
	}

	added, err := s.Db.AddNpc(context.Background(), npc)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Success! %s created.\n", added.Name)

	viewNpc(added)

	return nil
}

func viewNpc(npc *db.Npc) {
	intro := fmt.Sprintf("Info About NPC:\n")
	introTip := fmt.Sprintf("%s // Id: %d", npc.Name, npc.Id)
	fmt.Println(header.Render(intro + introTip))
	fmt.Printf(" Name: %s\n", npc.Name)
	fmt.Printf(" Race: %s\n", npc.Race)
	fmt.Printf(" Class: %s\n", npc.Class)
	fmt.Printf(" Subclass: %s\n", npc.Subclass)
	fmt.Printf(" Alignment: %s\n", npc.Alignment)
	fmt.Printf(" Sex: %s\n", npc.Sex)
	fmt.Printf(" Description : %s\n", npc.Description)
	fmt.Printf(" Level: %d\n", npc.Level)
	fmt.Printf(" Hitpoints: %d\n", npc.Hitpoints)
}

func editNpc(npc *db.Npc, s *State) error {
	name := npc.Name
	race := npc.Race
	class := npc.Class
	subclass := npc.Subclass
	alignment := npc.Alignment
	sex := npc.Sex
	desc := npc.Description
	languages := npc.Languages
	level := strconv.FormatInt(int64(npc.Level), 10)
	hitpoints := strconv.FormatInt(int64(npc.Hitpoints), 10)

	npcForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Value(&name).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("You have to enter a name!")
					}
					return nil
				}),

			huh.NewInput().
				Title("Race").
				Value(&race),

			huh.NewInput().
				Title("Class").
				Value(&class),

			huh.NewInput().
				Title("Subclass").
				Value(&subclass),

			huh.NewInput().
				Title("Alignment").
				Value(&alignment),

			huh.NewSelect[string]().
				Title("Sex").
				Options(
					huh.NewOption("Male", "male"),
					huh.NewOption("Female", "female"),
					huh.NewOption("Intersex", "intersex"),
					huh.NewOption("?", "?"),
					huh.NewOption("Other", "other"),
				).
				Value(&sex),

			huh.NewInput().
				Title("List known languages:").
				Value(&languages),
		),

		huh.NewGroup(
			huh.NewText().
				Title("Describe your NPC").
				Value(&desc),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("Level").
				Value(&level).
				Validate(func(s string) error {
					_, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return fmt.Errorf("%s is not a number!", s)
					}
					return nil
				}),

			huh.NewInput().
				Title("Hitpoints").
				Value(&hitpoints).
				Validate(func(s string) error {
					_, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return fmt.Errorf("%s is not a number!", s)
					}
					return nil
				}),
		),
	).WithTheme(huh.ThemeBase())

	err := npcForm.Run()
	if err != nil {
		return err
	}

	parsedLevel, _ := strconv.ParseInt(level, 10, 64)
	parsedHP, _ := strconv.ParseInt(hitpoints, 10, 64)

	updatedNpc := db.Npc{
		Id:          npc.Id,
		Name:        name,
		Race:        race,
		Class:       class,
		Subclass:    subclass,
		Alignment:   alignment,
		Sex:         sex,
		Description: desc,
		Languages:   languages,
		Level:       int(parsedLevel),
		Hitpoints:   int(parsedHP),
	}

	res, err := s.Db.EditNpcById(context.Background(), &updatedNpc)
	if err != nil {
		return err
	}

	viewNpc(res)

	return nil
}
