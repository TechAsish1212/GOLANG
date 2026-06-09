package main

import "fmt"

// Interfaces in Go are a powerful way to define behavior and achieve polymorphism. An interface is a collection of method signatures that a type can implement. When a type implements all the methods defined in an interface, it is said to satisfy that interface. This allows you to write functions that can work with any type that satisfies the interface, promoting code reuse and flexibility.

// syntax for defining an interface
// type InterfaceName interface {
//     Method1(params) returnType
//     Method2(params) returnType
//     ...
// }

// To implement an interface, a type must define all the methods declared in the interface. There is no explicit declaration of intent to implement an interface; if a type has the required methods, it automatically satisfies the interface.


type payment struct {

}

func (p payment) makePayment(amount float64) {
	raz:= razorpay{};
	raz.pay(amount);
}

type razorpay struct {

}
	
func (r razorpay) pay(amount float64) {
	fmt.Println("Making Payment using Razorpay:: ",amount);
}
func main() {
	newPayment := payment{};
	newPayment.makePayment(100.50);
}