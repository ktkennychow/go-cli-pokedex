package pokeapi

type RespLocationAreas struct {
	Count int 							`json:"count"`
	Next *string 						`json:"next"`
	Previous *string					`json:"previous"`
	Results []LocationArea 	`json:"results"`
}

type LocationArea struct {
	Name string 						`json:"name"`
	URL string 							`json:"url"`
}

type RespExploreLocation struct {
	Pokemon_encounters []Encounter 	`json:"pokemon_encounters"`
}


type Encounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon"`
}
