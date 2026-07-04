package main

import "fmt"

// enumered type for order status
// type OrderStatus int;

// const (
// 	Recieved OrderStatus=iota;
// 	Confirmed
// 	Shipped
// 	Delivered
// )

// another for string
type OrderStatus string

const (
	Recieved  OrderStatus = "recieved"
	Confirmed OrderStatus = "confirmed"
	Shipped   OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	// Enums in Go are a way to define a set of named constants that represent a specific type. They are often used to represent a fixed set of values, such as days of the week, months of the year, or status codes. In Go, enums are typically implemented using the `iota` identifier, which allows you to create a sequence of related constants with minimal code.

	// syntax for defining an enum
	// const (
	//     EnumValue1 = iota
	//     EnumValue2
	//     EnumValue3
	//     ...
	// )

	changeOrderStatus(Shipped)
}
