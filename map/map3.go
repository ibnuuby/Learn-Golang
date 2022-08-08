package main

import (
	"fmt"
)

func main() {
	var chicken1 = map[string]int{}
	var chicken2 = make(map[string]int)
	var chicken3 = *new(map[string]int)

	fmt.Println("chicken1: ", chicken1["chicken1"])
	fmt.Println("chicken2: ", chicken2["chicken2"])
	fmt.Println("chicken3: ", chicken3["chicken3"])

}
