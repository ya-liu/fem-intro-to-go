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
var initialPets map[string]string = map[string]string{
	"Sprinkles": "cat",
	"Chichi":    "dog",
	"Fido":      "dog",
}

func found(name string) bool {
	_, ok := initialPets[name]
	return ok
}

// Part 3
var groceries = []string{"yogurt", "apples", "carrots", "hummus", "pita"}

func addGroceries(newGroceries ...string) []string {
	foods := groceries
	// foods := append(groceries, newGroceries...)
	for _, g := range newGroceries {
		foods = append(foods, g)
	}
	fmt.Println(len(foods))
	fmt.Println(cap(foods))
	return foods
}

func main() {
	scores := [5]float64{3, 45, 666, 8, 99}
	fmt.Println(average(scores))

	pet := "nancy"
	fmt.Println(found(pet))

	groceryList := addGroceries("beets", "chocolate", "wine")
	fmt.Println(groceryList)
}
