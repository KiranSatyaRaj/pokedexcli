package utils

import (
	"fmt"
	"math/rand/v2"
	"encoding/json"
)

type Pokedex struct {
	Pokemons map[string]pokemon
}

type pokemon struct {
	name string
	height int
	weight int
	stats map[string]int
	types []string
}

var exp float64

var (
	maxNum int = 0
 	chance int = 0
	minAttempts int = 0
	catchAttempts int = 0
)
func CallPokeCatch(pokemon string) (bool, error) {
	url := "https://pokeapi.co/api/v2/pokemon"
	url = fmt.Sprintf("%s/%s", url, pokemon)
 
	resp, err := makeGetRequest(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	
	var m map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&m)
	
	baseExp := m["base_experience"].(float64)
	setParams(baseExp)
	return performCatch(baseExp), nil
}
func setParams(baseExp float64) {
	if maxNum == 0 && chance == 0 && minAttempts == 0 {
		maxNum = int(baseExp + 1)
		chance = int(baseExp)
		minAttempts = int(baseExp / 25)
	}
}

func unsetParams() {
	maxNum, chance, minAttempts, exp = 0, 0, 0, 0
}
 

func performCatch(baseExp float64) bool {
	exp = float64(rand.IntN(maxNum))

	if exp >= baseExp {
		//fmt.Printf("%v >= %v\n", exp, baseExp)
		unsetParams()
		return true
	}
	catchAttempts++

	if catchAttempts < minAttempts {
		//fmt.Printf("%v < %v\n", exp, baseExp)
		return false
	} else {
		maxNum += chance
		//chance += int(chance/2)
		//fmt.Printf("%v < %v\n", exp, baseExp)
		//fmt.Printf("%v %v\n", maxNum, chance)
		return false
	}
}
