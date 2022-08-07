package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "anggur", "mangga", "jeruk"}

	for i, buah := range buah {
		fmt.Println("Element :", i, buah)
	}
}
