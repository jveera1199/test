package main

import "fmt"

type Animal interface {
	Says() string
	NumOfLegs() int
}
type Dog struct {
	Name, Breed string
}

type Cat struct {
	Name, colour string
}

func main() {
	dog := Dog{
		Name:  "ricky",
		Breed: "Germak shepard",
	}
	PrintInfo(&dog)

	cat := Cat{
		Name:   "martin",
		colour: "orange",
	}
	PrintInfo(&cat)

}
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "And has ", a.NumOfLegs(), "legs")
}
func (d *Dog) Says() string {
	return "woof"
}
func (d *Dog) NumOfLegs() int {
	return 4
}
func (d *Cat) Says() string {
	return "meow"
}
func (d *Cat) NumOfLegs() int {
	return 4
}
