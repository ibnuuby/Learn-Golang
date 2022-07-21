package main

import "fmt"

func main() {
	var decimalNumber = 2.62
	fmt.Printf("bilangan decimal: %f\n", decimalNumber)
	fmt.Printf("bilangan decimal: %.3f\n", decimalNumber)

}

/*
	pada variabel "decimalNumber" adalah bertipe data "float32"
	%f digunakan untuk memformat data numerik decimal menjadi "string"
	%.3f digunakan untuk mengatur keluaran digit yang dihasilkan, jadi pada text "3" bisa diganti
	dengan kemauan kita sendiri.

*/
