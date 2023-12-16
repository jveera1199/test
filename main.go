package main

import (
	"fmt"
	"log"
)

// Package level scope, x is a global variable
// In this case 'var x' assumes a type from the assignment
var x = "This is a string, defined using the var keyword in the global namespace"

// You could initialise the variable x in formal long form as:
// 'var x string = "initial string"'

type info struct {
	result string
}
type myStruct struct {
	FirstName string
	LastName  string
}

func (m *info) PrintResult() string {
	return m.result
}

func main() {
	fmt.Printf("%v \n", x) // uses the package level variable x
	myvar1 := myStruct{
		FirstName: "SVEER",
	}
	log.Println("myStruct Fname is:=  ", myvar1.FirstName)
	myvar2 := info{
		result: "Veerama",
	}

	log.Println("function calling the struct is:", myvar2.PrintResult())

	myMap3 := make(map[string]myStruct)
	me := myStruct{
		FirstName: "SV",
		LastName:  "Veer",
	}
	myMap3["me"] = me
	log.Println("map on strut is :", myMap3["me"].FirstName, "", myMap3["me"].LastName)
	myMap := make(map[string]string)
	myMap["dog"] = "samson"
	myMap["otherdog"] = "den"
	log.Println("my dogs name is  :", myMap["dog"], "my other dogs name is :", myMap["otherdog"])

	myMap2 := make(map[string]int)
	myMap2["isString"] = 1
	myMap2["notString"] = 2
	log.Println("maps are", myMap2["isString"], myMap2["notString"])

}
