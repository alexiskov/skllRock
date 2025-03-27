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
	} else {
		list.Head.Previous = newNode
	}
	list.Head = newNode
}

// Печать всех элементов списка
func (list *LinkedList) PrintForward() {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Printf("%d ->previousVal-> ", node.Value)
		if node.Previous != nil {
			fmt.Println(node.Previous.Value)
		} else {
			fmt.Println(nil)
		}
	}
	fmt.Println("nil")
}

// Печать всех элементов списка
func (list *LinkedList) PrintBackward() {
	for node := list.Tail; node != nil; node = node.Previous {
		fmt.Printf("%d ->previousVal-> ", node.Value)
		if node.Next != nil {
			fmt.Println(node.Next.Value)
		} else {
			fmt.Println(nil)
		}
	}
	fmt.Println("nil")
}

func (list *LinkedList) Invert() {
	tempBuf := make([]*Node, 0)
	//наполняем буфер адресами нод
	for node := list.Head; node != nil; node = node.Next {
		tempBuf = append(tempBuf, node)
	}
	//перебираем буфер адресов нод  и меняем местами ссылки на соседние элементы в ноде
	for _, node := range tempBuf {
		node.Next, node.Previous = node.Previous, node.Next
	}
	//меняем местами адреса головы и хвоста
	list.Tail, list.Head = list.Head, list.Tail
}

func main() {
	list := &LinkedList{}
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	list.Invert()
	list.PrintForward() // 3 ->previousVal-> <nil>; 2 ->previousVal-> 3;	1 ->previousVal-> 2; nil

	fmt.Printf("list head val: %d\nlist tail val: %d\n", list.Head.Value, list.Tail.Value)

	list.PrintBackward() // 1 ->previousVal-> <nil>;	2 ->previousVal-> 1; 3 ->previousVal-> 2; nil
}
