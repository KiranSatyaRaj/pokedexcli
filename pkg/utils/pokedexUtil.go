package utils

import (
	"encoding/json"
)

func unmarshaller[T any](obj *T) error {
	if err := json.Unmarshal(pokeObject, obj); err != nil {
		return err
	}
	return nil
}

func getHeightAndWeight() (int, int, error) {
	pokeHtWt := struct {
		Height int `json:"height"`
		Weight int `json:"weight"`
	}{}
	if err := unmarshaller(&pokeHtWt); err != nil {
		return 0,0, err
	}
	return pokeHtWt.Height, pokeHtWt.Weight, nil
}

type Name struct {
	Name string `json:"name"`
}

type Type struct {
	PokeType Name `json:"type"`
}

type PokeTypes struct {
	Res []Type `json:"types"`
}

func getTypes() ([]string, error) {
	var pokeTypes PokeTypes
	if err := unmarshaller(&pokeTypes); err != nil {
		return []string{}, err
	}
	pokeTypeArr := []string{}
	for _, v := range pokeTypes.Res {
		pokeTypeArr = append(pokeTypeArr, v.PokeType.Name)
	}
	return pokeTypeArr, nil
}

type stat struct {
	statType string
	value 	 int
}

type StatTypes struct {
	StatType Name `json:"stat"`
	BaseStat int `json:"base_stat"`
}

type AllStats struct {
	Stats []StatTypes `json:"stats"`
}

type allStats struct {
	statType string
	statScore int
}

func getStats() (map[string]int, error) {
	var stats AllStats
	if err := unmarshaller(&stats); err != nil {
		return make(map[string]int), err
	}
	statMap := make(map[string]int)
	for _, v := range stats.Stats {
		statMap[v.StatType.Name] = v.BaseStat
	}
	return statMap, nil
}
