package main

import (
	"container/list"
	"fmt"
)

// IntNode - узел списка
type IntNode struct {
	Value int
	Next  *IntNode //адрес на следующий элемент
	Prev  *IntNode //адрес на предыдущий элемент
}

// DoubleLinkedList - список
type DoubleLinkedList struct {
	size int
	head *IntNode //указатель на начало
	tail *IntNode //указатель на конец или хвост
}

// NewIntNode - конструктор типа IntNode
func NewIntNode2(value int) *IntNode {
	return &IntNode{0, nil, nil}
}

// NewDoubleLinkedList - конструктор списка
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{0, nil, nil}
}

func main() {
	list := list.New()

	el1 := list.PushFront(100)
	el2 := list.PushBack(500)
	el3 := list.InsertBefore(400, el2)
	el4 := list.InsertAfter(200, el1)
	for el := list.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
	fmt.Println("///////////////////////")
	for el := list.Back(); el != nil; el = el.Prev() {
		fmt.Printf("%d\n", el.Value)
	}
	fmt.Println(el4.Value, el3.Value)
	fmt.Println(list.Front().Value)
	fmt.Println(list.Back().Value)
	fmt.Println(list.Front().Next().Value)

}
