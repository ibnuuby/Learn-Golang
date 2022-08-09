package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"john", "witch"}
	printMessage("Hello", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
