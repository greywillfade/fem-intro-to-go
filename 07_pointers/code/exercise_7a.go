package main

import "fmt"

//A user is a set of attributes describing a person
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// 2. Write a function called `updateEmail` that takes in a `*User` type
func updateEmail(u *User, newEmail string) {

	// 3. Update the user's email to something new
	u.Email = newEmail
}

func main() {
	fmt.Println("Pointers!")

	// 1. Define an instance of the User struct
	var user1 = User{ID: 1, FirstName: "Sally", LastName: "L", Email: "testemail@yo.no"}
	fmt.Println(user1)

	// 4. Call `updateEmail()` from `main()` and verify the updated email has persisted
	updateEmail(&user1, "newemail@whodis.com")
	fmt.Println(user1)

}
