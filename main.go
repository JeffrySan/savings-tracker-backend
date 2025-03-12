package main

import (
	"encoding/json"
	"fmt"
	"log"

	helpers "github.com/JeffrySan/go-packages-helpers/packages"
)

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

	var personResult []helpers.Person

	error := json.Unmarshal([]byte(myJson), &personResult)

	if error != nil {
		log.Println("Error on Unmarshalling")
		return
	}

	log.Printf("Success Unmarshalled %v", personResult)

	// Marshalling
	var mySlicePerson []helpers.Person

	var person1 helpers.Person
	person1.FirstName = "Bailey"
	person1.LastName = "Aisyle"
	person1.Age = 23
	person1.HasDog = false

	var person2 helpers.Person
	person2.FirstName = "Jason"
	person2.LastName = "Mars"
	person2.Age = 53
	person2.HasDog = true

	mySlicePerson = append(mySlicePerson, person1)
	mySlicePerson = append(mySlicePerson, person2)

	jsonResult, error := json.MarshalIndent(mySlicePerson, "", "   ")

	if error != nil {
		log.Println("Error Marshalling")
		return
	}

	fmt.Printf("Result JSON: %v", string(jsonResult))
}
