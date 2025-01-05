package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	firstTime, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		panic(err)
	}
	if firstTime.Hour() > 13 {
		new_dt := firstTime.AddDate(0, 0, 1)
		fmt.Print(new_dt.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
	}
}
