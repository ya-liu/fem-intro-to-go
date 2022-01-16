// // Uncomment the entire file

package main

import "fmt"

func main() {

	var city string

	switch city {
	case "Des Moines":
		fmt.Println("You live in Iowa")
	case "Minneapolis,", "St Paul":
		fmt.Println("You live in Minnesota")
	case "Madison":
		fmt.Println("You live in Wisconsin")
	default:
		fmt.Println("You're not from around here.")
	}

	// // ****************************
	// var i int

	// switch {
	// case i > 10:
	// 	fmt.Println("Greater than 10")
	// case i < 10:
	// 	fmt.Println("Less than 10")
	// default:
	// 	fmt.Println("Is 10")
	// }

	// ****************************
	var i int = 9

	switch {
	case i != 10:
		fmt.Println("Does not equal 10")
		fallthrough
		// the fallthrough statement is a way to control whether to keep the case statement to continue executing code once it reaches something that returns true
		// here the i is 9, not 10, it will keep going to the next case to evaluate whether i is less than 10
	case i < 10:
		fmt.Println("Less than 10")
	case i > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}
}
