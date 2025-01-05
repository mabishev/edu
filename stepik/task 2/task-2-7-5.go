package main

import "math"

var p, v, k float64

func M(x, y float64) float64 {
	return x * y
}

func W(w float64) float64 {
	c := w / M(p, v)
	return math.Sqrt(c)
}

func T() float64 {
	t := 6 / W(k)
	return t
}

func main() {

}
