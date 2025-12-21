package main

import "fmt"

// num will be copied by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1

	changeNum(&num)

	fmt.Println("Memory Address", &num)
	fmt.Println("After changeNum in main", num)
}
