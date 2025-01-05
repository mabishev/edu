package main

import (
	"strconv"
)

// short
/*func adding() {
	a := "%^-80"
	b := "hhhhh+20&&&&nd"

	a = strings.TrimFunc(a, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	b = strings.TrimFunc(b, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	a1, _ := strconv.ParseInt(a, 10, 64)
	b1, _ := strconv.ParseInt(b, 10, 64)
	fmt.Println(a1 + b1)

}
*/

func adding(x, y string) int64 {
	c := []byte("")
	for _, num := range x {
		if num > 47 && num < 58 {
			c = append(c, byte(num))
		}
	}
	d := []byte("")
	for _, num := range y {
		if num > 47 && num < 58 {
			d = append(d, byte(num))
		}
	}

	z, err := strconv.ParseInt(string(c), 10, 64)
	if err != nil {
		panic(err)
	}
	q, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		panic(err)
	}
	return z + q
}
