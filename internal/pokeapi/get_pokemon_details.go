package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails (pokemonName * string) (Pokemon, error){
	url := baseURL + "/pokemon/" + *pokemonName
	getPokemonDetailsResp := Pokemon{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	err = json.Unmarshal(dat, &getPokemonDetailsResp)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	return getPokemonDetailsResp, nil
}