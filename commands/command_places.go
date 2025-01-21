package commands

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/jake-abed/lore/internal/db"
)

func commandPlaces(s *State) error {
	args := s.Args[1:]

	// Break out if user did not provide enough flags.
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

	/* If the user is not searching, they have to have provided a type Flag
	somewhere in the arg list. If not, break out and error. If they are searching
	run the search functionality which will look up all places!
	*/
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
			fmt.Println("Uh oh! Lore errored out while adding this place: ")
			return err
		}
		printPlace(place)
		return nil
	case "-v":
		place, err := getPlaceByName(s, typeFlag, strings.ToLower(flagArg))
		if err != nil {
			fmt.Printf("Hmm... Lore couldn't find %s. Here's the error: \n", flagArg)
			return err
		}
		printPlace(place)
		return nil
	case "-e":
		place, err := getPlaceByName(s, typeFlag, strings.ToLower(flagArg))
		if err != nil {
			fmt.Printf("Hmm... Lore couldn't find %s. Here's the error: \n", flagArg)
			return err
		}

		updatedPlace, err := editPlace(s, place)
		if err != nil {
			return err
		}

		printPlace(updatedPlace)
		return nil
	case "-d":
		fmt.Println("Add `delete` fn!")
		return nil
	default:
		msg := fmt.Sprintf("Uh oh! You used an invalid flag or wrote your command wrong!")
		fmt.Println(ErrorMsg.Render(msg))
		return nil
	}
}

// Place Printers

func printPlace(p db.Place) {
	switch p.(type) {
	case *db.World:
		world := p.(*db.World)
		printWorld(world)
	case *db.Area:
		area := p.(*db.Area)
		printArea(area)
	case *db.Location:
		location := p.(*db.Location)
		printLocation(location)
	default:
		fmt.Println(p)
		fmt.Printf("Lore has no such place type as %T\n", p)
	}
}

func printWorld(w *db.World) {
	headerMsg := fmt.Sprintf("World: %-16s Id: %-2d", w.Name, w.Id)
	printHeader(headerMsg)
	fmt.Println(bold.Render("Description: ") + w.Desc)
}

func printArea(a *db.Area) {
	headerMsg := fmt.Sprintf("Area: %-16s Id: %-2d", a.Name, a.Id)
	printHeader(headerMsg)
	fmt.Println(bold.Render("Area Type: ") + a.Type)
	fmt.Println(bold.Render("Description: ") + a.Desc)
	fmt.Println(bold.Render("Belongs to World Id: ") +
		fmt.Sprintf("%d", a.WorldId))
}

func printLocation(l *db.Location) {
	headerMsg := fmt.Sprintf("Location: %-16s Id: %-2d", l.Name, l.Id)
	printHeader(headerMsg)
	fmt.Println(bold.Render("Area Type: ") + l.Type)
	fmt.Println(bold.Render("Description: ") + l.Desc)
	fmt.Println(bold.Render("Belongs to Area Id: ") +
		fmt.Sprintf("%d", l.AreaId))
}

/// Add Place Helper Function

func addPlace(s *State, typeFlag string) (db.Place, error) {
	switch typeFlag {
	case "--world":
		world, err := worldForm(db.World{})
		if err != nil {
			if err.Error() == "user aborted" {
				fmt.Println("User exited Lore form early!")
				os.Exit(0)
			}
			return &db.World{}, err
		}
		if world.Name == "" || world.Desc == "" {
			return &db.World{}, fmt.Errorf("World not added.")
		}
		worldParams := db.WorldParams{Name: world.Name, Desc: world.Desc}

		newWorld, err := s.Db.AddWorld(context.Background(), &worldParams)
		if err != nil {
			if err.Error() == "user aborted" {
				fmt.Println("User exited Lore form early!")
				os.Exit(0)
			}
			return &db.World{}, err
		}

		return newWorld, nil
	case "--area":
		area, err := areaForm(s, db.Area{})
		if err != nil {
			if err.Error() == "user aborted" {
				fmt.Println("User exited Lore form early!")
				os.Exit(0)
			}
			return &db.Area{}, err
		}

		areaParams := &db.AreaParams{
			Name:    area.Name,
			Type:    area.Type,
			Desc:    area.Desc,
			WorldId: area.WorldId,
		}

		newArea, err := s.Db.AddArea(context.Background(), areaParams)
		if err != nil {
			return &db.Area{}, err
		}

		return newArea, nil
	case "--location":
		location := locationForm(s, db.Location{})
		locationParams := &db.LocationParams{
			Name:   location.Name,
			Type:   location.Type,
			Desc:   location.Desc,
			AreaId: location.AreaId,
		}

		newLocation, err := s.Db.AddLocation(context.Background(), locationParams)
		if err != nil {
			return &db.Location{}, err
		}

		return newLocation, nil
	default:
		return nil, fmt.Errorf("%s is not a valid typeflag", typeFlag)
	}
}

