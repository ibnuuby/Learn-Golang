package main

import "fmt"

func main() {
	var buah = []string{"apel", "anggur", "jeruk", "mangga"}

	var aBuah = buah[0:3]
	var bBuah = buah[1:4]

	var aaBuah = buah[1:2]
	var bbBuah = buah[0:1]

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(bbBuah)

	bbBuah[0] = "pinaple"

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(bbBuah)

}
