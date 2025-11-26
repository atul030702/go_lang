package main

import (
	"fmt"
)

// slice -> dynamic arrays
// most commonly used construct in go
// + useful methods
func main() {
	/* 1. uninitialized slice is nil */
	//var num []int
	//fmt.Println(num == nil) Output = true
	//fmt.Println(len(num)) Output = 0

	/* 2. creating slice using make method and giving initial capacity */
	// var num = make([]int, 0, 5)
	// capacity -> maximum numbers of elements can fit (capacity increases by double when elements exceed initial given capacity)
	// fmt.Println(cap(num)) Output = 5
	// fmt.Println(num)
	// num = append(num, 1)
	// num = append(num, 2)
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	/**/
	// num := []int{}
	// num = append(num, 1)
	// num = append(num, 2)
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	// var num = make([]int, 0, 5)
	// num = append(num, 2)

	// var num2 = make([]int, len(num))

	/* copy function */
	// copy(num2, num)

	// fmt.Println(num, num2)

	/* Slice operator */
	// var num = []int{1, 2, 3, 4, 5}
	// fmt.Println(num[0:2])
	// fmt.Println(num[:2])
	// fmt.Println(num[1:])

	/* Slice Package (inbuilt package) */
	// var num1 = []int{1, 2, 3}
	// var num2 = []int{1, 2, 4}
	// fmt.Println(slices.Equal(num1, num2))

	/* 2d slices */
	var num = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(num)
}
