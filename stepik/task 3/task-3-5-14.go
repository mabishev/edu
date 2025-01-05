package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	args := os.Args[1:]
	if len(args) != 1 {
		return errors.New("file name is not specified")
	}

	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	position, err := FindZeroPosition(file)
	if err != nil {
		return err
	}

	fmt.Println(position)
	return nil
}

func FindZeroPosition(src io.Reader) (int, error) {
	reader := bufio.NewReader(src)

	var readCount int
	for {
		currentValue, err := reader.ReadString(';')
		if err == io.EOF {
			return 0, errors.New("zero not found")
		}
		if err != nil {
			return 0, err
		}
		readCount++
		if currentValue == "0;" {
			break
		}
	}

	return readCount, nil
}
