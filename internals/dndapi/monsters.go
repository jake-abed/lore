package dndapi

import (
	"net/http"
	"io"
	"encoding/json"
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
