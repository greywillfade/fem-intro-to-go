package main

import "fmt"

func main() {

	var mySentence string = "this is a sentence yo"
	for index, letter := range mySentence {
		if index%2 != 0 {
			fmt.Println("Index:", index, "Letter:", string(letter))
		}
	}

}
