package main

import "math"

func calculate(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	//hitung keliling
	var circumference = math.Pi * d
	return area, circumference
}
