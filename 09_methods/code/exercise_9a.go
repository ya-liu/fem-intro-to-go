package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group represents a set of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func (u User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// func describeGroup
// => "This user group has 19 users. The newest user is Joe Smith. Accepting New Users: true"

// If you are planning to modify the data in a permanent way, use a pointer value to gain access to the memory address
func (g *Group) describe() string {

	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting New Users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	u2 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

	u3 := User{ID: 3, FirstName: "Ingrid", LastName: "Bergman", Email: "ingrid.bergman@gmail.com"}

	g := Group{
		role:           "admin",
		users:          []User{u, u2, u3},
		newestUser:     u3,
		spaceAvailable: true,
	}

	fmt.Println(u.describe())
	fmt.Println(g.describe())
	fmt.Println(g)
}
