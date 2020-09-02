package utils

import (
	"fmt"
	"strings"
)

func printString(s string) {
	fmt.Println("The string: ", s)
}

//MakeExcited transforms to all caps with exclamation
func MakeExcited(s string) string {
	return strings.ToUpper(s) + "!"

}
