package main

import "fmt"

func main() {
	mySentence := "Go get them tiger"

	for index, letter := range mySentence {
		if index%2 != 0 {
			fmt.Println("Letter:", string(letter))
		}
	}
}
