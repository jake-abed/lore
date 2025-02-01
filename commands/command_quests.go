package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/jake-abed/lore/internal/db"
)

func commandQuests(s *State) error {
	args := s.Args[1:]

	// Break out if user did not provide enough flags.
	if len(args) == 0 {
		msg := ErrorMsg.Render("Quests command requires at least one argument!")
		fmt.Println(msg)
		questsHelp()
		return nil
	}

	flag, flagArg := parseQuestFlagArg(args)

	switch flag {
	case "-a":
		quest, err := addQuest(s)
		if err != nil {
			return err
		}

		printQuest(quest)

		return nil
	case "-e":
		quest, err := updateQuest(s, flagArg)
		if err != nil {
			return err
		}

		printQuest(quest)

		return nil
	case "-s":
		quests, err := searchQuests(s, flagArg)
		if err != nil {
			return err
		}

		for _, q := range quests {
			printQuestQuick(q)
		}

		return nil
	case "-v":
		quest, err := getQuestById(s, flagArg)
		if err != nil {
			return err
		}

		printQuest(quest)

		return nil
	case "-va":
		quests, err := getAllQuests(s)
		if err != nil {
			return err
		}

		for _, q := range quests {
			printQuestQuick(q)
		}

		return nil
	case "-d":
		err := deleteQuest(s, flagArg)
		if err != nil {
			return err
		}

		return nil
	default:
		fmt.Println("Unrecognized flag for quests command!")
	}

	questsHelp()
	return nil
}

func addQuest(s *State) (*db.Quest, error) {
	quest, err := questForm(s, db.Quest{})
	if err != nil {
		return nil, err
	}

	questParams := db.QuestParams{
		Name:       quest.Name,
		Desc:       quest.Desc,
		Rewards:    quest.Rewards,
		Notes:      quest.Notes,
		Level:      quest.Level,
		IsFinished: quest.IsFinished,
		IsStarted:  quest.IsStarted,
		WorldId:    quest.WorldId,
	}

	newQuest, err := s.Db.AddQuest(context.Background(), &questParams)
	if err != nil {
		if err.Error() == "user aborted" {
			fmt.Println("User exited Lore form early!")
			os.Exit(0)
		}
		return nil, err
	}

	return newQuest, nil
}

func updateQuest(s *State, id string) (*db.Quest, error) {
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	i := int(id64)

	quest, err := s.Db.GetQuestByIdQuery(context.Background(), i)
	if err != nil {
		return nil, err
	}

	updatedQuest, err := questForm(s, *quest)
	if err != nil {
		return nil, err
	}

	params := db.UpdateQuestParams{
		Name:       updatedQuest.Name,
		Desc:       updatedQuest.Desc,
		Rewards:    updatedQuest.Rewards,
		Notes:      updatedQuest.Notes,
		Level:      updatedQuest.Level,
		IsFinished: updatedQuest.IsFinished,
		IsStarted:  updatedQuest.IsStarted,
		WorldId:    updatedQuest.WorldId,
		Id:         updatedQuest.Id,
	}

	finalQuest, err := s.Db.UpdateQuestById(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return finalQuest, nil
}

func getQuestById(s *State, id string) (*db.Quest, error) {
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	i := int(id64)

	quest, err := s.Db.GetQuestByIdQuery(context.Background(), i)
	if err != nil {
		return nil, err
	}

	return quest, nil
}

// Technically, getAllQuests is grabbing '1,000,000,000' quests. The chance
// that a user is somehow getting storing more than that locally is next to
// nil in my opinion.
func getAllQuests(s *State) ([]*db.Quest, error) {
	quests, err := s.Db.GetXQuests(context.Background(), 1_000_000_000, 0)
	if err != nil {
		return nil, err
	}

	return quests, nil
}

func searchQuests(s *State, name string) ([]*db.Quest, error) {
	query := "%" + name + "%"

	quests, err := s.Db.GetQuestsByName(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return quests, nil
}

func deleteQuest(s *State, id string) error {
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	i := int(id64)

	quest, err := s.Db.GetQuestByIdQuery(context.Background(), i)
	if err != nil {
		return fmt.Errorf("Quest ID=%d not found!", i)
	}

	confirmMsg := fmt.Sprintf("Do you wish to delete the quest: '%s'?",
		quest.Name,
	)
	proceed, err := confirmForm(confirmMsg)
	if !proceed {
		return nil
	}

	err = s.Db.DeleteQuestById(context.Background(), i)
	if err != nil {
		return err
	}

	return nil

}

// Print functions

func printQuest(q *db.Quest) {
	var started string
	var finished string

	if q.IsStarted {
		started = bold.Render("Quest Has Been Started!")
	} else {
		started = bold.Render("Quest Has Not Been Started!")
	}

	if q.IsFinished {
		finished = bold.Render("Quest Is Finished!")
	} else {
		finished = bold.Render("Quest Is Not Finished!")
	}

	headerMsg := fmt.Sprintf("Quest: %-16s Id: %-2d", q.Name, q.Id)
	printHeader(headerMsg)
	fmt.Println(bold.Render("Description:"))
	fmt.Println(q.Desc)
	fmt.Println(bold.Render("Rewards:"))
	fmt.Println(q.Rewards)
	fmt.Println(bold.Render("Notes:"))
	fmt.Println(q.Notes)
	fmt.Println(bold.Render("Quest Level: ") +
		fmt.Sprintf("%d", q.Level))
	fmt.Println(bold.Render("Belongs to World Id: ") +
		fmt.Sprintf("%d", q.WorldId))
	fmt.Println(started)
	fmt.Println(finished)
}

func printQuestQuick(q *db.Quest) {
	descRunes := []rune(q.Desc)
	shortDesc := string(descRunes[0:31]) + "..."
	fmt.Printf("Id: %d | Name: %s | Desc: %s | Level: %d | World Id: %d\n",
		q.Id, q.Name, shortDesc, q.Level, q.WorldId,
	)
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

	level := fmt.Sprintf("%d", quest.Level)

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

func isQuestCommandFlag(flag string) bool {
	return flag == "-a" || flag == "-v" || flag == "-e" ||
		flag == "-d" || flag == "-s" || flag == "-va"
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

// Quests Help Command

func questsHelp() {
	intro := "Lore Quests Help\n"
	introTip := "Quests subcommand information"
	fmt.Println(header.Render(intro + introTip))
	add := bold.Render("  *** quests -a | ")
	addMessage := "Add a new quest."
	fmt.Println(add + addMessage)
	edit := bold.Render("  *** quests -e <id> | ")
	editMessage := "Edit a place by integer ID."
	fmt.Println(edit + editMessage)
	view := bold.Render("  *** quests -v <id> | ")
	viewMessage := "View a quest by integer ID."
	fmt.Println(view + viewMessage)
	delete := bold.Render("  *** quests -d <id> | ")
	deleteMessage := "Delete a quest by ID."
	fmt.Println(delete + deleteMessage)
	search := bold.Render("  *** quests -s <name> | ")
	searchMessage := "Searches the DB by quest name returning all results.\n"
	fmt.Println(search + searchMessage)
}
