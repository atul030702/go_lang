package main

import "fmt"

/* variadic functions are types of functions which can receive n number of parameters in go */
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	nums := []int{3, 4, 55, 6}
	result := sum(3, 4, 5, 6, 7, 25)
	result2 := sum(nums...)
	fmt.Println(result, result2)
}
