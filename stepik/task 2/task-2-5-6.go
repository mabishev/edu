package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myRune := []rune(text)
	out := "Ok"
	/*   //you can use len like that
	var a string
	fmt.Scan(&a)
	len([]rune(a))
	*/
	if utf8.RuneCountInString(text)-1 > 4 {
		for i := 0; i < utf8.RuneCountInString(text)-1; i++ {
			if !unicode.Is(unicode.Latin, myRune[i]) && !unicode.Is(unicode.Digit, myRune[i]) {
				out = "Wrong password"
				break
			}
		}
	} else {
		out = "Wrong password"
	}
	fmt.Println(out)

	/*  // old version
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myRune := []rune(text)
	var check bool

	if utf8.RuneCountInString(text)-1 > 4 {
		for i := 0; i < utf8.RuneCountInString(text)-1; i++ {
			if !unicode.Is(unicode.Latin, myRune[i]) && !unicode.Is(unicode.Digit, myRune[i]) {
				check = true
				break
			}
		}
	} else {
		check = true
	}

	if check == false {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
	*/
}
