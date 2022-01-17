package main

import "fmt"

// Part 1
func average(scores [5]float64) float64 {
	total := 0.0
	for _, number := range scores {
		total += number
	}
	return total / float64(len(scores))
}

// Part 2
func found(name string) bool {
	pets := make(map[string]string)
	pets["Sprinkles"] = "cat"
	pets["Chichi"] = "dog"
	name, ok := pets[name]
	return ok
}

// Part 3
var groceries = []string{"yogurt", "apples", "carrots", "hummus", "pita"}

func addGroceries(item ...string) {
	newGroceries := append(groceries, item...)
	fmt.Println(newGroceries)
	fmt.Println(len(newGroceries))
	fmt.Println(cap(newGroceries))
}

func main() {
	scores := [5]float64{3, 45, 666, 8, 99}
	fmt.Println(average(scores))
	fmt.Println(found("Sprinkles"))
	addGroceries("oranges", "milk", "sausages")
}
