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

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func checkSpace(g Group) bool {
	if len(g.users) < 2 {
		g.spaceAvailable = true //using this in next chapter
		return true
	} else {
		g.spaceAvailable = false
		return false
	}
}

func describeGroup(g Group) string {
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting new users is %t.", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Sally", LastName: "L", Email: "sallymail@gmail.com"}

	//Note that we haven't actually implemented any checks on membership using checkSpace yet
	g := Group{
		role:           "admin",
		users:          []User{u, u2},
		newestUser:     u2,
		spaceAvailable: true,
	}

	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
}
