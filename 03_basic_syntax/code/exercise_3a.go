package main

import "fmt"

func main() {
	mySentence := "Go get them tiger"

	for index, letter := range mySentence {
		if index%2 != 0 {
			fmt.Println("Letter:", string(letter))
		}
	}

	// if you don't wanna use a variable and want to avoid go's error message, use an underscore to replace any unused variable
	for _, value := range mySentence {
		fmt.Println(string(value))
	}
}
