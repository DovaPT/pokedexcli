package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationData struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}`json:"results"`
}

func getLocations(path *string) (LocationData, error){
	jsonData, err := queryApi(*path)
	if err != nil {
		return LocationData{}, err
	}
	var locations LocationData
	json.Unmarshal(jsonData, &locations)

	return locations, nil
}


func queryApi(url string) ([]byte, error){
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
