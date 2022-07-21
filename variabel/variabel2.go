package main

import "fmt"

func main() {
	var lastName string
	var firstName string = "John"

	fmt.Printf("Hello John Wick! \n")
	// "\n" digunakan untuk membuah break line
	fmt.Printf("Hello %s %s", firstName, lastName+"!\n")
	// "%s"digunakan sebagai pengganti data pada karakter "string" > lastName
	fmt.Println("Hello", firstName, lastName+"!")
	/*
		"+" digunakan untuk melakukan penggabungan string
	*/
}
