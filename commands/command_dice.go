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
		return fmt.Errorf("the dice command requires at least one argument")
	}

	if len(diceArgs) == 1 && diceArgs[0] == "help" {
		diceHelp()
		return nil
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
		return fmt.Errorf("invalid Dice Expression: %s", diceExpression)
	}

	switch flag {
	case "-a":
		num, dieVal, modifier := dice.ParseDiceExpression(diceExpression)
		var sum int
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
		msg := "Sum of all rolls & modifier: "
		msg = bold.Render(msg)
		fmt.Println(msg, sum+int(modifier))
	default:
		return fmt.Errorf("unknown flag for dice subcommand")
	}
	return nil
}

func diceHelp() {
	intro := "Lore Dice Help\n"
	introTip := "Dice subcommands information"
	fmt.Println(header.Render(intro + introTip))
	all := bold.Render("  *** dice -a <monster-1> <monster-2> | ")
	allMessage := "Roll all the dice of an expression together returning the result."
	fmt.Println(all + allMessage)
	individual := bold.Render("  *** dice -i <dice-expression> | ")
	individualMessage := "Roll each die of an expression one at a time."
	fmt.Println(individual + individualMessage)
	example := bold.Render("  *** Examples | ")
	exampleMessage := "d20, 1d6, 2d12+1, 1d4-1, 3d6+8, 5d20+12, 80d100-50, etc.\n"
	fmt.Println(example + exampleMessage)
}
