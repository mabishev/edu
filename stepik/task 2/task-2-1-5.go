package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(a ...int) (int, int) {
	var b int
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	return len(a), b
}

/*                                  //perfect, but not mine
func sumInt(n ...int ) (c, s int) {
    for _, e := range n {
        c++
        s+=e
    }
    return
}
*/
