package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"watermelon", "pinaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
