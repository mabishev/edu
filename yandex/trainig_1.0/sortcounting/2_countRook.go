package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Rook struct {
	Row int
	Col int
}

// Задача: на шахматной доске N x N находится М ладей(ладья бьет клетки на той же горизонтали или вертикали
// до ближайшей занятой). Определите сколько пар ладей бьют друг друга. Ладьи задаются парой чисел J и K,
// обозначающих координаты клетки. Сложность О(М)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	var rookCoords []Rook // Создаем срез для хранения координат ладей
	for i := 0; i < m; i++ {
		scanner.Scan()
		j, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		// Создаем экземпляр структуры Rook и добавляем его в срез
		rook := Rook{Row: j, Col: k}
		rookCoords = append(rookCoords, rook)
	}
	result := countBeatingRooks(rookCoords)
	fmt.Println(result)
}

func countBeatingRooks(rookCoords []Rook) int {
	rooksInRow := make(map[int]int)
	rooksInCol := make(map[int]int)
	//Функция для добавления ладьи в строку или столбец
	addRock := func(rowOrCol map[int]int, key int) {
		rowOrCol[key]++
	}
	// Функция для подсчета пар в строке или столбце
	countPairs := func(rowOrCol map[int]int) int {
		pairs := 0
		for _, count := range rowOrCol {
			if count == 1 {
				continue
			}
			pairs += count - 1
		}
		return pairs
	}
	// Count rooks in rows and columns
	for _, coords := range rookCoords {
		row, col := coords.Row, coords.Col
		addRock(rooksInRow, row)
		addRock(rooksInCol, col)
	}
	return countPairs(rooksInRow) + countPairs(rooksInCol)
}
