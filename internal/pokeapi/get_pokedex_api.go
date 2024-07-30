package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Res struct {
	Next string
	Previous string
	LocationAreas []LocationArea `json:"results"`
}
type LocationArea struct {
	Name string
}


func GetPokedexApi (url string) (Res, error){
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Res{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Res{}, err
	}
	
	currentResult := Res{}
	fmt.Println(body)
	err = json.Unmarshal(body, &currentResult)
	if err != nil {
		fmt.Println(err)
		return Res{}, err
	}
	res.Body.Close()

	return currentResult, nil
}