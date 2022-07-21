package main

import "fmt"

func main() {
	var firstText string = "John"
	lastText := "Wick"
	/*
		":=" digunakan sebagai penanda bahwa tipe data tidak digunakan
		tanda ":=" hanya dapat digunakan selama sekali di awal pada saat deklarasi.
	*/

	fmt.Printf("Helo %s %s! \n", firstText, lastText)
	/*
		"%s" digunakan sebagai media pengganti pada variable "fristText & lastText"
	*/
}
