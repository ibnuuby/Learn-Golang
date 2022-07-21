package main

import "fmt"

func main() {

	/*
		Mutli variabel
		Gulang bisa menggunakan multi variabel yaitu dengan menggunakan tanda "," untuk membuat
		setiap penamaan variabel lain.
		dan mutli variabel ini bisa menggunakan "type inference (metoe tanpa tipe data)"
	*/

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	seventh, eighth, ninth := "tujuh", "delapan", "sembilan"

	fmt.Printf(first, second, third)
	fmt.Printf(fourth, fifth, sixth)
	fmt.Printf(seventh, eighth, ninth)

}
