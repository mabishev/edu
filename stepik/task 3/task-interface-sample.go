package main

import (
	"errors"
	"fmt"
)

func main() {

	lawyer := Lawyer{
		name:     "John",
		position: "senior law consultant",
	}

	name, err := GetName(lawyer)
	if err != nil {
		panic(err)
	}

	fmt.Println(name)

	stud := Student{
		name:   "adf",
		course: 1,
	}

	name, err = GetName(stud)
	if err != nil {
		panic(err)
	}

	fmt.Println(name)
}

type Named interface {
	Name() (string, error)
}

func GetName(n Named) (string, error) {
	return n.Name()
}

type Student struct {
	name   string
	course int
}

func (s Student) Name() (string, error) {
	if s.name == "" {
		return "", errors.New("student not has a name")
	}
	return s.name, nil
}

type Lawyer struct {
	name     string
	position string
}

func (l Lawyer) Name() (string, error) {
	if l.name == "" {
		return "", errors.New("lawyer not has a name")
	}
	return l.name, nil
}
