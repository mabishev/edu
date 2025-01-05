package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var rev int
	for a > 0 {
		x := a % 10
		a /= 10
		for b2 := b; b2 > 0; {
			y := b2 % 10
			b2 /= 10
			if x == y {
				rev = rev*10 + x
			}
		}
	}
	for rev > 0 {
		x := rev % 10
		rev = rev / 10
		fmt.Print(x, " ")
	}
}

package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	for _, ch1 := range a {
		for _, ch2 := range b {
			if ch1 == ch2 {
				fmt.Printf("%s ", string(ch1))
				break
			}
		}
	}
}
