package main

import "fmt"

func main() {
	var chicken = map[string]int{"januari": 50, "februari": 40}

	fmt.Println("januari", chicken["januari"])
	fmt.Println("februari", chicken["februari"])

}
