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

func ParseRoll(roll string) (
	numDice int64,
	damageDice int64,
	bonus int64,
) {
	splitAtD := strings.Split(roll, "d")
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

func RollDamage(attackDie string) int {
	numDice, dieSize, bonus := ParseRoll(attackDie)
	var damageSum int64
	for numDice > 0 {
		damageSum += int64(rand.IntN(int(dieSize)-1) + 1)
		numDice -= 1
	}
	return int(damageSum + bonus)
}
