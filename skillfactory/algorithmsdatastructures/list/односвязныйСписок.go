// Вот пример функции, которая находит элемент с указанным индексом и возвращает его значение, эту же функцию можно использовать и для чтения:
package main

import (
	"fmt"
)

func main() {
	list := NewIntList()
	list.Add(3)
	list.Add(2)
	list.Add(1)
	list.Insert(0, 0)
	list.Remove(3)
	list.Print()
	fmt.Println(list.Size2())

}

// ErrIndexOutOfBounds - описание ошибки при обращении к элементу с неверным индексом
var ErrIndexOutOfBounds = fmt.Errorf("Ошибка доступа за границы структуры")

// Node - узел списка
type Node struct {
	Value int
	next  *Node
}

// List - связный список элементов, хранящих целые числа
type List struct {
	size int
	head *Node
}

// Создание нового узла списка
func NewIntNode(value int) *Node {
	return &Node{value, nil}
}

// создание нового списка целых чисел
func NewIntList() *List {
	return &List{0, nil}
}

// Find - поиск узла по индексу
func (linkedList List) Find(index int) (*Node, error) {
	if index < 0 || index >= linkedList.size {
		return nil, ErrIndexOutOfBounds
	}
	var node *Node = linkedList.head
	for i := 1; i <= index; i++ {
		node = node.next
	}
	return node, nil
}

// обновление произвольного элемента списка
func (l *List) Set(el int, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrIndexOutOfBounds
	}
	node, err := l.Find(index)
	if err != nil {
		return err
	}
	node.Value = el
	return nil
}

// получение размера списка
func (l List) Size2() int {
	return l.size
}

// Напишем примерный код функции определения размера списка:
func (linkedList List) Size() int {
	var size int = 0
	if linkedList.head != nil {
		var node *Node = linkedList.head.next
		for size = 1; node.next != nil; size++ {
			node = node.next
		}
	}
	return size
}

// Приведём пример функции, удаляющей произвольный элемент с индексом i:
func (linkedList *List) Remove(index int) (*Node, error) {
	removeNode, err := linkedList.Find(index)
	if err != nil {
		return nil, err
	}
	if index == 0 {
		linkedList.head = linkedList.head.next
	} else {
		prevNode, err := linkedList.Find(index - 1)
		if err != nil {
			return nil, err
		}
		prevNode.next = removeNode.next
	}
	linkedList.size--
	return removeNode, nil
}

// Добавление нового элемента в начало списка
func (l *List) Add(el int) {
	newNode := NewIntNode(el)
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

// вставка нвоого элемента в произвольную позицию
func (l *List) Insert(el int, index int) error {
	if index < 0 || index >= l.size {
		return ErrIndexOutOfBounds
	}
	newNode := NewIntNode(el)
	if index == 0 {
		l.Add(el)
		return nil
	}
	node, err := l.Find(index - 1)
	if err != nil {
		return err
	}
	newNode.next = node.next
	node.next = newNode
	l.size++
	return nil
}

// печать списка
func (l List) Print() {
	node := l.head
	if node != nil {
		for node != nil {
			fmt.Printf("%d\t", node.Value)
			node = node.next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Список пуст!")
	}
}
