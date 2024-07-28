// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.

package main

import (
	"log"
)

func GetAllKeys[K comparable, V string](m map[K]V) []K {
	tmpMap := make([]K, 0, len(m))
	for k := range m {
		tmpMap = append(tmpMap, k)
	}
	return tmpMap
}

type Bio struct {
	name string
	age  int
}
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
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

func (lst *List[T]) GetByIndex(index int) T {
	elems := lst.GetAll()
	if len(elems) > index-1 {
		log.Fatal("Index out of error")
		var t T
		return t
	}
	return lst.GetAll()[index]
}

type ListInt[T any] struct {
	next *ListInt[T]
	val  T
}

func main() {
	// ex1 := map[string]string{"name": "Budi", "age": "18"}
	// fmt.Println(GetAllKeys(ex1))
	// lst := List[int]{}
	// lst.Push(10)
	// fmt.Println(lst.GetAll())
	// fmt.Println(lst.GetByIndex(1))

	// lst := new(ListInt[int])
	// for i := 0; i < 5; i++ {
	// 	lst.next = lst.val{1: 1}
	// }

	// fmt.Println(&lst.val)

}
