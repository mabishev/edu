package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}
	if info.IsDir() {
		return nil // Проигнорируем директории
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	if len(rows) != 10 {
		return nil
	}

	if len(rows[4]) != 10 {
		return nil
	}

	fmt.Println(rows[4][2])
	return nil
}

func main() {

	const root = "./task" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}
