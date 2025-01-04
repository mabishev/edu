// односвязный список
package main

import "fmt"

type LinkNode struct {
	next  *LinkNode
	value int
}

// вывод через курсор
func (n *LinkNode) Print() {
	cur := n
	for cur != nil {
		splitter := ""
		if cur != n {
			splitter = " -> "
		}
		fmt.Printf("%s%d", splitter, cur.value)
		cur = cur.next
	}
}
func (n *LinkNode) Reverse() *LinkNode {
	var cur = n
	var prev *LinkNode
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	n1 := &LinkNode{value: 4}
	n2 := &LinkNode{value: 10}
	n3 := &LinkNode{value: 20}

	n1.next = n2
	n2.next = n3

	n1.Print()
	fmt.Println("")
	n1.Reverse().Print()

}
