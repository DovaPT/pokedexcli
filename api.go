package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	Name             string `json:"name"`
	PokemonEncounter []struct {
		Pokemon struct {
			Name string `json:"name"`
		}`json:"pokemon"`
	}`json:"pokemon_encounters"`
}

func getLocations(path *string) (LocationData, error) {
	jsonData, err := queryApi(*path)
	if err != nil {
		return LocationData{}, err
	}
	var locations LocationData
	json.Unmarshal(jsonData, &locations)

	return locations, nil
}

func getLocationInfo(location *string) (LocationArea, error) {
	jsonData, err := queryApi("https://pokeapi.co/api/v2/location-area/" + *location)
	if err != nil {
		return LocationArea{}, err
	}
	var info LocationArea
	json.Unmarshal(jsonData, &info)
	
	return info, nil
}

func queryApi(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return jsonData, nil
}
