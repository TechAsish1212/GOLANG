package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	IsAdult bool   `json:"is_adult"`	
}

func main() {
	fmt.Println("We are learning Json");
	person :=Person{
		Name:    "John Doe",
		Age:     30,
		Gender:  "Male",
		IsAdult: true,
	}
	fmt.Println("Person Details: ", person)

	// convert to json encoding format(marshaling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	fmt.Println("JSON Data: ", string(jsonData))

	// decoding json data(unmarshaling)
	var personData Person;
	err=json.Unmarshal(jsonData, &personData);
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Unmarshaled Person Data: ", personData)
}	
