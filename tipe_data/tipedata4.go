package main

import "fmt"

func main() {
	var message = `Nama saya "John wick".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Print("message: \n", message)
}

/*
	(`) tanda grave accent/backtiks digunakan akan semua karakter yang ada didalamnya tidak di "escape"
*/
