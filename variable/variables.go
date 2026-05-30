package main

import "fmt"

func main() {
	// string
	// var name string ="Asish Kumar Bera";
	// var name = "Asish Kumar Bera"; // type inference
	// short hand declaration
	name := "Asish Kumar Bera"
	fmt.Println(name)

	// int
	// var age int = 21;  //int is the default type for whole numbers . otherwise(if we use int32 or int64) it will be treated as int and it will be converted to int32 or int64 based on the platform we are running the code on.
	var age = 21 // type inference
	// short hand declaration
	// age := 21; // this will cause an error because we have already declared the variable age in the same scope. we can not redeclare a variable in the same scope. we can only reassign a value to the variable.
	fmt.Println(age)

	// boolean
	var isStudent bool = true
	fmt.Println(isStudent)

	// float
	// var pi float64 = 3.14;
	pi := 3.14 // type inference and short hand declaration
	fmt.Println("pi:", pi)
}
