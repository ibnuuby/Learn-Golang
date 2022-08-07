package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "mangga", "anggur", "lemon"}

	for i := 0; i < len(buah); i++ {
		fmt.Println("Element : ", i, buah[i])
	}
}
