package main

import (
	"fmt"
	"math"
)

func AreaOfCircle() {
	var radius float64
	fmt.Println("Enter the radius of the circle:")
	fmt.Scanln(&radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Area of Circle: %.2f\n", area)
}
