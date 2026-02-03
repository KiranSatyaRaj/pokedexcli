package cmd

import (
	"os"
	"fmt"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/utils"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/results"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/args"
)

var Cmds = map[string]cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		Callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		Callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays 20 location areas",
		Callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays previous 20 location areas",
		Callback: commandMapBack,
	},
	"explore": {
		name: "explore",
		description: "Displays pokemons in that location area",
		Callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Catches a pokemon and adds it to your Pokedex.",
		Callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "Inspects a Pokemon in your Pokedex.",
		Callback: commandInspect,
	},
	"pokedex": {
		name: "pokedex",
		description: "Displays your pokedex",
		Callback: commandPokedex,
	},
}

type cliCommand struct {
	name string
	description string
	Callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

func commandMap() error {
	locations, err := utils.CallPokeApi("map")
	if err != nil {
		return err
	}
	printLocations(locations)
	return nil
}

func commandMapBack() error {
	locations, err := utils.CallPokeApi("mapb")
	if err != nil {
		return err
	}
	printLocations(locations)
	return nil
}

func commandExplore() error {
	pokemons, err := utils.CallLocationAreaPoke(args.Args[0])
	if err != nil {
		return err
	}
	printPokemons(pokemons)
	return nil
}

func commandCatch() error {
	pokemon := args.Args[0]
	isCaught, err := utils.CallPokeCatch(pokemon)
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon) 
	if err != nil {
		return err
	}

	if isCaught {
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Println("You may now inspect it with inspect command")
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
}

func commandInspect() error {
	utils.Inspect(args.Args[0])
	return nil
}
func printPokemons(pokemons []string) {
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
}

func printLocations(locations []results.LocationArea) {
	for _, location := range locations { 
		fmt.Printf("%v\n", location.Name)
	}
}

func commandPokedex() error {
	if err := utils.ShowPokedex(); err != nil {
		return err
	}
	return nil
}
