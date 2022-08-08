package main

import (
	"fmt"
)

func main() {

	/*
		Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan
		dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference,
		 menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang
		  memiliki alamat memori yang sama.
	*/
	var buah = []string{"apel", "mangga", "anggur", "jeruk"}
	fmt.Println(buah[2])
}
