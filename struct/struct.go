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
// In Go, when you define a method with a value receiver (like `func (o Order) changeStatus(status string)`), it operates on a copy of the struct. This means that any changes made to the struct's fields within the method will not affect the original struct outside the method. In this case, when you call `myOrder1.changeStatus("new status")`, it will change the status of the copy of `myOrder1` inside the method, but the original `myOrder1` will remain unchanged. To modify the original struct, you would need to use a pointer receiver (e.g., `func (o *Order) changeStatus(status string)`), which allows you to modify the original struct through its memory address.
func (o *Order) changeStatus(status string) {
	o.status = status
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

	myOrder1.changeStatus("out of stock");

	// myOrder2 := Order{
	// 	id:        "ORDER 2",
	// 	amount:    4000.00,
	// 	status:    "out of stock",
	// 	createdAt: time.Now(),
	// }

	myOrder1.createdAt = time.Now()

	fmt.Println("Our Order Details:: ", myOrder1)
	fmt.Println("Order Status:: ", myOrder1.status)

	// fmt.Println("Our Order Details:: ", myOrder2)
	// fmt.Println("Order Status:: ", myOrder2.status)

}
