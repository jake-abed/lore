package dice

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

/*
A little ugly in my opinion. This function does too much, but that might be
unavoidable? Part of the ugliness comes from having to handle all the int64
to int conversions.
*/

/*
ParseDiceExpression takes a string diceExpression, and parses 3 strings from
it! The number of dice being rolled {X}, the value of the dice being rolled {Y},
and the modifier being applied to the roll {Z}.
An input dice roll should look like one of the following:
- XdY: roll X dice of Y value.
- XdY+Z: roll X dice of Y value adding Z.
- XdY-Z: roll X dice of Y value subtracting Z.
*/
func ParseDiceExpression(diceExpression string) (
	numDice int64,
	damageDice int64,
	bonus int64,
) {
	if strings.HasPrefix(diceExpression, "d") {
		trimmed := strings.ReplaceAll(diceExpression, "d", "")
		if strings.Contains(trimmed, "+") {
			splitAtPlus := strings.Split(trimmed, "+")
			damageDice, err := strconv.ParseInt(splitAtPlus[0], 10, 32)
			if err != nil {
				fmt.Println(err)
				return 0, 0, 0
			}

			bonus, err := strconv.ParseInt(splitAtPlus[1], 10, 32)
			if err != nil {
				fmt.Println(err)
				return 0, 0, 0
			}

			return 1, damageDice, bonus
		} else if strings.Contains(trimmed, "-") {
			splitAtMinus := strings.Split(trimmed, "-")
			damageDice, err := strconv.ParseInt(splitAtMinus[0], 10, 32)
			if err != nil {
				fmt.Println(err)
				return 0, 0, 0
			}

			bonus, err := strconv.ParseInt(splitAtMinus[1], 10, 32)
			if err != nil {
				fmt.Println(err)
				return 0, 0, 0
			}

			return 1, damageDice, bonus * -1
		} else {
			damageDice, err := strconv.ParseInt(trimmed, 10, 32)
			if err != nil {
				fmt.Println(err)
				return 0, 0, 0
			}

			return 1, damageDice, 0
		}
	}
	splitAtD := strings.Split(diceExpression, "d")
	numDice, err := strconv.ParseInt(splitAtD[0], 10, 32)
	if err != nil {
		fmt.Println(err)
		return 0, 0, 0
	}
	dieSizeStr := splitAtD[1]
	if strings.Contains(dieSizeStr, "+") {
		splitAtPlus := strings.Split(dieSizeStr, "+")
		damageDice, err = strconv.ParseInt(splitAtPlus[0], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		bonus, err = strconv.ParseInt(splitAtPlus[1], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if strings.Contains(dieSizeStr, "-") {
		splitAtMinus := strings.Split(dieSizeStr, "-")
		damageDice, err = strconv.ParseInt(splitAtMinus[0], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		bonus, err = strconv.ParseInt(splitAtMinus[1], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		bonus *= int64(-1)
	} else {
		damageDice, err = strconv.ParseInt(splitAtD[1], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func SumRollDice(diceExpression string) int {
	numDice, dieSize, bonus := ParseDiceExpression(diceExpression)
	var damageSum int64
	for numDice > 0 {
		damageSum += int64(rand.IntN(int(dieSize)-1) + 1)
		numDice -= 1
	}
	return int(damageSum + bonus)
}
