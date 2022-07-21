package main

import "fmt"

func main() {
	/*
		seleksi kondisi digunakan untuk mengontrol alur program, analoginya seperti rambu2 lalu lintas
		di jalan raya. kapan kendaraan di perbolehkan maju dan kapan harus berhenti

		yang dijadikan sebagai acuan oleh seleksi kondisi adalah bertipe data "bool" bisa berasal dari
		variabel ataupun hasil perbandingan nilai tsb yang menentukan blok kode mana yang ingin di eksekusi

		go memiliki 2 seleksi kondisi yaitu :
		-if else
		-switch
	*/
	var point = 3

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus, nilai kamu adalah:", point)
	}

}
