package utils 

import (
	"fmt"
)

func Inspect(pokemon string) {
	pokeDetails, ok := pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}
	fmt.Printf("Name: %v\n", pokeDetails.name)
	fmt.Printf("Height: %v\n", pokeDetails.height)
	fmt.Printf("Weight: %v\n", pokeDetails.weight)
	fmt.Println("Stats:")
	for k, v := range pokeDetails.stats {
		fmt.Printf("  -%v: %v\n", k, v)
	}
	fmt.Println("Types:")
	for _, v := range pokeDetails.types {
		fmt.Printf("  - %v\n", v)
	}
}
