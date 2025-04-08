package commands

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
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
		fmt.Println("Places command requires exactly 2 arguments!")
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
		return connectNpc(firstArg, secondArg)
	case "--quest":
		return connectQuest(firstArg, secondArg)
	default:
		return fmt.Errorf("first connection arg must be place or npc")
	}
}

func connectNpc(npcArg, secondArg ConnectArg) error {
	// Implement Connect NPC
	return nil
}

func connectQuest(questArg, secondArg ConnectArg) error {
	// Implement Connect Quest
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

func connectHelp() {
	intro := "Lore Connect Help\n"
	introTip := "Connect subcommand information"
	fmt.Println(header.Render(intro + introTip))
}
