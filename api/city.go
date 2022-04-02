package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Centre struct {
	Coordinates [2]float64 `json:"coordinates,omitempty"`
}

type City struct {
	Nom             string `json:"nom,omitempty"`
	CodeDepartement string `json:"codeDepartement,omitempty"`
	Population      int    `json:"population,omitempty"`
	Centre          Centre `json:"centre,omitempty"`
}

func FindCity(cityName string) City {
	body := Fetch(
		fmt.Sprintf(
			"https://geo.api.gouv.fr/communes?nom=%v&fields=nom,codesPostaux,codeDepartement,population,centre&format=json&geometry=centre",
			url.QueryEscape(cityName),
		),
	)

	var cities []City
	if err := json.Unmarshal(body, &cities); err != nil {
		panic(err)
	}

	return cities[0]
}
