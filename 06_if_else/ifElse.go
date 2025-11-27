package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is minor.")
	// }

	/*
		age := 3
		if age >= 16 {
			fmt.Println("Person is an adult.")
		} else if age >= 12 {
			fmt.Println("Person is teenager.")
		} else if age >= 0 && age < 3 {
			fmt.Println("Person is infant")
		} else {
			fmt.Println("Person is kid.")
		}
	*/

	/* We can declare a variable inside if block */
	if age := 25; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	} else {
		fmt.Println("Person is a child", age)
	}

	/* Go doesn't have ternary, you will have to use normal if else */
}
