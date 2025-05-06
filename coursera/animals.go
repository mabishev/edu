package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animal := make(map[string]Animal)
	for {
		fmt.Println(`> Enter "newanimal" or "query"`)
		input := inputStr()

		switch input {
		case "newanimal":
			fmt.Print("Enter the name of the new animal: ")
			nameOfAnimal := inputStr()

			fmt.Println(`Enter type of the new animal, either “cow”, “bird”, or “snake"`)

			typeOfAnimal := inputStr()

			err := whichAnimal(typeOfAnimal, nameOfAnimal, animal)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Created it!")

		case "query":
			fmt.Println("Enter the name of animal")
			nameOfAnimal := inputStr()

			if _, ok := animal[nameOfAnimal]; !ok {
				fmt.Println("animal not found")
				continue
			}

			fmt.Println(`Enter the name of the information requested about the animal, either “eat”, “move”, or “speak”`)
			reqInfo(inputStr(), nameOfAnimal, animal)

		default:
			fmt.Println("wrong input")
		}
	}

}

func inputStr() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func reqInfo(info, nameOfAnimal string, animal map[string]Animal) {
	if val, ok := animal[nameOfAnimal]; ok {
		switch info {
		case "eat":
			val.Eat()
		case "move":
			val.Move()
		case "speak":
			val.Speak()
		default:
			fmt.Println("wrong request")
		}
	}
}

func whichAnimal(typeOfAnimal, nameOfAnimal string, animal map[string]Animal) error {
	switch typeOfAnimal {
	case "cow":
		animal[nameOfAnimal] = &Cow{}
	case "bird":
		animal[nameOfAnimal] = &Bird{}
	case "snake":
		animal[nameOfAnimal] = &Snake{}
	default:
		return errors.New("wrong type of animal")
	}
	return nil
}
