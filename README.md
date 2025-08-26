# Pokedex CLI

A command-line interface (CLI) tool built in Go that allows users to explore the Pokémon world using the PokeAPI. This interactive Pokedex provides a simple and fun way to browse locations, discover Pokémon, catch them, and inspect their details.

## Features

- **map**: Displays the next 20 location areas.
- **mapb**: Displays the previous 20 location areas.
- **explore**: Lists all Pokémon found in a specified location area.
- **catch**: Attempts to catch a Pokémon by its name and add it to your Pokedex.
- **inspect**: Shows detailed information about a caught Pokémon, including its name, height, weight, stats, and type(s).
- **pokedex**: Lists all Pokémon you have caught.
- **help**: Displays a help message with available commands.
- **exit**: Closes the Pokedex application.

## Usage

Run the program (./Pokedex) and enter commands at the `Pokedex >` prompt. For example:

- `map` to browse the next set of locations.
- `explore pastoria-city-area` to see Pokémon in that area.
- `catch pikachu` to attempt catching Pikachu.

The Pokedex uses a caching mechanism to optimize API calls, ensuring a smooth and efficient experience.