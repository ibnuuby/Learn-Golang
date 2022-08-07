package main

import "fmt"

func main() {
	/*
		Array adalah kumpulan data yang bertipe sama, yang disimpan pada sebuah variabel.

	*/
	var names [4]string
	names[0] = "satu"
	names[1] = "dua"
	names[2] = "tiga"
	names[3] = "empat"
	fmt.Println(names[0], names[1], names[2], names[3])
}
