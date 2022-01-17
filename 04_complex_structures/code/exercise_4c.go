package main

import "fmt"

// Part 1
func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

// Part 2
func found(name string) bool {
	pets := make(map[string]string)
	pets["Sprinkles"] = "cat"
	pets["Chichi"] = "dog"
	name, ok := pets[name]
	return ok
}

func main() {
	fmt.Println(average(3, 45, 666))
	fmt.Println(found("Sprinkles"))
}
