package main

import "fmt"

/* Function to add two integer values, a, b int means both values will have type int and return type will also be int */
func add(a, b int) int {
	return a + b
}

/* Returning multiple values from function, we have to define their type in a parentheses in same order */
func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

/* Functions are first class citizens in Go, meaning you can pass a function as a parameter in it and return a function from a function, and also store it in a variable */
func processIt(fn func(a int) int) { // receiving a function as a param
	fn(1)
}

func processItTwo() func(a int) int { // returning a function from a function
	return func(a int) int {
		return a
	}
}

func main() {
	result := add(3, 5)               // calling add function and storing result in a variable
	lang1, _, lang3 := getLanguages() // storing three values in lang1, _, 3 respectively
	fmt.Println(result)
	fmt.Println(lang1, lang3)

	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	fun := processItTwo()
	fmt.Println(fun(6))
}
