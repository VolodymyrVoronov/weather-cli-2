package main

import (
	"fmt"
	"weather-cli/pkg/api"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Prints text in cyan.")

	forecast, err := api.GetWeather("sasasrk")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(forecast)
}
