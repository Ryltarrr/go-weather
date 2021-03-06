package api

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type WeatherMain struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  float64 `json:"humidity"`
}

type WeatherResponse struct {
	Weather []Weather   `json:"weather"`
	Main    WeatherMain `json:"main"`
}

func (wr *WeatherResponse) roundTemperatures() {
	wr.Main.Temp = math.Round(wr.Main.Temp)
	wr.Main.TempMin = math.Round(wr.Main.TempMin)
	wr.Main.TempMax = math.Round(wr.Main.TempMax)
	wr.Main.FeelsLike = math.Round(wr.Main.FeelsLike)
}

func GetWeather(coordinates [2]float64) WeatherResponse {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")

	body := Fetch(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lon=%v&lat=%v&appid=%v&units=metric", coordinates[0], coordinates[1], apiKey))

	return decode(body)
}

func decode(body []byte) WeatherResponse {
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		panic(err)
	}

	weather.roundTemperatures()
	return weather
}
