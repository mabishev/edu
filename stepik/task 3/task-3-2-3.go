package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	a := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, " ", ""), ",", "."), ";")

	conv, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		panic(err)
	}
	conv2, err := strconv.ParseFloat(a[1], 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.4f", conv/conv2)

}
