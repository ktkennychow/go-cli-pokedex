package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocationAreas (pageURL *string) (RespLocationAreas, error){
	// first page URL
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	locationAreasResp := RespLocationAreas{}
	dat, exist:= c.cache.Get(url)
	if exist {
		// Simulate Accessing
		fmt.Println("Accessing from cache...")
		time.Sleep(200 * time.Millisecond)
	}
	if !exist {
		// Simulate Fetching
		fmt.Println("Fetching from the Internet...")
		fmt.Println(url)
		time.Sleep(1 * time.Second)
		
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
			return RespLocationAreas{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return RespLocationAreas{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return RespLocationAreas{}, err
		}
		c.cache.Add(url, dat)
	}

	err := json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		fmt.Println(err)
		return RespLocationAreas{}, err
	}

	return locationAreasResp, nil
}