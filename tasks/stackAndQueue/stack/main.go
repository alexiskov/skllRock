// изначально пример был только для типа: int
// задача: Модифицируйте стек и _очередь, добавив функциональность для обработки разных типов данных (не только int).
// perfect!!! очевидно необходимы дженерики) пальчики мои не казенные))
package main

import (
	"fmt"
)

// определим ограничения(constraints) для используемых типов структур и методов
type Stack[T int | uint | float32 | rune | byte | string] struct {
	elements []T
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack[T]) Pop() (element T, isOk bool) {
	if len(s.elements) == 0 {
		return
	}
	index := len(s.elements) - 1
	element = s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

// показывает первый элемент стека
func (s *Stack[T]) Peek() (element T, isOk bool) {
	if len(s.elements) != 0 {
		return s.elements[0], true
	}
	return
}

func main() {
	//пусть в данный момент это будет строчка
	stack := Stack[string]{}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	fmt.Println(stack.Peek()) // Вывод: a, true

	fmt.Println(stack.Pop()) // Вывод: c, true
	fmt.Println(stack.Pop()) // Вывод: b, true
	fmt.Println(stack.Pop()) // Вывод: a, true
	fmt.Println(stack.Pop()) // Вывод:  , false

	fmt.Println(stack.Peek()) // Вывод:  , false

}
