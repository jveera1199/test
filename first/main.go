package main

import (
	"fmt"
)

// Package level scope, x is a global variable
// In this case 'var x' assumes a type from the assignment
var x = "This is a string, defined using the var keyword in the global namespace"

// You could initialise the variable x in formal long form as:
// 'var x string = "initial string"'

func main() {
	fmt.Printf("%v \n", x) // uses the package level variable x

}
