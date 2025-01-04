package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // aa b cc d
	scanner.Scan()

	start := time.Now()

	s := strings.Split(scanner.Text(), " ")

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
	fmt.Println(shortWords(s))
}

// находит короткие слова
func shortWords(arr []string) string {
	find := make([]string, 0, len(arr))

	for i := 0; i < len(arr); i++ { // исключаем цифры
		var flag bool
		for _, v := range arr[i] {
			if unicode.IsDigit(v) {
				flag = true
				break
			}
		}
		if flag {
			copy(arr[i:], arr[i+1:])
			arr = arr[:len(arr)-1]
			i--
		}
	}

	// Вариант №2. Если вы не хотите изменять исходный слайс и хотите сохранить его в неизменном виде, то
	//второй вариант (с созданием нового слайса) предпочтительнее. Если использование памяти критично, и
	//вы готовы изменить исходный слайс, то первый вариант (с изменением базового слайса) может быть более эффективным.
	//arr2 := []string{}
	//for _, val := range arr { // исключаем цифры
	//	var flag bool
	//	for _, v := range val {
	//		if unicode.IsDigit(v) {
	//			flag = true
	//		}
	//	}
	//	if !flag {
	//		arr2 = append(arr2, val)
	//	}
	//}

	minLenWords := utf8.RuneCountInString(arr[0])

	for _, val := range arr {
		if utf8.RuneCountInString(val) < minLenWords {
			minLenWords = utf8.RuneCountInString(val)
		}
	}

	for _, val := range arr {
		if utf8.RuneCountInString(val) == minLenWords {
			find = append(find, val)
		}
	}
	return strings.Join(find, " ") // создает одну строку из массива строк, разделяя пробелом
}
