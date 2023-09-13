package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"weather-cli/models"
	"weather-cli/pkg/api"
	"weather-cli/pkg/helpers"

	"github.com/fatih/color"
)

func main() {
	var cityName string
	var forecast models.ForecastDay
	var err error

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		color.Cyan("Enter the name of the city. Example: New-York")
		color.Yellow("Note: If your city name contains spaces, please use dash.")
		color.Magenta("----------------------------------\n\n")
		color.Cyan("Your city:")

		scanner.Scan()

		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		cityName = scanner.Text()

		forecast, err = api.GetWeather(cityName)

		if err != nil {
			fmt.Println()
			color.Red("No matching location found.\n\n")
			color.Magenta("----------------------------------")

			if !strings.Contains(err.Error(), "error") {
				break
			}
		} else {
			country := forecast.Location.Country
			city := forecast.Location.Name
			date := helpers.Reverse(strings.Split(forecast.Location.Localtime, " ")[0], "-")
			time := strings.Split(forecast.Location.Localtime, " ")[1]
			temp := forecast.Current.TempC
			tempFeelsLike := forecast.Current.FeelslikeC
			condition := forecast.Current.Condition.Text
			wind := forecast.Current.WindKph
			windDirection := forecast.Current.WindDir
			humidity := forecast.Current.Humidity
			cloud := forecast.Current.Cloud
			visibility := forecast.Current.VisKm

			fmt.Println()
			fmt.Fprintf(w, "┌────────────────┬──────────────────────────────────┐\n")
			fmt.Fprintf(w, "│ Country        │ %s\n", country)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ City           │ %s\n", city)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Date           │ %s\n", date)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Time           │ %s\n", time)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Temperature    │ %.0f °C\n", temp)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Feels like     │ %.0f °C\n", tempFeelsLike)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Condition      │ %s - %s\n", condition, helpers.GetWeatherEmoji(condition))
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Wind           │ %.1f km/h\n", wind)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Wind direction │ %s\n", windDirection)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Humidity       │ %d %%\n", humidity)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Cloud          │ %d %%\n", cloud)
			fmt.Fprintf(w, "├────────────────┼──────────────────────────────────┤\n")
			fmt.Fprintf(w, "│ Visibility     │ %.0f km\n", visibility)
			fmt.Fprintf(w, "└────────────────┴──────────────────────────────────┘\n")

			w.Flush()

			break
		}
	}
}
