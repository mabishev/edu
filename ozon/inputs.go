package edu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Inputs() {
	n, err := GetFloat()
	if err != nil {
		fmt.Println(n)
		return
	}
	fmt.Println(n)

	Scanner()
}

func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Println(scanner.Text())

	/*
	   scanner := bufio.NewScanner(os.Stdin)
	   	//scanner.Split(bufio.ScanWords) // без Split программа завершается с помощью enter
	   	for scanner.Scan() {
	   		number, err := strconv.ParseFloat(scanner.Text(), 64)
	   		if err != nil {
	   			fmt.Println(err)
	   			return
	   		}
	   		fmt.Println(number)
	   	}
	*/
}

func GetFloat() (float64, error) {
	reader1 := bufio.NewReader(os.Stdin)
	input, err := reader1.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
