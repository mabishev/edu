package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		text := scanner.Text()

		n, err := strconv.Atoi(text)
		if n < 0 || n > 100 {
			break
		}
		if err != nil {
			panic(err)
		}
		sum += n
	}
	os.Stdout.WriteString(strconv.Itoa(sum))
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
}
