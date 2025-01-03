package edu

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Comma() {
	s := "1234567"
	fmt.Println(comma(s))
	fmt.Println(commaBytes(s))
}

// рекурсия
func comma(s string) string {
	n := utf8.RuneCountInString(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]

}

// без конкатенации строк
func commaBytes(s string) string {
	var buf bytes.Buffer
	n := utf8.RuneCountInString(s)
	if n <= 3 {
		return s
	}

	firstCommaIndex := n % 3
	if firstCommaIndex == 0 {
		firstCommaIndex = 3
	}

	buf.WriteString(s[:firstCommaIndex])
	s = s[firstCommaIndex:]

	for len(s) > 0 {
		buf.WriteRune(',')
		buf.WriteString(s[:3])
		s = s[3:]
	}
	return buf.String()
}
