package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Clinton",
			"age": 23,
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"age": 33,
			"has_dog": false
		}
	]`

	var personResult []Person

	error := json.Unmarshal([]byte(myJson), &personResult)

	if error != nil {
		log.Println("Error on Unmarshalling")
		return
	}

	log.Printf("Success Unmarshalled %v", personResult)
}
