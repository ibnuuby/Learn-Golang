package main

import "fmt"

func main() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(len(chicken))
	fmt.Println(chicken)

	/*
		Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada
		variabel map. Cara penggunaannya, dengan memasukan objek map dan key item yang
		ingin dihapus sebagai parameter.
	*/
	delete(chicken, "januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)
}
