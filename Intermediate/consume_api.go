package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Response ...
type Response struct {
	Description  []Description          `json:"description"`
	ID           int                    `json:"id"`
	IsMainSeries bool                   `json:"is_main_series"`
	Name         string                 `json:"name"`
	Names        []Names                `json:"names"`
	PokeEntries  []PokeEntries          `json:"pokemon_entries"`
	Regions      map[string]interface{} `json:"region"`
	VersionGrp   map[string]interface{} `json:"version_groups"`
}

//Description ...
type Description struct {
	Desc     string                 `json:"description"`
	Language map[string]interface{} `json:"language"`
}

//Names ...
type Names struct {
	Language map[string]interface{} `json:"language"`
	Name     string                 `json:"name"`
}

//PokeEntries ...
type PokeEntries struct {
	EntryNumber int                    `json:"entry_number"`
	Species     map[string]interface{} `json:"pokemon_species"`
}

func main() {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		panic(err)
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(resData))
	var responseData Response
	json.Unmarshal(resData, &responseData)
	// fmt.Println(responseData, len(responseData.PokeEntries))

	for i, value := range responseData.PokeEntries {
		fmt.Println(i, value)
	}

}
