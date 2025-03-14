package main

import (
	"errors"
	"log"
)

func main() {

	result, error := divide(10, 0)

	if error != nil {
		log.Println("Error: ", error)
		return
	}

	log.Println("Result: ", result)
}

func divide(x, y float64) (float64, error) {

	var result float64

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil
}
