package main

import "fmt"

//1. Instantiate an array of scores
//- The array should have at least 5 elements of type `float64`
var scoresArray = [5]float64{1, 2, 3, 5, 10}
var petNames map[string]string

//1. Instantiate a slice that has an initial value of a collection of groceries.
var groceries = []string{"tea", "cake", "wine"}

//2. Write a function that calculates and returns the average score (also a float)
//- Use the `range` keyword
func calcAvg(inputArray [5]float64) float64 {
	var subtotal float64 = 0
	for i := range inputArray {
		subtotal += inputArray[i]
	}
	return subtotal / float64(len(inputArray))
}

//2. Write a function that takes a string argument and returns a boolean indicating whether or not that key exists in your map of pets.
func petExists(petCheck string) bool {
	_, ok := petNames[petCheck]
	return ok
}

//2. Write a function that takes one or more groceries as strings and appends them to the slice, printing out the resulting list of groceries.
func moarGroceries(addItems ...string) []string {
	for _, i := range addItems {
		groceries = append(groceries, i)
	}
	return groceries
}

func main() {
	//1. Define a map that contains a set of pet names, and their corresponding animal type. i.e.: `"fido": "dog"`.
	petNames = make(map[string]string)
	petNames["Kasumi"] = "Cat"
	petNames["Selina"] = "Cat"
	petNames["Lola"] = "Dog"
	petNames["Tarka"] = "Otter"

	fmt.Println(calcAvg(scoresArray))
	fmt.Println(petExists("Kasumi"))
	fmt.Println(petExists("Mal"))

	moarGroceries("Pringles", "lemons")
	for _, value := range groceries {
		fmt.Println(value)
	}

}
