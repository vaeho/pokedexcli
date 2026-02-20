package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location to explore")
		return nil
	}
	location := args[0]
	fmt.Println("Exploring", location)
	locationPokemons, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		fmt.Printf("Error exploring location %s: %v\n", location, err)
		return err
	}

	pokemonEncounters := locationPokemons.PokemonEncounters
	if len(pokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in location %s\n", location)
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}

	return nil
}
