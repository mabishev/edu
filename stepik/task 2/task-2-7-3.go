package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var max int32

	for _, num := range text {
		if num >= max {
			max = num
		}
	}
	fmt.Println(string(max))

	/*  // old version
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myRune := []rune(text)
	max := '0'
	for i := 0; i < utf8.RuneCountInString(string(myRune)); i++ {
		if myRune[i] >= max {
			max = myRune[i]
		}
	}
	fmt.Println(string(max))
	*/
}
