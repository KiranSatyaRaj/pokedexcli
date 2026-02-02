package utils

import (
	"fmt"
	"io"
	"math/rand/v2"
	"encoding/json"
)

type PokeExp struct {
	BaseExp float64 `json:"base_experience"`
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
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var exp PokeExp
	if err = json.Unmarshal(data, &exp); err != nil {
		return false, nil
	}
	
	setParams(exp.BaseExp)
	
	isCaught := performCatch(exp.BaseExp)
	if isCaught{
		addEntry(data, pokemon)
	}
	return isCaught, nil
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
		return false
	} else {
		maxNum += chance
		return false
	}
}
