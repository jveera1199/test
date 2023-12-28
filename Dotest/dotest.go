package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of division is ", result)

	addition, err := add(0, 11.1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Addition value is", addition)

}
func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil

}

func add(a, b float32) (float32, error) {
	var addition float32

	if a == 0 || b == 0 {
		return addition, errors.New("cannot add as 1 value is 0")
	}
	addition = a + b
	return addition, nil
}
