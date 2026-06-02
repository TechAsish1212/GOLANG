package main

import "fmt"


func main() {
	// range---->In Go, the `range` keyword is used to iterate over elements in various data structures such as arrays, slices, maps, strings, and channels. It provides a convenient way to loop through these data structures without needing to manage index variables or iterators manually. The `range` keyword can be used in a `for` loop to access both the index and the value of each element in the data structure.

	// syntax of range in a for loop
	// for index, value := range dataStructure {
	//     // code to execute for each element
	// }

	// Iterating over an array using range
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		println("Index:", index, "Value:", value)
	}

	// Iterating over a map using range
	m := map[string]string{
		"name": "Asish",
		"city": "Bangalore",
	}

	for key, value := range m {
		println("Key:", key, "Value:", value)
	}

	// Iterating over a slice using range
	s:=[]int{10,20,30,40,50};
	sum:=0;
	for _,v:=range s{
		sum=sum+v;
	}
	fmt.Printf("sum of all values:%d\n",sum);

	// Iterating over a string using range
	str := "Hello, World!"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}