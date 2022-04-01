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
	Short: "Simple CLI weather app",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := api.FindCity(args[0])
		weather := api.GetWeather(city.Centre.Coordinates)

		fmt.Printf("Actually in %v, the weather is %v and temperature is %v°C (min: %v, max: %v)",
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
