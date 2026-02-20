package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/vaeho/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		println("Please specify a Pokemon to catch")
		return nil
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		println("Error fetching Pokemon:", err.Error())
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	caught := attemptCatch(pokemon)
	if caught {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Println("Congratulations! You caught", pokemon.Name)
	} else {
		fmt.Println("Catch attempt was unsuccessful.")
	}

	return nil
}

func attemptCatch(pokemon pokeapi.RespPokemon) bool {
	experience := pokemon.BaseExperience
	catchProbability := 1.0 - (float64(experience) / 1000.0)
	if catchProbability < 0.0 {
		catchProbability = 0.0
	} else if catchProbability > 1.0 {
		catchProbability = 1.0
	}

	randomValue := rand.Float64()
	return randomValue < catchProbability
}
