package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*
func main() {
	row, _ := csv.NewReader(os.Stdin).Read()
	x, _ := time.Parse("02.01.2006 15:04:05", row[0])
	y, _ := time.Parse("02.01.2006 15:04:05", row[1])
	diff := x.Sub(y)
	if diff < 0 {
		diff *= -1
	}
	fmt.Println(diff)
}
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputs := strings.Split(input, ",")
	s1, s2 := inputs[0], inputs[1]

	firstTime, err := time.Parse("02.01.2006 15:04:05", s1)
	if err != nil && err != io.EOF {
		panic(err)
	}
	secTime, err := time.Parse("02.01.2006 15:04:05", s2)
	if err != nil && err != io.EOF {
		panic(err)
	}

	if firstTime.Sub(secTime) > 0 {
		fmt.Println(firstTime.Sub(secTime))
	} else {
		fmt.Println(secTime.Sub(firstTime))
	}
}
