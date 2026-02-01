package results

type LocationArea struct {
	Name string `json: "name"`
}

type PokeApiResults struct {
	Results []LocationArea `json: "results"`
}
