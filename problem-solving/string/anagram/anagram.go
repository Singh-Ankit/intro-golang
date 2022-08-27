// Q. Check Whether Two Strings Are Anagram of Each Other
package main

import (
	"fmt"
)

/*
	Approach 1: Approach1_Map Using Map/Dictionary
	Approach 2: Using Sorting
*/
func Approach1_Map(str1 string, str2 string) bool {
	//find the length of two strings and check if they are equal, if not equal in length simply return false and exit
	if len(str1) != len(str2) {
		return false
	}

	// create map
	isAnagram := make(map[string]int)

	//use the above map to keep track of the elements in the string
	for index, val := range str1 {
		isAnagram[string(str1[index])]++
		fmt.Println("val: ", val, " at Index: ", index)
	}

	for index, val := range str2 {
		isAnagram[string(str2[index])]--
		fmt.Println("val: ", val, " at index: ", index)
	}

	for key, val := range isAnagram {
		if val != 0 {
			fmt.Println("this is not repetitive: ", key)
			return false
		}
	}
	return true
}

/*func Approach2_Sort(str1 string, str2 string) bool {

	return false
}*/

func main() {
	var inpu_str1 string = "listen"
	var inpu_str2 string = "silent"

	var result bool = Approach1_Map(inpu_str1, inpu_str2)

	fmt.Println(result)
}
