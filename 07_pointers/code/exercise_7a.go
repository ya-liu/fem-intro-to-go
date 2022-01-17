package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// YOUR CODE HERE
var u1 = User{
	ID:        1,
	FirstName: "Adele",
	LastName:  "Y",
	Email:     "adele@adele.com",
}

// a * next to the type name indicates the variable will point to a memory location
func updateEmail(user *User) {
	user.Email = "new@newEmail.com"
}

func main() {
	fmt.Println("Pointers!")

	// a & next to the pointer variable name gives you the pointer address
	updateEmail(&u1)
	fmt.Println(u1.Email)
}
