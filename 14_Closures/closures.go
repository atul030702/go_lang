package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

/*
Closure in Go behave same as closure in JavaScript,
means the function remembers the value of it's outer scope (lexical scope) and keep it in memory reference
*/
func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}
