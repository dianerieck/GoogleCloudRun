package utils

import "math"

func CelsiusToFahrenheit(celsius float64) float64 {
	return math.Round((celsius*1.8)+32) * 100 / 100
}

func CelsiusToKelvin(celsius float64) float64 {
	return math.Round((celsius+273)*100) / 100
}
