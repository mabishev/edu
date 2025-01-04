package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	start := time.Now()

	s := scanner.Text()

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println(elapsed)
	fmt.Println(RLE(s))
}

// находит количество входов букв
func RLE(s string) string {
	r := []rune(s)
	nowWord := r[0]
	c := 1
	//total := ""
	var builder strings.Builder
	for i := 1; i < len(r); i++ {
		if r[i] == nowWord {
			c++
		} else {
			builder.WriteString(string(nowWord)) // переменная builder аккумулирует строку
			builder.WriteString(strconv.Itoa(c))
			//total += string(nowWord) + strconv.Itoa(c) //Использование оператора += к строке
			//в цикле может привести к созданию большого числа промежуточных строк, что может
			//сказаться на производительности. Чтобы это избежать, можно использовать
			//strings.Builder, который оптимизирован для конкатенации строк"
			nowWord = r[i]
			c = 1
		}
	}
	builder.WriteString(string(nowWord))
	builder.WriteString(strconv.Itoa(c))
	return builder.String()

	//total += string(nowWord) + strconv.Itoa(c)
	//return total
}
