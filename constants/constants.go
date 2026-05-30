package main

import "fmt"

func main() {
	const name = "Asish Kumar Bera"
	const age = 21
	const isStudent = true
	const pi = 3.14

	// name = "John Doe" // this will cause an error because we can not reassign a value to a constant variable.

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isStudent)
	fmt.Println(pi)

	// grouping constants
	const (
		country = "India"
		city    = "Kolkata"
	)

	fmt.Println("Country: " + country + ", " + "City: " + city);
}
