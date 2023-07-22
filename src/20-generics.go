package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{ val: v }
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{ val: v }
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{ 1: "A", 2: "B", 3: "C" }
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(42)
	lst.Push(1337)
	fmt.Println("list:", lst.GetAll())
}
