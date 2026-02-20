package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespLocationPokemons, error) {
	url := baseURL + "/location-area/" + locationName
	locationResp := RespLocationPokemons{}

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocationPokemons{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationPokemons{}, err
	}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
