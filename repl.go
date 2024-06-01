package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Lists the next page of location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the prevoius page of location areas",
			callback: commandMapB,
		},
	}
}

func getInput(prompt string, r* bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input)), err
}

func startREPL(cfg *config) {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmdName, _ := getInput("pokedex > ", reader)
		cmd, exists := getCommands()[cmdName]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
	
}