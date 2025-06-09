package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func GetPokemonBaseStats(generation uint32) []BaseStats {
	var stats []BaseStats;

	for k := 1; k <= int(Generations[generation]); k++ {
		formattedEndpoint := fmt.Sprintf("%s/%s", PokemonEndpoint, fmt.Sprintf("pokemon/%d", k))
		response, err := http.Get(formattedEndpoint)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		
		var apiResponse PokeAPIResponse; 
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			log.Fatal(err)
		}

		var b BaseStats;
		b.ID = apiResponse.ID
		b.Name = apiResponse.Name
		b.Generation = int(generation)
		for _, s := range apiResponse.Stats {
			switch s.Stat.Name {
			case "hp":
				b.HP = s.BaseStat
			case "attack":
				b.Attack = s.BaseStat
			case "defense":
				b.Defense = s.BaseStat
			case "special-attack":
				b.SpAttack = s.BaseStat
			case "special-defense":
				b.SpDefense = s.BaseStat
			case "speed":
				b.Speed = s.BaseStat
			}
		}

		stats = append(stats, b)
	}

	return stats
}

func GetPokemonTypes(generation uint32) [][]string {
	var pokemonTypes [][]string;

	for k := 1; k <= int(Generations[generation]); k++ {
		formattedEndpoint := fmt.Sprintf("%s/%s", PokemonEndpoint, fmt.Sprintf("pokemon/%d", k))
		response, err := http.Get(formattedEndpoint)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		var apiResponse PokeAPIResponse; 
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			log.Fatal(err)
		}

		var typeNames []string
		for _, t := range apiResponse.Types {
			typeNames = append(typeNames, t.Type.Name)
		}
		
		pokemonTypes = append(pokemonTypes, typeNames)
	}

	return pokemonTypes
}