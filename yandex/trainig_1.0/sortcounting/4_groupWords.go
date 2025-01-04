package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Задача: Сгруппировать слова по общим буквам. Напр: input: ["eat", "tea", "ate", "tan, "bat", "nat"].
//Output: [["ate", "tea", "eat"], ["nat", "tan"], ["bat"]]

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var words []string
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		words = append(words, input)
	}

	result := groupWords(words)
	fmt.Println(result)
}

func groupWords(s []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range s {
		wordSlice := strings.Split(word, "")
		sort.Strings(wordSlice)
		sortedWord := strings.Join(wordSlice, "")
		groups[sortedWord] = append(groups[sortedWord], word)
	}

	ans := make([][]string, 0, len(groups))
	for key := range groups {
		ans = append(ans, groups[key])
	}
	return ans
}
