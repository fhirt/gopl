package main

import (
	"fmt"
	"gopl/ch2/echo"
	"gopl/ch2/temperature"
)

func main() {
	echo.Echo()
	temperature.Boiling()
	const freezingF, boilingF = temperature.Fahrenheit(32.0), temperature.Fahrenheit(212.0)
	fmt.Printf("%s = %s\n", freezingF, temperature.FToC(freezingF))
	fmt.Printf("%s = %s\n", boilingF, temperature.FToC(boilingF))
	fmt.Printf("%s = %s\n", temperature.BoilingC, temperature.CToK(temperature.BoilingC))
	fmt.Printf("%s = %s\n", temperature.AbsoluteZeroC, temperature.CToK(temperature.AbsoluteZeroC))
	fmt.Printf("%s = %s\n", temperature.FreezingC, temperature.CToK(temperature.FreezingC))
}
