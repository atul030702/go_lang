package main

import (
	"fmt"
	"maps"
)

// maps -> similar to JS object
func main() {
	/* creating map using make method */
	object := make(map[string]string)

	/* setting as element in map */
	object["name"] = "golang"
	object["area"] = "backend"

	/* get an element */
	// fmt.Println(object["name"], object["area"])

	/* If key doesn't exist in map then it returns zero value */
	// fmt.Println(object["phone"]) // Output: zero value ("") in string case, 0 in int case

	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 50
	// fmt.Println(m["phone"])

	/* Getting the length of the map using len() method */
	// fmt.Println(len(m))

	/* To delete the element in the map */
	delete(m, "price")
	// fmt.Println(m)

	/* To clear all the elements in the map */
	clear(m)
	// fmt.Println(m)

	/* Creating map if elements are known */
	map3 := map[string]int{"price": 40, "age": 23, "phones": 3}
	// fmt.Println(map3)
	_, ok := map3["phones"] // _ is for value of the key given, here phones means _ = 3, ok is boolean value means that key exist in map
	// fmt.Println()

	if ok {
		// fmt.Println("all ok")
	} else {
		// fmt.Println("not ok")
	}

	/* To check two maps are equal or not */
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 8}

	fmt.Println(maps.Equal(m1, m2))
}
