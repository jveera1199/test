package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	age       int
}

type Partnership struct {
	People     []Person
	commanName string
}

func (p *Person) Fullname() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
func (p *Person) ChangeLastName(s string) {
	p.LastName = s

}
func (p *Partnership) setCommanName() {
	if p.commanName != "" {
		for i := 0; i < len(p.People); i++ {
			p.People[i].ChangeLastName(p.commanName)

		}
	}

}

func whoamI(p Person) {
	fmt.Printf("The name of the Person is %s and age of the Person is %d  \n ", p.Fullname(), p.age)
}

func main() {
	p1 := Person{
		FirstName: "SVeer",
		LastName:  "Veerama",
		age:       27,
	}
	p2 := Person{
		FirstName: "VeeraS",
		LastName:  "neni",
		age:       25,
	}
	rel := Partnership{[]Person{p1, p2}, "Mixers"}

	for _, p := range rel.People {
		whoamI(p)
	}
	rel.People[1].ChangeLastName("Veeramachaneni")

	rel.setCommanName()
	for _, p := range rel.People {
		whoamI(p)
	}

}
