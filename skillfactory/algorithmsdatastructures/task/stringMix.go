package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputStr := "ыыааррыыффррф"
	str2 := []string{}
	for _, v := range inputStr {
		str2 = append(str2, string(v))
	}
	sort.Strings(str2)
	s2 := ""
	for _, v := range str2 {
		s2 += v
	}
	fmt.Println(s2)

	str := ""
	var x int
	for _, v := range s2 {
		if str != "" && strings.Contains(str, string(v)) {
			continue
		}
		x = strings.Count(s2, string(v))
		s := strconv.Itoa(x)
		str += string(v) + s

	}

	fmt.Println(str)
}
