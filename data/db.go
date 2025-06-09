package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/lpernett/godotenv"
)

func Connect() *pgx.Conn {
	godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func InsertPokemon(conn *pgx.Conn, pokemon BaseStats) int {
	query := `INSERT INTO base_stats (id, name, generation, hp, attack, defense, sp_attack, sp_defense, speed) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO NOTHING`

	_, err := conn.Exec(
		context.Background(), 
		query, 
		pokemon.ID, 
		pokemon.Name, 
		pokemon.Generation, 
		pokemon.HP,
		pokemon.Attack,
		pokemon.Defense,
		pokemon.SpAttack,
		pokemon.SpDefense,
		pokemon.Speed,
	)

	if err != nil {
		log.Fatal(err)
	}

	return pokemon.ID
}

func InsertPokemonType(conn *pgx.Conn, pokemonTypes [][]string) {
	var count int
	err := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM types").Scan(&count)
	if err != nil {
		log.Fatalf("failed to count rows in types table: %v", err)
	}
	if count > 0 {
		log.Println("types table already populated. skipping insert.")
		return
	}

	query := `INSERT INTO types (type) VALUES ($1)`	

	for i := range pokemonTypes {
		if len(pokemonTypes[i]) == 1 {
			conn.Exec(context.Background(), query, pokemonTypes[i][0])
		} else {
			for j := range pokemonTypes[i] {
				_, err := conn.Exec(context.Background(), query, pokemonTypes[i][j])
				if err != nil {
					log.Fatal(err)	
				}
			}
		}
	}
}

func InsertPokemonTypeLink(conn *pgx.Conn, pokemonID int, typeNames []string) {
	joinTableQuery := `INSERT INTO pokemon_types_link (pokemon_id, type_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING`

	for _, typeName := range typeNames {
		var typeID int
		err := conn.QueryRow(
			context.Background(), 
			"SELECT id FROM types WHERE type=$1",
			typeName,
		).Scan(&typeID)
		if err != nil {
			log.Fatalf("failed to get type ID for %s: %v", typeName, err)
		}

		_, err = conn.Exec(context.Background(), joinTableQuery, pokemonID, typeID)
		if err != nil {
			log.Fatalf("failed to insert pokemon_types link: %v", err)
		}
	}
}
