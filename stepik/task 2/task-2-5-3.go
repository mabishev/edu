package main

import (
	"fmt"
	"strings"
)

func main() {
	str1, str2 := "", "" // а мог так: var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(strings.Index(str1, str2))
}
