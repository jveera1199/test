package main

import "log"

func main() {
	var mySlice []string
	mySlice = append(mySlice, "jeevan")
	mySlice = append(mySlice, "sai")
	mySlice = append(mySlice, "subhash")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	log.Println(numbers)
	log.Println(numbers[1:3])

}
