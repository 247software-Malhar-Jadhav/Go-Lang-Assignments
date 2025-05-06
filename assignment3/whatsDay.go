package main

import (
	"fmt"
)

func WhatsDay() {
	// Create a map with index as key and day as value
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var index int
	fmt.Print("Enter the index (1-7): ")
	fmt.Scanln(&index)

	// Check if the index exists in the map
	if day, exists := days[index]; exists {
		fmt.Println(day)
	} else {
		fmt.Println("Not a day")
	}
}
