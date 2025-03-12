package helpers

import "math/rand"

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	HasDog    bool   `json:"has_dog"`
}

func RandomNumber(numberPool int) int {
	return rand.Intn(numberPool)
}
