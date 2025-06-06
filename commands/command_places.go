package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/jake-abed/lore/internal/db"
	"github.com/jake-abed/lore/internal/utils"
)

func commandPlaces(s *State) error {
	args := s.Args[1:]

	// Break out if user did not provide enough flags.
	if len(args) < 2 {
		fmt.Println("Places command requires at least two arguments!")
		placesHelp()
		return nil
	}

	var typeFlag string

	for _, arg := range args {
		if isPlaceTypeFlag(arg) {
			typeFlag = arg
		}
	}

	flag, flagArg := parsePlaceFlagArg(args)

	/* If the user is not searching, they have to have provided a type Flag
	somewhere in the arg list. If not, break out and error. If they are searching
	run the search functionality which will look up all places!
	*/
	if flag != "-s" && typeFlag == "" {
		return fmt.Errorf("flag %s requires a place type flag as well", flag)
	} else if flag == "-s" && flagArg != "" {
		searchTerm := "%" + strings.ToLower(flagArg) + "%"
		err := searchPlaceByName(s, typeFlag, searchTerm)
		if err != nil {
			return err
		}
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
	case "-va":
		err := viewAllPlaces(s, typeFlag)
		if err != nil {
			return err
		}

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
		id, err := strconv.ParseInt(flagArg, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot delete by name. Must delete by ID")
		}
		err = deletePlaceById(s, typeFlag, int(id))
		if err != nil {
			return err
		}

		return nil
	default:
		msg := "uh oh! You used an invalid flag or wrote your command wrong"
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

func printPlaceQuick(p db.Place) {
	switch p.(type) {
	case *db.World:
		world := p.(*db.World)
		printWorldQuick(world)
	case *db.Area:
		area := p.(*db.Area)
		printAreaQuick(area)
	case *db.Location:
		location := p.(*db.Location)
		printLocationQuick(location)
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

func printWorldQuick(w *db.World) {
	s := utils.TruncateString(w.Desc, 32) + "..."
	fmt.Printf("Id: %d | Name: %s | Desc: %s\n", w.Id, w.Name, s)
}

func printArea(a *db.Area) {
	headerMsg := fmt.Sprintf("Area: %-16s Id: %-2d", a.Name, a.Id)
	printHeader(headerMsg)
	fmt.Println(bold.Render("Area Type: ") + a.Type)
	fmt.Println(bold.Render("Description: ") + a.Desc)
	fmt.Println(bold.Render("Belongs to World Id: ") +
		fmt.Sprintf("%d", a.WorldId))
}

func printAreaQuick(a *db.Area) {
	s := utils.TruncateString(a.Desc, 32) + "..."
	msg := fmt.Sprintf("Id: %d | Name: %s | Type: %s | Desc: %s | World Id: %d",
		a.Id, a.Name, a.Type, s, a.WorldId,
	)
	fmt.Println(msg)
}

func printLocation(l *db.Location) {
	headerMsg := fmt.Sprintf("Location: %-16s Id: %-2d", l.Name, l.Id)
	printHeader(headerMsg)
	fmt.Println(bold.Render("Area Type: ") + l.Type)
	fmt.Println(bold.Render("Description: ") + l.Desc)
	fmt.Println(bold.Render("Belongs to Area Id: ") +
		fmt.Sprintf("%d", l.AreaId))
}

func printLocationQuick(l *db.Location) {
	s := utils.TruncateString(l.Desc, 32) + "..."
	msg := fmt.Sprintf("Id: %d | Name: %s | Type: %s | Desc: %s | World Id: %d",
		l.Id, l.Name, l.Type, s, l.AreaId,
	)
	fmt.Println(msg)

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
			return &db.World{}, fmt.Errorf("world not added")
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
		location, err := locationForm(s, db.Location{})
		if err != nil {
			if err.Error() == "user aborted" {
				fmt.Println("User exited Lore form early!")
				os.Exit(0)
			}
			return nil, err
		}

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
	switch p := place.(type) {
	case *db.World:
		world, err := worldForm(*p)
		if err != nil {
			return nil, err
		}

		updatedWorld, err := s.Db.UpdateWorldById(context.Background(), world)
		if err != nil {
			return nil, err
		}

		return updatedWorld, nil
	case *db.Area:
		area, err := areaForm(s, *p)
		if err != nil {
			return nil, err
		}

		updatedArea, err := s.Db.UpdateAreaById(context.Background(), area)
		if err != nil {
			return nil, err
		}

		return updatedArea, nil
	case *db.Location:
		location, err := locationForm(s, *p)
		if err != nil {
			if err.Error() == "user aborted" {
				fmt.Println("User exited Lore form early!")
				os.Exit(0)
			}
			return nil, err
		}

		updatedLocation, err := s.Db.UpdateLocationById(context.Background(),
			location,
		)
		if err != nil {
			return nil, err
		}

		return updatedLocation, err
	default:
		fmt.Println("Uh oh, cannot edit this place.")
		return nil, fmt.Errorf("on editable place type: %T", place)
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

func searchPlaceByName(s *State, placeType string, name string) error {
	switch placeType {
	case "--world":
		worlds, err := s.Db.SearchWorldsByName(
			context.Background(),
			db.SearchParams{Name: name, Limit: 1000, Offset: 0},
		)
		if err != nil {
			return err
		}

		for _, world := range worlds {
			id, worldName := world.Inspect()
			fmt.Printf("World Id: %d | World Name: %s\n", id, worldName)
		}

		return nil
	case "--area":
		areas, err := s.Db.SearchAreasByName(
			context.Background(),
			db.SearchParams{Name: name, Limit: 1000, Offset: 0},
		)
		if err != nil {
			return err
		}

		for _, area := range areas {
			id, areaName := area.Inspect()
			fmt.Printf("Area Id: %d | Area Name: %s\n", id, areaName)
		}

		return nil
	case "--location":
		locations, err := s.Db.SearchLocationsByName(
			context.Background(),
			db.SearchParams{Name: name, Limit: 1000, Offset: 0},
		)
		if err != nil {
			return err
		}

		for _, location := range locations {
			id, locationName := location.Inspect()
			fmt.Printf("Location Id: %d | Location Name: %s\n", id, locationName)
		}

		return nil
	default:
		return fmt.Errorf("%s not recognized when searching", placeType)
	}
}

func deletePlaceById(s *State, placeType string, id int) error {
	switch placeType {
	case "--world":
		msg := fmt.Sprintf("delete World ID '%d'", id)
		cont, _ := confirmForm(msg)
		if !cont {
			return nil
		}
		err := s.Db.DeleteWorldByIdQuery(context.Background(), id)
		return err
	case "--area":
		msg := fmt.Sprintf("delete Area ID '%d'", id)
		cont, _ := confirmForm(msg)
		if !cont {
			return nil
		}
		err := s.Db.DeleteAreaByIdQuery(context.Background(), id)
		return err
	case "--location":
		msg := fmt.Sprintf("delete Location ID '%d'", id)
		cont, _ := confirmForm(msg)
		if !cont {
			return nil
		}
		err := s.Db.DeleteLocationByIdQuery(context.Background(), id)
		return err
	default:
		return fmt.Errorf("place type not recognized, could not delete")
	}
}

func viewAllPlaces(s *State, placeFlag string) error {
	switch placeFlag {
	case "--world":
		worlds, err := s.Db.GetXWorlds(context.Background(), 1_000_000_000, 0)
		if err != nil {
			return nil
		}

		fmt.Println("Viewing all worlds: ")

		for _, w := range worlds {
			printWorldQuick(w)
		}

		return nil
	case "--area":
		areas, err := s.Db.GetAllAreas(context.Background())
		if err != nil {
			return nil
		}

		fmt.Println("Viewing all areas: ")

		for _, a := range areas {
			printAreaQuick(a)
		}

		return nil
	case "--location":
		locations, err := s.Db.GetAllLocations(context.Background())
		if err != nil {
			return nil
		}

		fmt.Println("Viewing all locations: ")

		for _, l := range locations {
			printLocationQuick(l)
		}

		return nil
	default:
		return fmt.Errorf("no such place flag exists")
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
		newPlaceSelectGroup(worlds, area.Name, &area.WorldId),
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

func locationForm(s *State, location db.Location) (db.Location, error) {
	worlds, _ := s.Db.GetXWorlds(context.Background(), 10, 0)
	var worldId int
	worldForm := huh.NewForm(
		newPlaceSelectGroup(worlds, location.Name, &worldId),
	).WithTheme(huh.ThemeBase16())

	err := worldForm.Run()
	if err != nil {
		return db.Location{}, err
	}

	areas, err := s.Db.GetXAreas(context.Background(), worldId, 10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	areaForm := huh.NewForm(
		newPlaceSelectGroup(areas, location.Name, &location.AreaId),
	).WithTheme(huh.ThemeBase16())

	err = areaForm.Run()
	if err != nil {
		return db.Location{}, err
	}

	mainForm := huh.NewForm(
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
	).WithTheme(huh.ThemeBase16())

	err = mainForm.Run()
	if err != nil {
		return db.Location{}, err
	}

	return location, nil
}

func newPlaceSelectGroup(
	places interface{},
	msg string,
	val *int,
) *huh.Group {
	if msg == "" {
		msg = "Which place does this entry belong to?"
	}

	options := []huh.Option[int]{}

	switch places.(type) {
	case []*db.World:
		worlds := places.([]*db.World)
		for _, world := range worlds {
			id, name := world.Inspect()
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
			Title(msg).
			Options(options...).
			Value(val),
	)
}

func confirmForm(msg string) (bool, error) {
	var cont bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Value(&cont).
				Title("Are you sure you would like to " + msg + "?"),
		),
	)

	err := form.Run()
	if err != nil {
		return false, err
	}

	return cont, nil
}

// Flag helper functions

func isPlaceTypeFlag(flag string) bool {
	return flag == "--world" || flag == "--area" || flag == "--location"
}

func isPlaceCommandFlag(flag string) bool {
	return flag == "-a" || flag == "-v" || flag == "-va" || flag == "-e" ||
		flag == "-d" || flag == "-s"
}

func parsePlaceFlagArg(args []string) (string, string) {
	for i, arg := range args {
		if isPlaceCommandFlag(arg) && (1+i) < len(args) {
			return arg, strings.Join(args[i+1:], " ")
		} else if isPlaceCommandFlag(arg) {
			return arg, ""
		}
	}

	return "", ""
}

// Places Help Command

func placesHelp() {
	intro := "Lore Places Help\n"
	introTip := "Places subcommand information"
	fmt.Println(header.Render(intro + introTip))
	placeFlagIntro := bold.Render("Place Flags: ")
	placeFlags := "--world, --area, --location"
	fmt.Println(placeFlagIntro + placeFlags)
	add := bold.Render("  *** places <place flag> -a ")
	addMessage := "| Add a new place. Must specifiy place flag."
	fmt.Println(add + addMessage)
	edit := bold.Render("  *** places <place flag> -e <name> ")
	editMessage := "| Edit a place by name. Must specify place flag."
	fmt.Println(edit + editMessage)
	view := bold.Render("  *** places <place flag> -v <name> ")
	viewMessage := "| View a place by name (case-insensitive)."
	fmt.Println(view + viewMessage)
	viewAll := bold.Render("  *** places <place-flag> -va ")
	viewAllMsg := "| View all places by place-flag. Great for piping."
	fmt.Println(viewAll + viewAllMsg)
	delete := bold.Render("  *** places <place flag> -d <id> ")
	deleteMessage := "| Delete a place by ID."
	fmt.Println(delete + deleteMessage)
	search := bold.Render("  *** places -s <place flag> <name> ")
	searchMessage := "| Searches the DB by place name returning all results.\n"
	fmt.Println(search + searchMessage)
}
