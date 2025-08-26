package repl

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/hachimB/Pokedex/internal/api"
	"math/rand"
)

func CleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}

func commandExit(p *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(p *config, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key := range registerCommands() {
		fmt.Printf("%s: %s\n", registerCommands()[key].name, registerCommands()[key].description)
	}
	return nil
}

func commandMap(p *config, args []string) error {
	var url string
	cache := p.cache
	if p.Next != nil {
		url = *(p.Next)
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	var parser api.Parse
	if entry, ok := cache.Get(url); ok {
		err := json.Unmarshal(entry, &parser)
		if err != nil {
			return err
		}
		for _, value := range parser.Results{
			fmt.Println(value.Name)
		}
		p.Next = parser.Next
		p.Previous = parser.Previous
		return nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&parser)
		if err != nil {
			return err
		}
		p.Next = parser.Next
		p.Previous = parser.Previous
		for _, val := range parser.Results {
			fmt.Println(val.Name)
		}

		ps, err := json.Marshal(parser)
		if err != nil {
			return err
		}
		cache.Add(url, ps)
	}
	return nil
}


func commandMapb(p *config, args []string) error {
	var url string
	cache := p.cache
	if p.Previous != nil {
		url = *(p.Previous)
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	var parser api.Parse
	if entry, ok := cache.Get(url); ok {
		err := json.Unmarshal(entry, &parser)
		if err != nil {
			return err
		}
		for _, value := range parser.Results{
			fmt.Println(value.Name)
		}
		p.Next = parser.Next
		p.Previous = parser.Previous
		return nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&parser)
		if err != nil {
			return err
		}
		p.Next = parser.Next
		p.Previous = parser.Previous
		for _, val := range parser.Results {
			fmt.Println(val.Name)
		}

		ps, err := json.Marshal(parser)
		if err != nil {
			return err
		}
		cache.Add(url, ps)
	}
	return nil
}

func commandExplore(p *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please add location's name")
		return fmt.Errorf("Please add location's name or ID")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	var locationName api.GetWithLocationName
	if entry, ok := p.cache.Get(url); ok {
		err := json.Unmarshal(entry, &locationName)
		if err != nil {
			return err
		}
		fmt.Printf("Exploring %s...\n", args[0])
		fmt.Println("Found Pokemon:")
		for _, val := range locationName.PokemonEncounters {
			println(val.Pokemon.Name)
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusOK {
			fmt.Printf("Exploring %s...\n", args[0])
			fmt.Println("Found Pokemon:")
		} else {
			fmt.Println("Not found: Make sure You are using the correct name")
		}
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&locationName)
		if err != nil {
			return err
		}

		for _, val := range locationName.PokemonEncounters {
			fmt.Println(val.Pokemon.Name)
		}
		data, err := json.Marshal(locationName)
		if err != nil {
			return err
		}
		p.cache.Add(url, data)
	}
	return nil
}

func commandCatch(p *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please add pokemon's name")
		return fmt.Errorf("Please add pokemon's name or ID")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	
	var pokemon *api.Pokemon

	if entry, ok := p.cache.Get(url); ok {
		err := json.Unmarshal(entry, &pokemon)
		if err != nil {
			return err
		}
		c := rand.Intn(pokemon.BaseExperience)
		fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
		if c < pokemon.BaseExperience / 4 {
			if p.pokedex[pokemon.Name] != nil {
				fmt.Printf("%s has already been caught!\n", args[0])
				return fmt.Errorf("Pokemon already cautch")
			}
			fmt.Printf("%s was caught!\n", args[0])
			p.pokedex[pokemon.Name] = pokemon
		} else {
			if p.pokedex[pokemon.Name] != nil {
				fmt.Printf("%s has already been caught!\n", args[0])
				return fmt.Errorf("Pokemon already cautch")
			}
			fmt.Printf("%s escaped!\n", args[0])
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusOK {
			fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
		} else {
			fmt.Println("Not found: Make sure You are using the correct name")
		}
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&pokemon)
		if err != nil {
			return err
		}

		c := rand.Intn(pokemon.BaseExperience)

		if c < pokemon.BaseExperience / 4 {
			fmt.Printf("%s was caught!\n", args[0])
			p.pokedex[pokemon.Name] = pokemon
		} else {
			fmt.Printf("%s escaped!\n", args[0])
		}

		data, err := json.Marshal(pokemon)
		if err != nil {
			return err
		}
		p.cache.Add(url, data)
	}
	return nil
}

func commandInspect(p *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please add pokemon's name")
		return fmt.Errorf("Please add pokemon's name or ID")
	}
	pokemon := p.pokedex
	if pokemon[args[0]] == nil {
		fmt.Println("you have not caught that pokemon")
		return fmt.Errorf("Pokemon not caught")
	} else {
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon[args[0]].Name, pokemon[args[0]].Height, pokemon[args[0]].Weight)
	fmt.Print("Stats:")
	for _, s := range(pokemon[args[0]].Stats) {
		fmt.Printf("\n  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Print("Types:")
	for _, t := range(pokemon[args[0]].Types) {
		fmt.Printf("\n  - %s\n", t.Type.Name)
	}
	}
	return nil
}

func commandPokedex(p *config, args []string) error {
	if len(args) != 0 {
		fmt.Println("Pokedex command does not take any argument")
		fmt.Errorf("Pokedex command does not take any argument")
	}
	pokemon := p.pokedex
	if len(pokemon) == 0 {
		fmt.Println("Your Pokedex is empty.")
	} else {
		fmt.Println("Your Pokedex:")
		for name := range pokemon {
			fmt.Printf("- %s\n", name)
		}
	}
	return nil
}
