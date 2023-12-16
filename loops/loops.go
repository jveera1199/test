package main

import "log"

func main() {
	animals := []string{"cat", "dog", "fish", "cow"}
	for _, animal := range animals {
		log.Println(animal)
	}

	type User struct {
		FirstName string
		LastName  string
		email     string
		age       int
	}

	var users []User
	users = append(users, User{"john", "smith", "js@gmail", 12})
	users = append(users, User{"john", "th", "js@gmail", 22})
	users = append(users, User{"ohn", "ther", "js@gmail", 21})

	for i, k := range users {
		log.Println(i, k)
	}

}
