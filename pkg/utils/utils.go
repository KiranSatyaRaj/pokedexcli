package utils

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
	"time"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/pokecache"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/results"
)

const interval = time.Duration(1 * time.Minute)

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

var offset = -20

var cache pokecache.Cache

func makeGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return &http.Response{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	
	return resp, nil
}
func CallPokeApi(cmd string) ([]results.LocationArea, error) {
	if offset == 0 && cmd == "mapb" {
		return []results.LocationArea{}, fmt.Errorf("Zero Location Areas cannot map back!")
	}
	
	if cache.IsEmpty() {
		cache = pokecache.NewCache(interval)
	}

	if cmd == "map" {
		offset += 20
	}

	if cmd == "mapb" {
		offset -= 20
	}

	url := fmt.Sprintf("%s?offset=%d", baseUrl, offset) 
	resp, err := makeGetRequest(url)
	if err != nil {
		return []results.LocationArea{}, err
	}

	defer resp.Body.Close()	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []results.LocationArea{}, err
	}
	
	var par results.PokeApiResults
	
	if err = json.Unmarshal(data, &par); err != nil {
		return []results.LocationArea{}, err
	}

	
	if results, ok := cache.Get(url); ok {
		return results, nil
	} else {
		cache.Add(url, par.Results)
	}

	return par.Results, nil
}



func CallLocationAreaPoke(locationArea string) ([]string, error) {
	url := fmt.Sprintf("%s/%s", baseUrl, locationArea)
	resp, err := makeGetRequest(url)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	pokemons, err := parseLocationAreaPokemons(resp)
	
	if err != nil {
		return []string{}, nil
	}

	return pokemons, nil
}

func parseLocationAreaPokemons(resp *http.Response) ([]string, error) {
	
	dec := json.NewDecoder(resp.Body)
	
	var m map[string]interface{}
	err := dec.Decode(&m)
	
	if err != nil {
		return []string{}, err
	}
	pokemons := []string{}
	vals := m["pokemon_encounters"].([]interface{})

	for _, val := range vals {
		pokemon := val.(map[string]interface{})["pokemon"].(map[string]interface{})["name"]
		pokemons = append(pokemons, pokemon.(string))
	}
	return pokemons, nil
}
