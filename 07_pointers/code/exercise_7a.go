package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// YOUR CODE HERE
func updateEmail(user *User) {
	user.Email = "new@newEmail.com"
}

func main() {
	fmt.Println("Pointers!")
	u1 := User{
		ID:        1,
		FirstName: "Adele",
		LastName:  "Y",
		Email:     "adele@adele.com",
	}
	updateEmail(&u1)
	fmt.Println(u1.Email)
}
