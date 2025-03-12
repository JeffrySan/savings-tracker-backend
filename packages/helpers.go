package helpers

import "math/rand"

type Person struct {
	first  string `json: "first"`
	last   string `json: "last"`
	Age    int    `json: "age"`
	HasDog bool   `json: "hasDog"`
}

func RandomNumber(numberPool int) int {
	return rand.Intn(numberPool)
}
