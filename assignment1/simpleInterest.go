package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SimpleInterest() {
	var input string
	fmt.Println("Enter Principal, Rate, Time (comma-separated):")
	fmt.Scanln(&input)

	values := strings.Split(input, ",")
	if len(values) != 3 {
		fmt.Println("Please enter exactly 3 comma-separated values.")
		return
	}

	principal, _ := strconv.ParseFloat(strings.TrimSpace(values[0]), 64)
	rate, _ := strconv.ParseFloat(strings.TrimSpace(values[1]), 64)
	time, _ := strconv.ParseFloat(strings.TrimSpace(values[2]), 64)

	simpleInterest := (principal * rate * time) / 100
	fmt.Printf("Simple Interest: %.2f\n", simpleInterest)
}
