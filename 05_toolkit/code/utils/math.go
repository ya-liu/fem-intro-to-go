package utils

import "fmt"

// this function is not exported from  package utils
func printNum(num int) {
	fmt.Println("Current Number: ", num)
}

// Add function adds multiple numbers
// Capitalized functions in a package are visible and exported
func Add(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		printNum(v)
		total += v
	}
	return total
}
