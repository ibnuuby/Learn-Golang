package main

import "fmt"

func main() {
	/*
		fungsi "new" digunakan untuk membuat variabel "pointer" dengan tipe data tertentu. nilai
		data defaultnya akan menyesuaikan tipe datanya.

	*/
	name := new(string)
	fmt.Println(name)  // 0xc00008a210
	fmt.Println(*name) // ""

	/*
		Variabel name menampung data bertipe pointer string. Jika ditampilkan yang muncul bukanlah nilainya melainkan
		alamat memori nilai tersebut (dalam bentuk notasi heksadesimal). Untuk menampilkan nilai aslinya, variabel
		tersebut perlu di-dereference terlebihdahulu, menggunakan tanda asterisk (*).
	*/
}
