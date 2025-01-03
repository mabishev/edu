package edu

import (
	"fmt"
	"strconv"
)

func SprintfItoa() {
	x := 123
	y := fmt.Sprintf("%d", x) //аналог itoa
	fmt.Println(y)
	fmt.Println(strconv.Itoa(x))
	s := fmt.Sprintf("x=%b", x) // в двоичной системе исчисления, %d(десятичной), %o(восьмиричной), %x(шестнадцетиричной)
	fmt.Println(s)
}
