package main

import (
	"log"

	helpers "github.com/jveera1199/test/Helpers"
)

func main() {
	var myVar helpers.MyNameis
	myVar.Name = "Sveer"
	myVar.Lname = "Veerama"
	log.Println(myVar, myVar.Name, myVar.Lname)
}
