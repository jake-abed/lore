package commands

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jake-abed/lore/internal/db"
)

var validConnectTypes = [5]string{
	"--npc",
	"--quest",
	"--world",
	"--area",
	"--location",
}

type ConnectArg struct {
	Type string
	Id   int
}

func commandConnect(s *State) error {
	args := s.Args[1:]

	// Break out if user did not provide enough flags.
	if len(args) != 2 {
		fmt.Println("Connect command requires exactly 2 arguments!")
		connectHelp()
		return nil
	}

	firstArg, err := parseConnectArg(args[0])
	if err != nil {
		return err
	}
	secondArg, err := parseConnectArg(args[1])
	if err != nil {
		return err
	}

	switch firstArg.Type {
	case "--npc":
		return connectNpc(s, firstArg, secondArg)
	case "--quest":
		return connectQuest(s, firstArg, secondArg)
	default:
		return fmt.Errorf("first connection arg must be place or npc")
	}
}

func connectNpc(s *State, npcArg, secondArg ConnectArg) error {
	if npcArg.Type != "--npc" {
		return fmt.Errorf("first argument must be an npc")
	}
	npc, err := s.Db.GetNpcById(context.Background(), npcArg.Id)
	if err != nil {
		return fmt.Errorf("no such npc in database: %w", err)
	}

	params := db.ConnectionParams{
		FirstId:  npc.Id,
		SecondId: secondArg.Id,
	}

	switch secondArg.Type {
	case "--quest":
		quest, err := s.Db.GetQuestById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such quest in database: %w", err)
		}

		_, err = s.Db.CreateNpcQuestConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect npc_name=%s(id=%d) to quest_id=%d - error: %w",
				npc.Name,
				npc.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "NPC", Name: npc.Name, Id: npc.Id},
			SuccessEntry{TypeName: "Quest", Name: quest.Name, Id: quest.Id},
		)
		return nil
	case "--world":
		world, err := s.Db.GetWorldById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such world in database: %w", err)
		}

		_, err = s.Db.CreateNpcWorldConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect npc_name=%s(id=%d) to world_id=%d - error: %w",
				npc.Name,
				npc.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "NPC", Name: npc.Name, Id: npc.Id},
			SuccessEntry{TypeName: "World", Name: world.Name, Id: world.Id},
		)
	case "--area":
		fmt.Println(secondArg.Id)
		area, err := s.Db.GetAreaById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such area in database: %w", err)
		}

		_, err = s.Db.CreateNpcAreaConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect npc_name=%s(id=%d) to location_id=%d - error: %w",
				npc.Name,
				npc.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "NPC", Name: npc.Name, Id: npc.Id},
			SuccessEntry{TypeName: "Area", Name: area.Name, Id: area.Id},
		)
	case "--location":
		location, err := s.Db.GetLocationById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such world in database: %w", err)
		}

		_, err = s.Db.CreateNpcAreaConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect npc_name=%s(id=%d) to location_id=%d - error: %w",
				npc.Name,
				npc.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "NPC", Name: npc.Name, Id: npc.Id},
			SuccessEntry{TypeName: "Location", Name: location.Name, Id: location.Id},
		)
	default:
		return fmt.Errorf("unknown Type for second argument: %s", secondArg.Type)
	}

	return nil
}

