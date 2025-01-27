package commands

import (
	"fmt"

	"github.com/jake-abed/lore/internal/dice"
)

func commandDice(s *State) error {
	fmt.Println(dice.SumRollDice("2d6-1"))
	fmt.Println("Dice command place holder!")
	return nil
}
