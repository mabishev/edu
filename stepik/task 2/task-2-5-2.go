package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	/*            //beautiful, but not mine
	 str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	    text := []rune(str)
	    for i, j := range text{
	        if text[len(text)-i-1] != j{
	            fmt.Print("Нет")
	            return
	        }
	    }
	    fmt.Print("Палиндром")
	}
	*/
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var revStr string
	myRune := []rune(text)

	for i := utf8.RuneCountInString(string(myRune)) - 2; i >= 0; i-- {
		revStr = revStr + string(myRune[i])
	}

	if revStr == text {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
