package main

import (
	"fmt"
	"time"
)

// order
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func main() {
	// Structs---->A struct (short for "structure") is a composite data type in Go that groups together zero or more fields with varying data types. It allows you to create complex data structures that can represent real-world entities or concepts. Structs are defined using the `type` keyword, and you can create instances of a struct to store and manipulate data.

	// syntax for defining a struct
	// type StructName struct {
	// 	Field1 DataType1
	// 	Field2 DataType2
	// 	...
	// }

	myOrder1 := Order{
		id:     "ORDER 1",
		amount: 2000.00,
		status: "available",
	}

	myOrder2 := Order{
		id:        "ORDER 2",
		amount:    4000.00,
		status:    "out of stock",
		createdAt: time.Now(),
	}

	myOrder1.createdAt = time.Now()

	// fmt.Println("Our Order Details:: ", myOrder1)
	// fmt.Println("Order Status:: ", myOrder1.status)

	fmt.Println("Our Order Details:: ", myOrder2)

	fmt.Println("Order Status:: ", myOrder2.status)

}
