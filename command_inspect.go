package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Pokemon to inspect")
		return nil
	}
	pokemonName := args[0]
	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Printf("You haven't caught %s yet. Try catching it first!\n", pokemonName)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("  - %s\n", ability.Ability.Name)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
	return nil
}
