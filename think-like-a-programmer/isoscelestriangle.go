package edu

import (
	"fmt"
	"math"
)

// равнобедренный треугольник
// #
// ##
// ###
// ####
// ###
// ##
// #
func IsoscelesTriangle() {
	for row := 1; row <= 7; row++ {
		for i := 1; i <= 4-int(math.Abs(4-float64(row))); i++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
