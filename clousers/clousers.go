package main

import "fmt"

func counter() func() int { 
	var count int = 0
	return func() int {  // This inner function is a closure because it captures the `count` variable from the outer `counter` function. Each time the inner function is called, it has access to the `count` variable and can modify it, allowing us to keep track of the count across multiple calls to the closure.
		count++
		return count
	}
}

func main() {
	// Closures---->A closure is a function that captures and retains access to variables from its surrounding scope, even after that scope has finished executing. In Go, closures are created when you define a function inside another function and the inner function references variables from the outer function.

	increment := counter()

	fmt.Println(increment());
	fmt.Println(increment());
	fmt.Println(increment());
	fmt.Println(increment());


}
