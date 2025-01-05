package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myRune := []rune(text)

	for i := 0; i < utf8.RuneCountInString(text)-1; i++ {
		fmt.Print((myRune[i] - '0') * (myRune[i] - '0'))
	}

}
