package main

import "fmt"

func main() {
	var value = (((2+6)%3)*4 - 2) / 2
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}

/*
	(%d) untuk memunculkan nilai int ke dalam bentuk string
	(%t) untuk memunculkan nilai bool pada (printf)
*/
