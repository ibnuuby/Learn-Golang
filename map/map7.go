package main

import "fmt"

func main() {
	var chicken = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chicken {
		fmt.Println(chicken["name"], chicken["gender"])
	}
}
