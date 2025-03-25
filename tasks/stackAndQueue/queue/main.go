package main

import "fmt"

type Queue[T int | float32 | string] struct {
	elements []T
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue удаляет элемент из начала очереди и возвращает его
func (q *Queue[T]) Dequeue() (element T, isOk bool) {
	if len(q.elements) == 0 {
		return
	}
	element = q.elements[0]
	q.elements = q.elements[1:]
	return
}

// Peek  позволяeт посмотреть на элемент в начале структуры без его удаления.
func (q *Queue[T]) Peek() (element T, isOk bool) {
	if len(q.elements) != 0 {
		return q.elements[0], true
	}
	return
}

func main() {
	queue := &Queue[float32]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Peek())    // Вывод: 1, true
	fmt.Println(queue.Dequeue()) // Вывод: 1, true
	fmt.Println(queue.Dequeue()) // Вывод: 2, true
	fmt.Println(queue.Dequeue()) // Вывод: 3, true
	fmt.Println(queue.Dequeue()) // Вывод: 0, false
	fmt.Println(queue.Peek())    // Вывод: 0, false
}
