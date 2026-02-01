package repl

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/cmd"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/args"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func TakeInput() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd.Cmds["help"].Callback()
	for {
		fmt.Print("Pokedex >")
		var input string
		if scanner.Scan(){
			input = scanner.Text()
			cleanText := cleanInput(input)
			validCmd, isValid := cmd.Cmds[cleanText[0]]
			if isValid {
				if len(cleanText) > 1 {
					args.CreateArgs(cleanText[1:])
				}
				validCmd.Callback()
			} else {
				fmt.Println("Unknown Command")
			}
		}
	}
}


