package main

import (
	"fmt"
	"time"
)

func main() {

	var s string
	fmt.Scanln(&s)
	firstTime, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format("Mon Jan _2 15:04:05 MST 2006"))

}
