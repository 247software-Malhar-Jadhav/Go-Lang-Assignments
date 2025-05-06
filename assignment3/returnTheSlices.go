package main

import (
	"fmt"
	"os"
)

func ReturnTheSlices() {
	// Hardcoded array
	arr := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	// Read input indices
	var index1, index2 int
	fmt.Print("Enter two space-separated indices: ")
	fmt.Scanln(&index1, &index2)

	// Check for index bounds
	if index1 < 0 || index2 < 0 || index1 > len(arr) || index2 > len(arr) {
		fmt.Println("Incorrect Indexes")
		os.Exit(0)
	}

	// Create and print slices
	slice1 := arr[:index1+1]
	slice2 := arr[index1 : index2+1]
	slice3 := arr[index2:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
