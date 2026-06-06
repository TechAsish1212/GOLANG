package main

import "fmt"

type Address struct {
	City     string
	Country  string
	PinCode  string
	District string
}

type Employee struct {
	EmployeeId int
	Name    string
	Age     int
	Address // This is struct embedding. By embedding the `Address` struct within the `Employee` struct, we can directly access the fields of `Address` from an instance of `Employee` without needing to reference the embedded struct. This promotes code reuse and allows us to organize related data together. For example, we can create an instance of `Employee` and set the city, country, pin code, and district directly on the employee instance, as if those fields were part of the `Employee` struct itself.
}

func main() {
	// Struct Embedding---->Struct embedding is a powerful feature in Go that allows you to create complex data structures by embedding one struct within another. This promotes code reuse and helps to organize related data together. When you embed a struct, the fields and methods of the embedded struct become part of the outer struct, allowing you to access them directly without needing to reference the embedded struct.

	// syntax for struct embedding
	// type OuterStruct struct {
	// 	EmbeddedStruct
	// 	Field1 DataType1
	// 	Field2 DataType2
	// 	...
	// }

	address1 := Address{
		City:     "kolkata",
		Country:  "India",
		PinCode:  "721443",
		District: "Paschim Medinipur",
	}

	employee1 := Employee{
		EmployeeId: 101,
		Name:    "Asish Kumar Bera",
		Age:     21,
		Address: address1,
	}

	fmt.Println("Employee Details:: ", employee1);
	fmt.Println("Employee Id: ", employee1.EmployeeId);
	fmt.Println("Employee Name:: ", employee1.Name);

}
