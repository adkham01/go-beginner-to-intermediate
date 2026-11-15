package main

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		return *new(T), false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	a, b := 1, 2
	println("Before swap:", a, b)
	a, b = swap(a, b)
	println("After swap:", a, b)

	x, y := "hello", "world"
	println("Before swap:", x, y)
	x, y = swap(x, y)
	println("After swap:", x, y)

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	for !intStack.IsEmpty() {
		element, _ := intStack.Pop()
		println("Popped from intStack:", element)
	}
}
