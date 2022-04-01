package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	res, err := http.Get(fmt.Sprintf("https://geo.api.gouv.fr/communes?nom=%v&fields=nom,codesPostaux,codeDepartement,population,centre&format=json&geometry=centre", cityName))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var cities []City
	if err := json.Unmarshal(body, &cities); err != nil {
		panic(err)
	}

	return cities[0]
}
