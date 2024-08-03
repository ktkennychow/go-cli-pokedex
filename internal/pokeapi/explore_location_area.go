package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ExploreLocationArea (areaName * string) (RespExploreLocation, error){
	url := baseURL + "/location-area/" + *areaName
	exploreLocationResp := RespExploreLocation{}
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
			return RespExploreLocation{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return RespExploreLocation{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return RespExploreLocation{}, err
		}
		c.cache.Add(url, dat)
	}

	err := json.Unmarshal(dat, &exploreLocationResp)
	if err != nil {
		fmt.Println(err)
		return RespExploreLocation{}, err
	}

	return exploreLocationResp, nil
}