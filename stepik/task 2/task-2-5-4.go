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
	my2r := []rune("")

	for i := 1; i < utf8.RuneCountInString(string(myRune))-1; i += 2 {
		my2r = append(my2r, myRune[i])
	}

	fmt.Println(string(my2r))
}

/*    // which right and easier?
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myRune := []rune(text)

	for i := 1; i < utf8.RuneCountInString(string(myRune)); i += 2 {
		fmt.Print(string(myRune[i]))
	}

}
*/
