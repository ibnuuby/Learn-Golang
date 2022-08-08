package main

import "fmt"

func main() {
	/*
		Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice.
		Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len, tapi
		bisa berubah seiring operasi slice yang dilakukan. Agar lebih jelas, silakan
		disimak kode berikut.
	*/
	var buah = []string{"apel", "anggur", "pisang", "melon"}
	var abuah = append(buah, "manggis")
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
	fmt.Println(abuah)

	/*
		append() digunakan untuk menambahkan elemen pada slice.
		 Elemen baru tersebut diposisikan setelah indeks paling akhir
	*/

}
