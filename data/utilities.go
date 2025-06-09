package main

var (
	Generations = map[uint32]uint32{
		1: 151,
		2: 100,
		3: 135,
		4: 107,
		5: 156,
		6: 72,
		7: 88,
		8: 96,
		9: 120,
	}

	PokemonEndpoint = "https://pokeapi.co/api/v2/"
)

type PokeAPIResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type BaseStats struct {
	ID         int
	Name       string
	Generation int
	HP         int
	Attack     int
	Defense    int
	SpAttack   int
	SpDefense  int
	Speed      int
}