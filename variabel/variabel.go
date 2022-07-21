package main

import "fmt"

func main() {

	/*
		variabel di golang mempunyai 2 jenis penulisan, yaitu :
		- dituliskan dengan tipe datanya
		- dituliskan tanpa menggunakan tipe datanya

		variabel sendiri digunakan sebagai media penyimpanan sementara dan memiliki bermacam-macam tipe
		datanya.

	*/
	var firstName string = "Si"
	/*
		"var" digunakan sebagai mendeklarasikan sebuah variabel dan "string" merupakan tipe data
		 yang digunakan
		firstName : nama variabel
	*/
	var lastName string
	lastName = "Koding"

	/*
		"fmt.Println" digunakan untuk mencetak hasil dari isi variabel
	*/
	fmt.Println(firstName, lastName)

}
