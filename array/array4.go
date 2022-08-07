package main

import "fmt"

func main() {
	var number1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var number2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("number1 :", number1)
	fmt.Println("number2 :", number2)
}
