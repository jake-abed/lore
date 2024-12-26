package dndapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetAllMonsters() ([]MonsterSearchResult, error) {
	url := baseUrl + "/monsters"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []MonsterSearchResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []MonsterSearchResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []MonsterSearchResult{}, err
	}

	monsterResp := MonsterSearchResp{}
	err = json.Unmarshal(body, &monsterResp)
	if err != nil {
		return []MonsterSearchResult{}, err
	}

	return monsterResp.Results, nil
}

func (c *Client) GetMonster(monster string) (Monster, error) {
	sanitizedId := monsterToId(monster)
	url := fmt.Sprintf("%s/monsters/%s", baseUrl, sanitizedId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Monster{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Monster{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		err = fmt.Errorf(
			"%s not found. Some monsters are not featured in the API. Check for typos.",
			monster,
		)
		return Monster{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Monster{}, err
	}

	monsterInfo := Monster{}
	err = json.Unmarshal(body, &monsterInfo)
	if err != nil {
		return Monster{}, err
	}

	return monsterInfo, nil
}

/*
	Takes an input that may or may not be formatted correctly.

Processes it to ensure that it is formatted as lower case, no whitespace
and "-" to delimit words.
*/
func monsterToId(monster string) string {
	lower := strings.ToLower(monster)
	spacesToDashes := strings.ReplaceAll(lower, " ", "-")
	return strings.TrimSpace(spacesToDashes)
}

/*
I don't entirely think this section should be here? It seems slightly out of
place to me. TODO: Move to a more logical place in the codebase.
*/
type Attack struct {
	Name        string
	AttackBonus int
	Damage      []Damage
}

type AttackDamage struct {
	Name   string
	Type   string
	Damage int
	AttackBonus int
}

func (m *Monster) ParseAttacks() []*Attack {
	attacks := []*Attack{}
	for _, a := range m.Actions {
		if a.Damage == nil {
			continue
		}
		newAttack := &Attack{
			Name:        a.Name,
			AttackBonus: a.AttackBonus,
			Damage:      a.Damage,
		}

		attacks = append(attacks, newAttack)
	}

	return attacks
}

func UseRandomAttack(attacks []*Attack) *AttackDamage {
	if len(attacks) == 0 {
		return nil
	}
	attackIndex := rand.IntN(len(attacks)) - 1
	if attackIndex < 0 {
		attackIndex = 0
	}
	attack := attacks[attackIndex]
	damages := []*AttackDamage{}
	for _, damage := range attack.Damage {
		attackDamage := &AttackDamage{
			Name:   attack.Name,
			Type:   damage.DamageType.Name,
			Damage: rollDamage(damage.DamageDice),
			AttackBonus: attack.AttackBonus,
		}
		damages = append(damages, attackDamage)
	}
	return damages[0]
}

/*
A little ugly in my opinion. This function does too much, but that might be
unavoidable? Part of the ugliness comes from having to handle all the int64
to int conversions.
*/

func parseDamageDice(attackDie string) (
	numDice int64,
	damageDice int64,
	bonus int64,
) {
	splitAtD := strings.Split(attackDie, "d")
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
	} else {
		damageDice, err = strconv.ParseInt(splitAtD[1], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func rollDamage(attackDie string) int {
	numDice, dieSize, bonus := parseDamageDice(attackDie)
	var damageSum int64
	for numDice > 0 {
		damageSum += int64(rand.IntN(int(dieSize)-1) + 1)
		numDice -= 1
	}
	return int(damageSum + bonus)
}