func connectQuest(s *State, questArg, secondArg ConnectArg) error {
	if questArg.Type != "--quest" {
		return fmt.Errorf("first argument must be a quest")
	}
	quest, err := s.Db.GetQuestById(context.Background(), questArg.Id)
	if err != nil {
		return fmt.Errorf("no such quest in database: %w", err)
	}

	params := db.ConnectionParams{
		FirstId:  quest.Id,
		SecondId: secondArg.Id,
	}

	switch secondArg.Type {
	case "--npc":
		npc, err := s.Db.GetNpcById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such npc in database: %w", err)
		}

		// Basically, swap the parameter order because Db.CreateNpcQuestConnection
		// expects the Npc ID to be under FirstId.
		params.FirstId = secondArg.Id
		params.SecondId = quest.Id

		_, err = s.Db.CreateNpcQuestConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect quest_name=%s(id=%d) to npc_id=%d - error: %w",
				quest.Name,
				quest.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "Quest", Name: quest.Name, Id: quest.Id},
			SuccessEntry{TypeName: "NPC", Name: npc.Name, Id: npc.Id},
		)
	case "--world":
		world, err := s.Db.GetWorldById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such world in database: %w", err)
		}

		_, err = s.Db.CreateQuestWorldConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect quest_name=%s(id=%d) to world_id=%d - error: %w",
				quest.Name,
				quest.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "Quest", Name: quest.Name, Id: quest.Id},
			SuccessEntry{TypeName: "World", Name: world.Name, Id: world.Id},
		)
	case "--area":
		area, err := s.Db.GetAreaById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such area in database: %w", err)
		}

		_, err = s.Db.CreateQuestAreaConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect quest_name=%s(id=%d) to world_id=%d - error: %w",
				quest.Name,
				quest.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "Quest", Name: quest.Name, Id: quest.Id},
			SuccessEntry{TypeName: "Area", Name: area.Name, Id: area.Id},
		)
	case "--location":
		location, err := s.Db.GetAreaById(context.Background(), secondArg.Id)
		if err != nil {
			return fmt.Errorf("no such location in database: %w", err)
		}

		_, err = s.Db.CreateQuestLocationConnection(
			context.Background(),
			params,
		)
		if err != nil {
			return fmt.Errorf(
				"could not connect quest_name=%s(id=%d) to location_id=%d - error: %w",
				quest.Name,
				quest.Id,
				secondArg.Id,
				err,
			)
		}

		printSuccessMessage(
			SuccessEntry{TypeName: "Quest", Name: quest.Name, Id: quest.Id},
			SuccessEntry{TypeName: "Location", Name: location.Name, Id: location.Id},
		)
	default:
		return fmt.Errorf("unknown Type for second argument: %s", secondArg.Type)
	}

	return nil
}

func parseConnectArg(arg string) (ConnectArg, error) {
	splitArg := strings.Split(arg, "=")

	if len(splitArg) != 2 {
		return ConnectArg{}, fmt.Errorf("malformed connection argument")
	}

	argType, err := validateConnectType(splitArg[0])
	if err != nil {
		return ConnectArg{}, err
	}

	id, err := strconv.ParseInt(splitArg[1], 10, 64)
	if err != nil {
		return ConnectArg{}, err
	}

	return ConnectArg{Type: argType, Id: int(id)}, nil
}

func validateConnectType(possibleType string) (string, error) {
	if !slices.Contains(validConnectTypes[:], strings.ToLower(possibleType)) {
		return "", fmt.Errorf("%s is not a valid Connection Type", possibleType)
	}

	return possibleType, nil
}

type SuccessEntry struct {
	Name     string
	TypeName string
	Id       int
}

func printSuccessMessage(entryOne, entryTwo SuccessEntry) {
	intro := "Success!"
	fmt.Println(header.Render(intro))

	idOneMsg := italic.Render(fmt.Sprintf("(ID: %d)", entryOne.Id))
	idTwoMsg := italic.Render(fmt.Sprintf("(ID: %d)", entryTwo.Id))

	msg := fmt.Sprintf("%s %s %s connected to %s %s %s",
		bold.Render(entryOne.TypeName), bold.Render(entryOne.Name),
		idOneMsg, bold.Render(entryTwo.TypeName),
		bold.Render(entryTwo.Name), idTwoMsg)

	fmt.Println(msg)
}

func connectHelp() {
	intro := "Lore Connect Help\n"
	introTip := "Connect subcommand information"
	fmt.Println(header.Render(intro + introTip))

	npc := bold.Render("  *** connect --npc=<id> --{quest|place-type}=<id> ")
	npcMessage := "| Connect an NPC to a quest or place by IDs."
	fmt.Println(npc + npcMessage)

	quest := bold.Render("  *** connect --quest=<id> --{npc|place-type}=<id> ")
	questMessage := "| Connect a quest to an NPC or place by IDs."
	fmt.Println(quest + questMessage)
}

func printNameAndId(name string, id, i, length int) {
	fmt.Printf("%s (ID: %d)", name, id)

	if i < length-1 {
		fmt.Printf(", ")
	} else {
		fmt.Printf("\n")
	}
}
