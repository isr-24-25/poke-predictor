package main

import "context"

func main() {
	conn := Connect()
	defer conn.Close(context.Background())

	stats := GetPokemonBaseStats(1)
	types := GetPokemonTypes(1)

	for i := range stats {
		pokemonID := InsertPokemon(conn, stats[i])
		InsertPokemonType(conn, types)
		InsertPokemonTypeLink(conn, pokemonID, types[i])
	}
}
