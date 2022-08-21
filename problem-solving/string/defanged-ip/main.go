package main

import (
	"fmt"
)

func DefangedIp(str string) string {
	result := ""
	for _, i := range str {
		if string(i) == "." {
			result += "[.]"
		} else {
			result += string(i)
		}
	}
	return result
}

func main() {
	var input string = "255.100.50.0"
	var result = DefangedIp(input)
	fmt.Println(result)
}
