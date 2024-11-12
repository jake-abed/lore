package dndapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		err = fmt.Errorf("%s not found. Please check for mistakes", monster)
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
