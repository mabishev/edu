package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(n uint) uint {
		str := strconv.FormatUint(uint64(n), 10)
		apstr := []rune{}
		for _, elem := range str {

			if elem%2 == 0 && elem != '0' {
				apstr = append(apstr, elem)
			}
		}

		bar, err := strconv.ParseUint(string(apstr), 10, 64)
		if bar == 0 {
			return 100
		}
		if err != nil {
			panic(err)
		}

		return uint(bar)
	}
	fmt.Println(fn(713))

}
