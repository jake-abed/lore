package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/jake-abed/lore/internal/db"
)

func commandQuests(s *State) error {
	args := s.Args[1:]

	// Break out if user did not provide enough flags.
	if len(args) == 0 {
		fmt.Println("Quests command requires at least one argument!")
		return nil
	}

	var stepFlag string

	for _, arg := range args {
		if isStepFlag(arg) {
			stepFlag = arg
		}
	}

	flag, flagArg := parseQuestFlagArg(args)

	if stepFlag == "--step" {
		fmt.Println("Implement step subcommands.")
		return nil
	}

	switch flag {
	case "-a":
		quest, err := addQuest(s)
		if err != nil {
			return err
		}

		fmt.Println(quest)
	default:
		return fmt.Errorf("Unrecognized flag for quests command!")
	}

	fmt.Println(stepFlag, flag, flagArg)

	fmt.Println("Placeholder for quest command!")
	return nil
}

func addQuest(s *State) (*db.Quest, error) {
	quest, err := questForm(s, db.Quest{})
	if err != nil {
		return nil, err
	}

	questParams := db.QuestParams{
		Name:        quest.Name,
		Desc:        quest.Desc,
		Rewards:     quest.Rewards,
		Notes:       quest.Notes,
		Level:       quest.Level,
		IsFinished:  quest.IsFinished,
		IsStarted:   quest.IsStarted,
		CurrentStep: quest.CurrentStep,
		WorldId:     quest.WorldId,
	}

	newQuest, err := s.Db.AddQuest(context.Background(), &questParams)
	if err != nil {
		return nil, nil
	}

	return newQuest, nil
}

// Forms
func questForm(s *State, quest db.Quest) (db.Quest, error) {
	worldCount, _ := s.Db.WorldCount(context.Background())
	if worldCount == 0 {
		return db.Quest{}, fmt.Errorf("Please create a world to add quests to!")
	}

	worlds, err := s.Db.GetXWorlds(context.Background(), 300, 0)
	if err != nil {
		return db.Quest{}, err
	}

	var level string

	nameForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Quest Name: ").
				Value(&quest.Name),
		),
		newPlaceSelectGroup(worlds, "Which world does this quest belong to?", &quest.WorldId),
		huh.NewGroup(
			huh.NewText().
				Title("Quest Description: ").
				Value(&quest.Desc),
			huh.NewText().
				Title("Quest Rewards: ").
				Value(&quest.Rewards),
			huh.NewText().
				Title("Quest Notes: ").
				Value(&quest.Notes),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Suggested Level for Quest: ").
				Value(&level).
				Validate(func(s string) error {
					i, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return fmt.Errorf("Must be a number!")
					}

					num := int(i)
					if num < 0 {
						return fmt.Errorf("Must be greater than or equal to 0!")
					}

					return nil
				}),
			huh.NewConfirm().
				Title("Is the quest started?").
				Value(&quest.IsStarted),
			huh.NewConfirm().
				Title("Is the quest finished?").
				Value(&quest.IsFinished),
		),
	).WithTheme(huh.ThemeBase16())

	err = nameForm.Run()
	if err != nil {
		return db.Quest{}, err
	}

	i, err := strconv.ParseInt(level, 10, 64)
	if err != nil {
		return db.Quest{}, err
	}

	quest.Level = int(i)

	return quest, nil
}

// Flag Helpers

func isStepFlag(arg string) bool {
	return arg == "--step"
}

func isQuestCommandFlag(flag string) bool {
	return flag == "-a" || flag == "-v" || flag == "-e" ||
		flag == "-d" || flag == "-s"
}

func parseQuestFlagArg(args []string) (string, string) {
	for i, arg := range args {
		if isQuestCommandFlag(arg) && (1+i) < len(args) {
			return arg, args[i+1]
		} else if isQuestCommandFlag(arg) {
			return arg, ""
		}
	}

	return "", ""
}
