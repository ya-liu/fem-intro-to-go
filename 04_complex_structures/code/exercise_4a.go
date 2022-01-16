package main

import "fmt"

func main() {
	fmt.Println(average(1, 2, 3))
}

func average(arg1, arg2, arg3 float64) float64 {
	return (arg1 + arg2 + arg3) / 3
}
