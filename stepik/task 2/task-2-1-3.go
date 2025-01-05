package main

import "fmt"

func main() {
	a := vote(0, 1, 1)
	fmt.Println(a)
}

func vote(x int, y int, z int) int {

	if x == 0 && y == 0 {
		return 0
	}
	if x == 0 && z == 0 {
		return 0
	}
	if y == 0 && z == 0 {
		return 0
	}
	return 1
	/*
		if x == y {    // додумались же
	        return x
	    }
	    return z
	*/
}
