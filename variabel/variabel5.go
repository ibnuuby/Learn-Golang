package main

import "fmt"

func main() {
	/*
		Underscore "_" adalah reserved variabel yang bisa dimanfaatkan untuk menampung nilai
		yang tidak dipakai. bisa di bilang variabel ini adalah "keranjang sampah"
	*/

	_ = "belajar golang"
	_ = "golang itu mudah"
	text, _ := "John", "Wick"

	/*
		pada contoh di atas variabel "text"  yang berisikan "John", sedangkan nilai "wick" akan ditampung
		oleh variabel "_" yang menandakan bahwa nilai tsb tidak akan digunakan
	*/

	fmt.Println(text)

}
