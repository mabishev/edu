package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	// var m, s time.Duration  // NOT MINE SOLUTION
	// fmt.Scanf("%d мин. %d сек.", &m, &s)
	// t := time.Unix(1589570165, 0).UTC().Add(m * time.Minute).Add(s * time.Second)
	// fmt.Println(t.Format(time.UnixDate))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	rep := strings.Replace(input, "мин.", "m", 1)
	rep = strings.Replace(rep, "сек.", "s", 1)
	tr := strings.ReplaceAll(rep, " ", "")

	dur, err := time.ParseDuration(tr)
	if err != nil {
		panic(err)
	}
	res := dur.Round(time.Nanosecond).Nanoseconds()
	unixTime := time.Unix(
		now, // секунды
		res, // наносекунды
	).UTC()

	fmt.Println(unixTime.Format(time.UnixDate))
}
