package main

import "fmt"


func main() {
	// map---->In Go, a map is a built-in data structure that stores key-value pairs, similar to objects in JavaScript or dictionaries in Python. Maps are used to associate unique keys with corresponding values, allowing for efficient retrieval of data based on the key. A map is defined using the built-in `map` type, and it can be created using the `make` function or by using a map literal.

	// creating map
	m :=make(map[string]string) // creating a map with string keys and string values
	m["name"]="Asish" // adding key-value pairs to the map
	fmt.Println(m["name"]) // printing the value associated with the "name" key

	// Declaring and Initializing a Map
	 student := map[string]int{
        "Asish": 25,
        "Rahul": 30,
        "Amit":  28,
    }
    fmt.Println(student) // printing the entire map
	
	// Checking if a Key Exists
	age, exists := student["Asish"] // retrieving a value from the map using a key
	fmt.Println("Age:", age, "Exists:", exists) // printing the retrieved value and whether the key exists in the map

	// updateing a value in the map
	student["Asish"]=21;
	fmt.Println(student) // printing the updated map
}