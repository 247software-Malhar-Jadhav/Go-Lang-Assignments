package main

import (
	"fmt"
	"strings"
)

func WordCount() {
	// Take input string
	var input string
	fmt.Println("Enter a sentence:")
	getline := func() string {
		var line string
		for {
			var temp string
			n, _ := fmt.Scanln(&temp)
			if n == 0 {
				break
			}
			if line != "" {
				line += " "
			}
			line += temp
			if strings.HasSuffix(line, ".") {
				break
			}
		}
		return line
	}
	input = getline()

	// Split input into words
	words := strings.Fields(input)

	// Count frequency of each word
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}

	// Find the maximum frequency
	maxFreq := 0
	for _, count := range freqMap {
		if count > maxFreq {
			maxFreq = count
		}
	}

	// Collect words with highest frequency in order of appearance
	seen := make(map[string]bool)
	result := []string{}
	for _, word := range words {
		if freqMap[word] == maxFreq && !seen[word] {
			result = append(result, word)
			seen[word] = true
		}
	}

	fmt.Println(result)
}
