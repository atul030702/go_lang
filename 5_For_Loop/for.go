package main

import "fmt"

// for loop -> only construct in go for looping
func main() {
	/* while loop implementation using for */
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	/* infinite loop */
	// for {
	// 	println("1")
	// }

	/* classic for loop */
	// for i := 0; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue // Skips the i == 2 iteration
	// 	}

	// 	fmt.Println(i)
	// }

	// range
	for i := range 3 {
		fmt.Println(i)
	}
}
