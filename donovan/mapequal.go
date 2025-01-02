package edu

import "fmt"

// Чтобы проверить, содержат ли два отображения одни и те же ключи и связанные с ними значения
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

func EqualMap() {
	s := map[string]int{"hello": 2, "dolly": 5}
	s2 := map[string]int{"hello": 2, "dolly": 5}
	fmt.Println(equal(s, s2))
}
