package main

import "fmt"

type Address struct {
	City     string
	Country  string
	PinCode  string
	District string
}

type Employee struct{
	Name string
	EmployeeId string
	Age int64
	Gender string
	Address
}

func main() {

	address1:=Address{
		City: "Kolkata",
		PinCode: "721443",
		District: "Paschim Medinipur",
		Country: "India",
	}

	e1:=Employee{
		EmployeeId: "CX-4321",
		Name: "Asish Kumar Bera",
		Age: 22,
		Gender: "Male",
		Address: address1,
	}

	fmt.Println(e1);
}
