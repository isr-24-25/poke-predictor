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