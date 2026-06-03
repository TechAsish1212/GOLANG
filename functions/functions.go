package main

import "fmt"

func add(a int, b int) int {
	return a + b 
}

func subtract(a, b int) int {  // we can omit the type of the second parameter if it is the same as the first parameter. so we can write it as subtract(a, b int) int
	return a - b
}

func getLang()(string,string ){ // we can return multiple values from a function in Go. we can specify the return types in the function signature and then return the values in the return statement.
	return "java","GO";
}


func main() {
	// Functions---->In Go, a function is a reusable block of code that performs a specific task. Functions are defined using the `func` keyword, followed by the function name, parameters (if any), and the return type (if any). Functions can take zero or more parameters and can return zero or more values.

	// syntax of a function
	// func functionName(parameters) returnType {
	//     // code to execute
	// }
	a :=add(4,7);

	fmt.Printf("add: %d ",a);
	sub:=subtract(10,5);
	fmt.Printf("subtract: %d \n",sub);

	lang1,lang2:=getLang();
	fmt.Printf("lang1: %s, lang2: %s \n",lang1,lang2);

}
