package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type People struct {
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name":"clarke",
			"last_name":"Kent",
			"has_dog":true

		},
		{
			"first_name":"subhash ",
			"last_name":"veeramachaneni",
			"has_dog":false


		}
	]`

	var unmarshalled []People

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error Unmarshalling json", err)
	}
	log.Printf("unmarshalled : %v", unmarshalled)

	//Writing json from a struct

	var mySlice []People

	var m1 People
	m1.Firstname = "Jeevan"
	m1.LastName = "veera"
	m1.HasDog = true

	mySlice = append(mySlice, m1)
	var m2 People
	m2.Firstname = "Jan"
	m2.LastName = "era"
	m2.HasDog = false
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error Marshalling", err)

	}
	fmt.Println(string(newJson))
}
