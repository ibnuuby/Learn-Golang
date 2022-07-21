package main

import "fmt"

func main() {
	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("perfect")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more!")
		}
	}
}

/*
	fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke satu case selanjutnya dengan tanpa
	menghiraukan nilai kondisinya, jadi satu case di pengecekan selanjutnya tersebut selalu dianggap benar
	(meskipun aslinya adalah salah). Dalam sebuah case switch lebih dari satu fallthrough
	bisa di tempatkan untuk memaksa melanjutkan proses pengecekan ke satu setelahnya.
*/
