package main

type Stack[T []int | []uint | []float32 | []rune | []byte] struct {
	elements []T
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func main() {
	//stackInt := StackInt{}
}
