package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
easy solution

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(text)
	if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

*/
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myRune := []rune(text)
	for i := 0; i < utf8.RuneCountInString(string(myRune))-1; i++ {
		if unicode.IsUpper(myRune[0]) && myRune[utf8.RuneCountInString(text)-2] == '.' {
			fmt.Println("Right")
			break
		} else {
			fmt.Println("Wrong")
			break
		}
	}
}
