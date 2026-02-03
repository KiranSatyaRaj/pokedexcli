package utils

import (
	"fmt"
)

var	pokedex map[string]*pokemon

type pokemon struct {
	name string
	height int
	weight int
	stats map[string]int
	types []string
}

var pokeObject []byte
var pokeName string

func addEntry(data []byte, pokename string) error {
	if pokedex == nil {
		pokedex = make(map[string]*pokemon)
	}
	pokeObject, pokeName = data, pokename
	pokedex[pokeName] = &pokemon{}
	setName()
	err1, err2, err3 := setHeightAndWeight(), setStats(), setTypes()
	if err1 != nil {
		return err1
	} 
	if err2 != nil {
		return err2 
	}
	if err3 != nil {
		return err3
	}
	return nil
}

func setName() {
	pokedex[pokeName].name = pokeName
}

func setHeightAndWeight() error {
	height, weight, err := getHeightAndWeight()
	if err != nil {
		return err
	}
	pokedex[pokeName].height, pokedex[pokeName].weight = height, weight
	return nil
}

func setStats() error {
	var err error
	pokedex[pokeName].stats, err = getStats()
	if err != nil {
		return err
	}
	return nil
}

func setTypes() error {
	var err error
	pokedex[pokeName].types, err = getTypes()
	if err != nil {
		return err
	}
	return nil
}

func ShowPokedex() error {
	if pokedex == nil {
		return fmt.Errorf("You have not caught any pokemons.")
	}
	fmt.Println("Your Pokedex:")
	for pokemon, _ := range pokedex {
		fmt.Printf("  - %s\n", pokemon)
	}
	return nil
}

