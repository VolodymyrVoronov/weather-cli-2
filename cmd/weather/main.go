package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weather-cli/models"
	"weather-cli/pkg/api"

	"github.com/fatih/color"
)

func main() {
	var cityName string
	var forecast models.ForecastDay
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	for {
		color.Cyan("Enter the name of the city. Example: New-York")
		color.Yellow("Note: If your city name contains spaces, please use dash.")
		color.Magenta("----------------------------------")
		fmt.Println()
		color.Cyan("Your city:")

		scanner.Scan()

		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		cityName = scanner.Text()

		forecast, err = api.GetWeather(cityName)

		if err != nil {
			color.Red("No matching location found.")
			fmt.Println()
			color.Magenta("----------------------------------")

			if !strings.Contains(err.Error(), "error") {
				break
			}
		} else {
			fmt.Println(forecast)
			break
		}
	}

}
