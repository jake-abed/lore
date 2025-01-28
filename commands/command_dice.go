package commands

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/jake-abed/lore/internal/dice"
)

func commandDice(s *State) error {
	diceArgs := s.Args[1:]

	if len(diceArgs) == 0 {
		return fmt.Errorf("The dice command requires at least one argument!")
	}

	var flag string
	var diceExpression string

	if len(diceArgs) == 1 && diceArgs[0][0] != 45 {
		flag = "-a"
		diceExpression = diceArgs[0]
	} else {
		flag = diceArgs[0]
		diceExpression = strings.ToLower(diceArgs[1])
	}

	if !strings.Contains(diceExpression, "d") {
		return fmt.Errorf("Invalid Dice Expression: %s", diceExpression)
	}

	switch flag {
	case "-a":
		num, dieVal, modifier := dice.ParseDiceExpression(diceExpression)
		var sum int
		fmt.Println(dieVal)
		for range num {
			sum += rand.IntN(int(dieVal)) + 1
		}
		msg := fmt.Sprintf("Result of %s roll: ", diceExpression)
		msg = bold.Render(msg)
		fmt.Println(msg, sum+int(modifier))
	case "-i":
		num, dieVal, modifier := dice.ParseDiceExpression(diceExpression)
		if num == 0 {
			num = 1
		}
		var sum int
		for i := range num {
			roll := rand.IntN(int(dieVal)) + 1
			sum += roll
			msg := fmt.Sprintf("Roll %d: ", i+1)
			msg = bold.Render(msg)
			fmt.Println(msg, roll)
		}
		msg := fmt.Sprintf("Sum of all rolls & modifier: ")
		msg = bold.Render(msg)
		fmt.Println(msg, sum+int(modifier))
	default:
		return fmt.Errorf("Unknown flag for dice subcommand!")
	}
	return nil
}
