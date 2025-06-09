package main

import "context"

func main() {
	conn := Connect()

	stats := GetPokemonBaseStats(1)
	types := GetPokemonTypes(1)

	for i := range stats {
		InsertPokemon(conn, stats[i])
	}
	InsertPokemonType(conn, types)

	defer conn.Close(context.Background())
}