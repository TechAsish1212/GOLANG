package main

import "fmt"

func main() {
	// while loop is not available in Go, but we can achieve the same functionality using for loop.
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	fmt.Println("1")
	// }

	// for loop with initialization, condition and post statement
	// for i :=1;i<=5;i++{
		// fmt.Println(i);
	// }

	// for i := 1; i <= 5; i++ {
	// 	if i==2{
	// 		continue; // this will skip the current iteration and move to the next iteration of the loop.
	// 	}
	// 	fmt.Println(i)
	// }

	// range loop
	for i := range 3{
		fmt.Println(i)
	}
}
