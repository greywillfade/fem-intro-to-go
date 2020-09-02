package main

import (
	"fem-intro-to-go/05_toolkit/code/utils"
	"fmt"
)

func calculateImportantData() int {
	totalValue := utils.Add(34, 2, 8)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
}
