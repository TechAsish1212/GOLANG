package main

import "fmt"

func main() {
	// array declaration 
	// var nums [5]int;
	// nums[0] = 10;
	// fmt.Println("Arrays",nums[0]);
	// fmt.Println("Total size of array: ",len(nums));   // length of array

	// var num int;
	// fmt.Println("Enter the size of array:: ");
	// fmt.Scan(&num);

	// var arr [100]int;

	// fmt.Println("Enter the values:: ");
	// for i:=1;i<=num;i++{
	// 	fmt.Scan(&arr[i]);
	// }

	// fmt.Println("Output::---->");
	// for i:=1;i<=num;i++{
	// 	fmt.Println(arr[i]);
	// }

	// array initialization
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", nums)
	
	// 2d array
	// var matrix [3][3]int
	// matrix[0][0] = 1
	// matrix[0][1] = 2
	// matrix[0][2] = 3
	// matrix[1][0] = 4
	// matrix[1][1] = 5
	// matrix[1][2] = 6
	// matrix[2][0] = 7
	// matrix[2][1] = 8
	// matrix[2][2] = 9

	// fmt.Println("2D Array (Matrix):")
	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		fmt.Printf("%d ", matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// 2d array initialization
	   matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    fmt.Println(matrix)

	// array uses
	// 1. Storing multiple values of the same type
	// 2. Iterating over a collection of data
	// 3. Passing arrays to functions
	// 4. Multidimensional arrays for complex data structures
	// 5. fixed-size collections where the size is known at compile time
	// 6. performance-critical applications where memory allocation and access patterns are important
	// 7. contiguous memory allocation for better cache performance
} 