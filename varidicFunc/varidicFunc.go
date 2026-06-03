package main

import "fmt"


func sum(nums ...int)int{ // the `...` syntax indicates that the function can accept a variable number of arguments of type `int`. Inside the function, `nums` is treated as a slice of integers, allowing us to iterate over it and perform operations on the provided arguments.
	total:=0;
	
	for _,num:=range nums{
		total=total+num;
	}

	return total;
}

func main() {
	// Variadic Functions---->In Go, a variadic function is a function that can accept a variable number of arguments. This is useful when you want to create functions that can handle different numbers of inputs without having to define multiple function signatures. Variadic functions are defined using the `...` syntax in the parameter list.

	res:=sum(1,2,3,4,5,6,7,8,9);
	fmt.Println(res);
}