package main

import "fmt"

func main() {
	/* Ways of declaring constants in Go, you can't modify them after initialization */

	const name = "golang"
	const age = 30

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println()
}
