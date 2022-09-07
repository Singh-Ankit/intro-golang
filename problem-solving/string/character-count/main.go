package main

import (
	"fmt"
)

func MaxOccurrenceCharacter(input_str string) (count int, char string) {
	var maxCount int = 0
	var maxCountChar string
	//create a map
	charCount := make(map[string]int)

	//iterate over the input string
	for index, _ := range input_str {
		//fmt.Println(index, string(value))
		charCount[string(input_str[index])]++
	}

	for key, val := range charCount {
		if val > maxCount {
			maxCount = val
			maxCountChar = key
		}
	}

	return maxCount, maxCountChar
}

func main() {
	fmt.Println("print the character with max count in a string")
	var input_str string = "ajbhghsffsfff"
	count, char := MaxOccurrenceCharacter(input_str)

	fmt.Println(count, char)
}