func editPlace(s *State, place db.Place) (db.Place, error) {
	switch place.(type) {
	case *db.World:
		world := *place.(*db.World)
		world, err := worldForm(world)
		if err != nil {
			return nil, err
		}

		updatedWorld, err := s.Db.UpdateWorldById(context.Background(), world)
		if err != nil {
			return nil, err
		}

		return updatedWorld, nil
	case *db.Area:
		area := *place.(*db.Area)
		area, err := areaForm(s, area)
		if err != nil {
			return nil, err
		}

		updatedArea, err := s.Db.UpdateAreaById(context.Background(), area)
		if err != nil {
			return nil, err
		}

		return updatedArea, nil
	case *db.Location:
		location := *place.(*db.Location)
		location = locationForm(s, location)

		updatedLocation, err := s.Db.UpdateLocationById(context.Background(),
			location,
		)
		if err != nil {
			return nil, err
		}

		return updatedLocation, err
	default:
		fmt.Println("Uh oh, cannot edit this place.")
		return nil, fmt.Errorf("Non editable place type: %T", place)
	}
}

func getPlaceByName(s *State, typeFlag string, arg string) (db.Place, error) {
	switch typeFlag {
	case "--world":
		world, err := s.Db.GetWorldByName(context.Background(), arg)
		if err != nil {
			return nil, err
		}
		return world, nil
	case "--area":
		area, err := s.Db.GetAreaByName(context.Background(), arg)
		if err != nil {
			return nil, err
		}
		return area, nil
	case "--location":
		location, err := s.Db.GetLocationByName(context.Background(), arg)
		if err != nil {
			return nil, err
		}
		return location, nil
	default:
		return nil, nil
	}
}

// Form Functions

func worldForm(world db.World) (db.World, error) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("World Name: ").
				Value(&world.Name),
			huh.NewText().
				Title("World Description: ").
				Value(&world.Desc),
		),
	).WithTheme(huh.ThemeBase16())

	err := form.Run()
	if err != nil {
		return db.World{}, err
	}

	return world, nil
}

func areaForm(s *State, area db.Area) (db.Area, error) {
	worlds, _ := s.Db.GetXWorlds(context.Background(), 10, 0)
	form := huh.NewForm(
		newPlaceSelectGroup(worlds, &area.WorldId),
		huh.NewGroup(
			huh.NewInput().
				Title("Area Name: ").
				Value(&area.Name),
			huh.NewInput().
				Title("Area Type: ").
				Value(&area.Type),
			huh.NewText().
				Title("Area Description: ").
				Value(&area.Desc),
		),
	).WithTheme(huh.ThemeBase16())

	err := form.Run()
	if err != nil {
		return db.Area{}, err
	}

	return area, nil
}

func locationForm(s *State, location db.Location) db.Location {
	worlds, _ := s.Db.GetXWorlds(context.Background(), 10, 0)
	var worldId int
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Location Name: ").
				Value(&location.Name),
			huh.NewInput().
				Title("Location Type: ").
				Value(&location.Type),
			huh.NewText().
				Title("Location Description: ").
				Value(&location.Desc),
		),
		newPlaceSelectGroup(worlds, &worldId),
	).WithTheme(huh.ThemeBase16()).Run()

	areas, err := s.Db.GetXAreas(context.Background(), worldId, 10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	huh.NewForm(
		newPlaceSelectGroup(areas, &location.AreaId),
	).WithTheme(huh.ThemeBase16()).Run()

	return location
}

func newPlaceSelectGroup(places interface{}, val *int) *huh.Group {
	options := []huh.Option[int]{}
	var placeType db.PlaceType

	switch places.(type) {
	case []*db.World:
		worlds := places.([]*db.World)
		for _, world := range worlds {
			id, name := world.Inspect()
			placeType = world.PlaceType()
			option := huh.NewOption(fmt.Sprintf("%d - %s", id, name), id)

			options = append(options, option)
		}
	case []*db.Area:
		areas := places.([]*db.Area)
		for _, area := range areas {
			id, name := area.Inspect()
			option := huh.NewOption(fmt.Sprintf("%d - %s", id, name), id)

			options = append(options, option)
		}
	default:
		fmt.Println(ErrorMsg.Render("Uh oh! You need a place this can belong to!"))
		os.Exit(1)
	}

	return huh.NewGroup(
		huh.NewSelect[int]().
			Title(fmt.Sprintf("Which %s will this place belong to?", placeType)).
			Options(options...).
			Value(val),
	)
}

// Flag helper functions

func isPlaceTypeFlag(flag string) bool {
	return flag == "--world" || flag == "--area" || flag == "--location" ||
		flag == "--sublocation"
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
