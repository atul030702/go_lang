package main

import "fmt"

// Range: keyword used to iterate over data structures
func main() {
	/* Iterating slice using range */
	nums := []int{6, 7, 8}

	for i, num := range nums {
		fmt.Println(num, i)
		// i is index of element, num is the element of slice
	}

	/* Iterating maps using range */
	first_map := map[string]string{"first_name": "John", "last_name": "Abraham"}

	for k, v := range first_map {
		fmt.Println(k, v) // k is key, v is value of map

	}

	/* Iterating over strings using range */
	for i, c := range "golang" {
		fmt.Println(i, c, string(c)) // i is starting byte of rune, c is unicode point rune for character, string(c) will convert unicode into alphabet character
		/*
			Generally characters are stored in 255 in ascii encoding, till 255 it takes 1byte of memory and after 255 it takes 2byte,
			if first characters is 300 in ascii encoding it will take 2byte, so second character will show 2 as i instead of 1 as 0, 1 is taken by 1st character
		*/
	}
}
