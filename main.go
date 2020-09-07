package main

import (
	"github.com/fr3fou/margov/margov"
)

const (
	Sunny  = "Sunny"
	Rainy  = "Rainy"
	Cloudy = "Cloudy"
)

func main() {
	chain := margov.New()

	// Given that the current day is Sunny.
	chain.Set(Sunny, Sunny, 0.8)
	chain.Set(Rainy, Sunny, 0.05)
	chain.Set(Cloudy, Sunny, 0.15)

	// Given that the current day is Rainy.
	chain.Set(Sunny, Rainy, 0.2)
	chain.Set(Rainy, Rainy, 0.6)
	chain.Set(Cloudy, Rainy, 0.2)

	// Given that the current day is Rainy.
	chain.Set(Sunny, Cloudy, 0.2)
	chain.Set(Rainy, Cloudy, 0.3)
	chain.Set(Cloudy, Cloudy, 0.5)
}
