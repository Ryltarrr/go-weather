package cmd

import (
	"fmt"
	"go-weather/api"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-weather",
	Short: "Simple weather CLI app",
	Long: `go-weather is a simple CLI tool to get the weather of a french city`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := api.FindCity(args[0])
		weather := api.GetWeather(city.Centre.Coordinates)

		fmt.Printf("Actually in %v, the weather is %v and temperature is %v°C (min: %v°C, max: %v°C)",
			city.Nom, weather.Weather[0].Description, weather.Main.Temp, weather.Main.TempMin, weather.Main.TempMax)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-weather.yaml)")
}
