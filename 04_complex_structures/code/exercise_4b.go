package main

import (
	"fmt"
	// "reflect"
)

func main() {
	fmt.Println(average(1, 2, 3))
}

func average(args ...float64) float64 {
	// fmt.Println(reflect.TypeOf(len(args)))
	total := 0.0
	if len(args) == 3 {
		for _, num := range args {
			total += num
		}
	}
	return total / 3
}
