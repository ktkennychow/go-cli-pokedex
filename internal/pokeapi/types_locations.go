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
	
type Pokemon struct {
	Name string `json:"name"`
	Height int
	Weight int
	Stats []Stat
	Types []Type
	Base_experience int `json:"base_experience"`
}

type Stat struct {
	Base_stat int `json:"base_stat"`
	Stat struct{
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Type struct{
		Name string `json:"name"`
	} `json:"type"`
}