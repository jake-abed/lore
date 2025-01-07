package commands

import (
	"fmt"
)

func commandPlaces(s *State) error {
	args := s.Args[1:]

	if len(args) < 2 {
		return fmt.Errorf("Places command requires at least two arguments!")
	}

	var typeFlag string

	for _, arg := range args {
		if isPlaceTypeFlag(arg) {
			typeFlag = arg
		}
	}

	flag, flagArg := parseFlagArg(args)

	if flag != "-s" && typeFlag == "" {
		return fmt.Errorf("Flag %s requires a place type flag as well.", flag)
	} else if flag == "-s" && flagArg != "" {
		fmt.Println("Add `search fn!")
		return nil
	}

	switch flag {
	case "-a":
		fmt.Println("Add `add` fn!")
		return nil
	case "-v":
		fmt.Println("Add `view` fn!")
		return nil
	case "-e":
		fmt.Println("Add `edit` fn!")
		return nil
	case "-d":
		fmt.Println("Add `delete` fn!")
		return nil
	default:
		fmt.Println("Help!")
		return nil
	}
}

func isPlaceTypeFlag(flag string) bool {
	return flag == "--world" || flag == "--region" || flag == "-location"
}

func isCommandFlag(flag string) bool {
	return flag == "-a" || flag == "-v" || flag == "-e" ||
		flag == "-d" || flag == "-s"
}

func parseFlagArg(args []string) (string, string) {
	for i, arg := range args {
		if isCommandFlag(arg) && (1+i) < len(args) {
			return arg, args[i+1]
		} else if isCommandFlag(arg) {
			return arg, ""
		}
	}

	return "", ""
}
