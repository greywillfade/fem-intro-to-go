package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group is a collection of similar users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

//Refactoring the original function to be a method instead
func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) describe() string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting new users is %t.", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Sally", LastName: "L", Email: "sallymail@gmail.com"}
	u3 := User{ID: 3, FirstName: "Sally2", LastName: "L", Email: "sallymail2@gmail.com"}

	g := Group{
		role:           "admin",
		users:          []User{u, u2, u3},
		newestUser:     u2,
		spaceAvailable: true,
	}

	userDesc := u.describe()
	gDesc := g.describe()
	fmt.Println(userDesc)
	fmt.Println(gDesc)
}
