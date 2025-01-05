package main

import (
	"fmt"
	"math"
)

/*  // short variant
var a, b float64
fmt.Scan(&a, &b)
fmt.Print(math.Hypot(a, b))
*/

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	c := a*a + b*b
	fmt.Println(math.Sqrt(c))
}
