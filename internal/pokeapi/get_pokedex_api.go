package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ktkennychow/go-cli-pokedex/internal/pokecache"
)

type Res struct {
	Next string
	Previous string
	LocationAreas []LocationArea `json:"results"`
}
type LocationArea struct {
	Name string
}

func GetPokedexApi (url string, cache *pokecache.Cache) (Res, error){
	currentResult := Res{}
	body, exist := cache.Get(url)
	if !exist {
		// Simulate Fetching
		fmt.Println("Fetching from the Internet...")
		fmt.Println(url)
		time.Sleep(1 * time.Second)
		
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return Res{}, err
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return Res{}, err
		}
		
		cache.Add(url, body)
		res.Body.Close()
	}
	if exist {
		// Simulate Accessing
		fmt.Println("Accessing from cache...")
		time.Sleep(200 * time.Millisecond)
	}
	err := json.Unmarshal(body, &currentResult)
	if err != nil {
		fmt.Println(err)
		return Res{}, err
	}

	return currentResult, nil
}