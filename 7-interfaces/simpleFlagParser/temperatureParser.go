package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c *Celsius) String() string {
	return fmt.Sprintf("%f°C", *c)
}

func fahrenheitToCelsius(t float64) Celsius {
	return Celsius((t - 32) * 5 / 9)
}

func kelvinToCelsius(t float64) Celsius {
	return Celsius(t - 273.15)
}

type CelsiusFlag struct{ Celsius }

func (c *CelsiusFlag) Set(input string) error {
	var unit string
	var temp float64
	_, err := fmt.Sscanf(input, "%f%s", &temp, &unit)
	if err != nil {
		fmt.Println("set error", err)
	}
	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(temp)
		return nil
	case "F", "°F":
		c.Celsius = fahrenheitToCelsius(temp)
		return nil
	case "K", "°K":
		c.Celsius = kelvinToCelsius(temp)
		return nil
	}
	return fmt.Errorf("No unit found")
}

func ParseTemperatureFlag(c Celsius) *Celsius {
	cf := CelsiusFlag{c}
	flag.CommandLine.Var(&cf, "temp", "input temperature")
	return &cf.Celsius
}

var c *Celsius

func init() {
	c = ParseTemperatureFlag(20)
}

func main() {
	flag.Parse()
	fmt.Printf("Temperature in Celsius: %s \n", c)
}
