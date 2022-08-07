package main

import "fmt"

func main() {
	var buah = [4]string{
		"apple",
		"grape",
		"banana",
		"melon"}
	fmt.Println("Jumlah buah \t\t", len(buah))
	fmt.Println("Isi semua element \t", buah)
}
