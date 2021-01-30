package main

import (
	"encoding/json"
	"log"
)

// Define schema
// Translate "first_name" in json to FirstName in Go and vice-versa.
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJSON := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true 
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false 
		}
	]
	`

	// JSON to STRUCT
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJSON), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled : %v", unmarshalled)

	// STRUCT to JSON
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Prince"
	m2.LastName = "Diana"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")

	if err != nil {
		log.Println("Error marshalling into json", err)
	}

	log.Printf("marshalled : %v", string(newJson))

}
