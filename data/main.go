package main

import "context"

func main() {
	conn := Connect()

	// collect the data
	stats := GetPokemonData(1)

	for i := range stats {
		AddEntry(conn, stats[i])
	}

	defer conn.Close(context.Background())
}