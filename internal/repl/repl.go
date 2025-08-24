package repl

import (
	"fmt"
	"bufio"
	"time"
	"os"
	"github.com/hachimB/Pokedex/internal/pokecache"
	"github.com/hachimB/Pokedex/internal/api"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	commandsMap := registerCommands()

	// conf := new(config)
	// conf.cache = pokecache.NewCache(5 * time.Minute)
	// conf.pokedex = make(map[string]*api.Pokemon)

	conf := &config{
		cache : pokecache.NewCache(5 * time.Minute),
		pokedex : make(map[string]*api.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input := scanner.Text()
			words := CleanInput(input)
			if len(words) == 0 {
				break
			} else {
				b := false
				for key := range commandsMap {
					if words[0] == commandsMap[key].name {
						b = true
						commandsMap[key].callback(conf, words[1:])
					}
				}
				if b == false {
					fmt.Println("Unknown command")
					break
				}
			}
		}
	}
}