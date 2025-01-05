package main

import "fmt"

func main() {
	var n int
	var cow string
	fmt.Scan(&n)
	if n%10 == 1 && n != 11 {
		cow = "korova"
	} else if n%10 == 2 && n != 12 || n%10 == 3 && n != 13 || n%10 == 4 && n != 14 {
		cow = "korovy"
	} else {
		cow = "korov"
	}
	fmt.Printf("%d %s\n", n, cow)
}
