package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice --> In Go, a slice is a flexible, dynamically-sized view into an underlying array. Slices are one of the most commonly used data structures in Go. They provide a more powerful and convenient way to work with sequences of data compared to arrays. A slice is defined by three components: a pointer to the underlying array, the length of the slice, and its capacity.

	// var nums []int // declaring a slice of integers
	// fmt.Println("Slice:", nums)
	// fmt.Println(nums==nil)

	// Creating a slice using make function
	// var nums=make([]int, 5) // creating a slice of integers with length 5
	// fmt.Println("Slice:", nums)
	// fmt.Println(nums==nil)
	// fmt.Println(cap(nums)); // capacity of the slice

	// nums=append(nums, 10) // appending a value to the slice
	// nums=append(nums, 20)
	// fmt.Println("Slice after appending:", nums);
	// fmt.Println(cap(nums)); // capacity of the slice


	// slice initialization
	// nums := []int{}
	// nums=append(nums,1,2);
	// fmt.Println("Initialized slice:", nums)
	// fmt.Println(cap(nums)) // capacity of the slice

	// slice operator
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println("Original slice:", nums)

	// // slicing the slice
	// subSlice := nums[1:4] // creates a new slice from index 1 to 3 (4 is exclusive)
	// fmt.Println("Sub-slice:", subSlice)


	
	nums1:= []int{1, 2, 3, 4, 5}
	nums2:= []int{1, 2, 3, 4, 5}
	// nums2:= []int{6, 7, 8, 9, 10}

	fmt.Println(slices.Equal(nums1,nums2)); // checks if two slices are equal

	// 2d slice
	var matrix [][]int // declaring a 2D slice
	matrix = append(matrix, []int{1, 2, 3}) // appending rows to the 2D slice
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	fmt.Println("2D Slice (Matrix):")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	
}