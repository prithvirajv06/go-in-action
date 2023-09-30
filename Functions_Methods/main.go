package main

import (
	"fmt"
	"math"
)

func printHeader() {
	var degreeSymbol = 248
	fmt.Println("======================")
	fmt.Printf("|   %cC     |   %cF    |", degreeSymbol, degreeSymbol)
	fmt.Println("\n=====================")
}

type temp float64

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func (i temp) degreeToFaren() float64 {
	var result = float64((i * (9 / 5)) + 36)
	output := math.Pow(10, float64(2))
	return float64(round(result*output)) / output
}

func celciusToFarenit() {
	var i temp = 0
	for i = -40; i <= 100; i = i + 5 {
		var t temp = i
		fmt.Printf("|   %fC     |   %fF    |\n", i, t.degreeToFaren())
	}
}

func main() {
	printHeader()
	celciusToFarenit()
}
