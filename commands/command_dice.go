package commands

import (
	"fmt"

	"github.com/jake-abed/lore/internal/dice"
)

func commandDice(s *State) error {
	diceArgs := s.Args[1:]

	var flag string
	var diceExpression string

	if len(diceArgs) == 1 && diceArgs[0][0] != 45 {
		fmt.Println("just default roll")
		flag = "-a"
		diceExpression = diceArgs[0]
	}

	fmt.Println(flag, diceExpression)

	fmt.Println(diceArgs)
	fmt.Println(dice.SumRollDice("2d6-1"))
	fmt.Println("Dice command place holder!")
	return nil
}
