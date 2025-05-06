package main

import (
	"fmt"
)

func getValue(symbol string) int {
	switch symbol {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func main() {
	roman := "MCMXCIV"
	total := 0
	i := 0

	for i < len(roman) {
		// Current symbol and its value
		current := string(roman[i])
		currentValue := getValue(current)

		// Check next symbol
		if i+1 < len(roman) {
			next := string(roman[i+1])
			nextValue := getValue(next)

			// Subtraction case
			if currentValue < nextValue {
				total += nextValue - currentValue
				i += 2
				continue
			}
		}

		// Normal addition case
		total += currentValue
		i++
	}

	fmt.Println("The number is:", total)
}
