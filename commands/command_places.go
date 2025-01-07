package commands

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/jake-abed/lore/internal/db"
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
		place, err := addPlace(s, typeFlag)
		if err != nil {
			return err
		}
		fmt.Println(place)
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

func addPlace(s *State, typeFlag string) (db.Place, error) {
	switch typeFlag {
	case "--world":
		world := worldForm(db.World{})
		worldParams := db.WorldParams{Name: world.Name, Desc: world.Desc}

		newWorld, err := s.Db.AddWorld(context.Background(), &worldParams)
		if err != nil {
			return &db.World{}, err
		}

		return newWorld, nil
	default:
		return nil, fmt.Errorf("%s is not a valid typeflag", typeFlag)
	}
}

// Form Functions

func worldForm(world db.World) db.World {
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("World Name: ").
				Value(&world.Name),
			huh.NewText().
				Title("World Description: ").
				Value(&world.Desc),
		),
	).WithTheme(huh.ThemeBase16()).Run()
	return world
}

// Flag helper functions

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
