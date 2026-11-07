package main

import "fmt"


func main() {
	fmt.Println("We are learning defer in golang");
	fmt.Println("start line");
	defer fmt.Println("middle line");
	defer fmt.Println("end line");

	//  if two defer statements are used then the last defer statement will be executed first and the first defer statement will be executed last. LIFO (Last In First Out) principle is followed in defer statements.
}