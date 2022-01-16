// Uncomment the entire file

package main

import "fmt"

func main() {

	// ****************************

	// i := 1

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(i)
	// }

	// 	// ****************************

	// i := 1

	// for i <= 100 {
	// 	fmt.Println(i)
	// 	// This will behave like a while loop
	// 	i += 1
	// }

	// 	// ****************************

	var mySentence = "This is a sentence"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", letter)
		// if we don't convert the letter to type string, golang returns the byte of the letter, as an integer
		fmt.Println("Index:", index, "Letter:", string(letter))
	}
}
