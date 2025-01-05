package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
type BatteryForTest string

func (b BatteryForTest) String() string{
    col := strings.Count(string(b), "1")
    return fmt.Sprintf("[%10s]", strings.Repeat("X", col))
}
func main() {
    var batteryForTest BatteryForTest
    fmt.Scan(&batteryForTest)
	fmt.Print(batteryForTest)
*/

type Battery struct {
	Number string
}

func (s Battery) String() string {
	return fmt.Sprintf("[%v]", s.Number)
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")

	r1 := []rune{}
	var count int
	for _, v := range text {
		if v == '0' {
			r1 = append(r1, ' ')
		}
		if v == '1' {
			count++
		}
	}
	for i := 0; i < count; i++ {
		r1 = append(r1, 'X')
	}

	batteryForTest := Battery{
		Number: string(r1),
	}

	fmt.Println(batteryForTest)

}
