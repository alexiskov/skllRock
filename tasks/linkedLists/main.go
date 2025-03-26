package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

// Добавление элемента в начало списка
func (list *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: list.Head}
	if list.Head == nil && list.Tail == nil {
		list.Tail = newNode
	}
	list.Head = newNode
}

// Печать всех элементов списка
func (list *LinkedList) Print() {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	list.Print() // Вывод: 3 -> 2 -> 1 -> nil
}
