package main

import (
	"fmt"
	// "time"
)

func main() {

	// user input
	// var age int
	// fmt.Println("Enter your age: ")
	// fmt.Scan(&age)

	// if condition
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else {
	// 	fmt.Println("You are a minor.")
	// }

	// else if condition
	// if age>=18{
	// 	fmt.Println("You are Adult");
	// }else if age >=12 {
	// 	fmt.Println("You are Teenager");
	// }else{
	// 	fmt.Println("You are Child");
	// }


	// switch case
	// simple switch case
	// day := 0
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// multiple cases
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a weekday.")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I am an integer: %d\n", v)
		case string:
			fmt.Printf("I am a string: %s\n", v)
		default:
			fmt.Printf("I am of a different type: %T\n", v)
		}
	}

	whoAmI(3.4)
}
