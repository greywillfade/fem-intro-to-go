package main

import "fmt"

// func average(f1, f2, f3 float32) float32 {
// 	return (f1 + f2 + f3) / 3
// }

func average(inputs ...float32) float32 {
	var total float32

	for _, num := range inputs {
		total += num
	}
	return (total / float32(len(inputs)))
}

func main() {
	var result float32 = average(1, 2, 3)
	fmt.Println(result)
}
