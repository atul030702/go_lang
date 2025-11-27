package main

import "fmt"

func main() {
	/* Int type Array */
	// int -> 0, string -> "", bool -> false (default values of array types)
	// var nums [4]int

	// nums[0] = 1
	// fmt.Println(nums[0]) // element at 0th index
	// fmt.Println(nums)
	// fmt.Println(len(nums)) // array length

	/* Bool type Array */
	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	/* String type Array  */
	// var name [3]string
	// name[0] = "Atul"
	// fmt.Println(name)

	/* To declare array in single line */
	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	/* 2D Arrays */
	nums := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums)

	/*
		- If fixed size, that is predictable
		- Memory optimization
		- constant time access
	*/
}
