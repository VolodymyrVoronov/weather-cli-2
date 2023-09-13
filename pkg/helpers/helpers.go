package helpers

import "strings"

func GetWeatherEmoji(weatherCondition string) string {

	if weatherCondition == "" {
		return "🌎"
	}

	weatherEmojis := map[string]string{
		"Sunny":                                    "☀️",
		"Partly cloudy":                            "⛅️",
		"Cloudy":                                   "☁️",
		"Overcast":                                 "☁️",
		"Mist":                                     "🌫️",
		"Patchy rain nearby":                       "🌦️",
		"Patchy snow nearby":                       "🌨️",
		"Patchy sleet nearby":                      "🌨️",
		"Patchy freezing drizzle nearby":           "🌨️",
		"Thundery outbreaks in nearby":             "⛈️",
		"Blowing snow":                             "🌨️ ❄️",
		"Blizzard":                                 "❄️ ☃️",
		"Fog":                                      "🌫️",
		"Freezing fog":                             "🌫️ ❄️",
		"Patchy light drizzle":                     "🌧️",
		"Light drizzle":                            "🌧️",
		"Freezing drizzle":                         "🌧️ ❄️",
		"Heavy freezing drizzle":                   "🌧️ ❄️",
		"Patchy light rain":                        "🌧️",
		"Light rain":                               "🌧️",
		"Moderate rain at times":                   "🌧️",
		"Moderate rain":                            "🌧️",
		"Heavy rain at times":                      "🌧️",
		"Heavy rain":                               "🌧️ ",
		"Light freezing rain":                      "🌧️ ❄️",
		"Moderate or heavy freezing rain":          "🌧️ ❄️",
		"Light sleet":                              "🌨️ ❄️",
		"Moderate or heavy sleet":                  "🌨️ ❄️",
		"Patchy light snow":                        "🌨️ ❄️",
		"Light snow":                               "🌨️ ❄️",
		"Patchy moderate snow":                     "🌨️ ❄️",
		"Moderate snow":                            "🌨️ ❄️",
		"Patchy heavy snow":                        "🌨️ ❄️",
		"Heavy snow":                               "🌨️ ❄️",
		"Ice pellets":                              "🌨️ ❄️",
		"Light rain shower":                        "🌧️",
		"Moderate or heavy rain shower":            "🌧️",
		"Torrential rain shower":                   "🌧️",
		"Light sleet showers":                      "🌨️ ❄️",
		"Moderate or heavy sleet showers":          "🌨️ ❄️",
		"Light snow showers":                       "🌨️ ❄️",
		"Moderate or heavy snow showers":           "🌨️ ❄️",
		"Light showers of ice pellets":             "🌨️ ❄️",
		"Moderate or heavy showers of ice pellets": "🌨️ ❄️",
		"Patchy light rain in area with thunder":   "🌩️ 🌧️",
		"Moderate or heavy rain in area with thunder": "🌩️ 🌧️",
		"Patchy light snow in area with thunder":      "🌩️ 🌨️ ❄️",
		"Moderate or heavy snow in area with thunder": "🌩️ 🌨️ ❄️",
	}

	emoji := weatherEmojis[weatherCondition]

	return emoji
}

func Reverse(s string, separator string) string {
	arr := strings.Split(s, separator)
	reversed := ""
	for i := len(arr) - 1; i >= 0; i-- {
		reversed += arr[i]
		if i != 0 {
			reversed += separator
		}
	}

	return reversed
}
