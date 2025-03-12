package helpers

import "math/rand"

type SomeStruct struct {
	FirstName string
	LastName  string
}

func RandomNumber(numberPool int) int {
	return rand.Intn(numberPool)
}
