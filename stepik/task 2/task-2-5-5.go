package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	/* easy code but not mine

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	for _, s := range text {
		if strings.Count(text, string(s)) < 2 {
			fmt.Print(string(s))
		}
	}*/

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	for i := 0; i < utf8.RuneCountInString(text)-1; i++ {
		if strings.Count(text, string(text[i])) == 1 {
			fmt.Print(string(text[i]))
		}
	}
}
