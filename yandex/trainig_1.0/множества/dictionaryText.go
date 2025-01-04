package main

import (
	"fmt"
)

func main() {
	dictionary := []string{"apple", "banana"}
	text := []string{"apple", "anan"}
	fmt.Println(wordsIndict(dictionary, text))

}

// Дан словарь из N слов,длина каждого не превосходит К. В записи каждого из М слов текста(каждое длиной до К)
// может быть пропущена одна буква. Для каждого слова сказать, входит ли оно(возможно, с одной пропущенной
// буквой), в словарь. Сложность O(NK+M)
func wordsIndict(dictionary, text []string) []bool {
	goodWords := make(map[string]bool)
	for _, word := range dictionary {
		goodWords[word] = true
		for delPos := range word {
			deletedWord := word[:delPos] + word[delPos+1:] //[pple:true aple:true appe:true appl:true]
			goodWords[deletedWord] = true
		}
	}
	//fmt.Println(goodWords) // [apple:true pple:true aple:true appe:true appl:true]

	var ans []bool
	for _, word := range text {
		ans = append(ans, goodWords[word])
	}
	return ans
}
