package repl

import (
	"github.com/hachimB/Pokedex/internal/pokecache"
	"github.com/hachimB/Pokedex/internal/api"
)

type config struct {
	Next *string
	Previous *string
	cache *pokecache.Cache
	pokedex map[string]*api.Pokemon
}



type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error
}


func registerCommands() map[string]cliCommand{
	return map[string]cliCommand {
		"exit" : {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help" : {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map" : {
			name: "map",
			description: "Displays the names of 20 next location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb" : {
			name: "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world",
			callback: commandMapb,
		},
		"explore" : {
			name: "explore",
			description: "Displays the list of all the Pokémon located in a place",
			callback: commandExplore,
		},
		"catch" : {
			name: "catch",
			description: "Allows to catch a Pokemon by it's name",
			callback: commandCatch,
		},
		"inspect" : {
			name: "inspect",
			description: "Prints the name, height, weight, stats and type(s) of the Pokemon",
			callback: commandInspect,
		},
		"pokedex" : {
			name: "pokedex",
			description: "Lists all the names of the pokemon the user has caught",
			callback: commandPokedex,
		},
	}
}