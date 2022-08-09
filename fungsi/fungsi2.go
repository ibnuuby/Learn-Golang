package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	rand.Seed() digunakan untuk penentuan nilai seed.
	Fungsi rand.Int() digunakan untuk generate angka random dalam bentuk numerik bertipe int.
*/

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
