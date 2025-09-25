package pokeapi


type LocationArea struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type CatchPokemon struct {
	ID int `json:"id"`
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationAreaDetails struct {
	ID int `json:"id"`
	Name string `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type ResponseData struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []LocationArea `json:"results"`
}