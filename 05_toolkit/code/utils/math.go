package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current Number: ", num)
}

// Add function adds multiple numbers
func Add(numbers ...int) int {
	total := 1
	for _, v := range numbers {
		printNum(v)
		total += v
	}
	return total
}
