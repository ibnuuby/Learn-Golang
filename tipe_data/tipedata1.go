package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Println("bilangan positive: ", positiveNumber)
	fmt.Println("bilangan negative: ", negativeNumber)
}

/*
	variabel positiveNumber bertipe data "uint8" cangkupan nilai bilangannya yaitu (0  <==> 255)
	sedangkan negativeNumber merupakan tipe data "int32" dengan cangkupan nilainya yaitu (-2147483648 <==> 2147483647)


	Note:
	uint digunakan untuk bilangan cacah (bilangan negative)
	int digunakan untuk bilangan bulat (bilangan negative dan positive)

	di anjur untuk tidak menggunakan sembarang tipe data, karena akan mempengarungi alokasi memory variabel.
*/
