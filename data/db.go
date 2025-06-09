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

func AddEntry(conn *pgx.Conn, pokemon BaseStats) {
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
}