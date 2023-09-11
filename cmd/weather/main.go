package main

import (
	"bufio"
	"fmt"
	"os"
	"weather-cli/models"
	"weather-cli/pkg/api"

	"github.com/fatih/color"
)

func main() {
	var cityName string
	var forecast models.ForecastDay
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	color.Cyan("Enter the name of the city. Example: New-York")
	color.Yellow("Note: If your city name contains spaces, please use dash.")
	color.Magenta("----------------------------------")
	color.Cyan("Your city:")

	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	cityName = scanner.Text()

	forecast, err = api.GetWeather(cityName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(forecast)

}
