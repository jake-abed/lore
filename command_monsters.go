package main

import (
	"fmt"
	"github.com/jake-abed/auxquest/internals/dndapi"
	"time"
)

func commandMonsters(state *State) error {
	client := dndapi.NewClient(5 * time.Second)
	
	monsters, err := client.GetAllMonsters()
	if err != nil {
		return err
	}

	fmt.Println(bold.Render("Showing all monsters!"))
	for _, monster := range monsters {	
		fmt.Printf("Monster: %s - Index: %s\n", monster.Name, monster.Index)
	}
	return nil
}
