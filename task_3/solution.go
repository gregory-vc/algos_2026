package task3

import "reflect"

type Node[T any] struct {
	value      T
	prev, next *Node[T]
	list       *LinkedList[T]
}

type LinkedList[T any] struct {
	first, last *Node[T]
	size        int
}

func (list *LinkedList[T]) AddFirst(value T) *Node[T] {
	node := &Node[T]{value: value, next: list.first, list: list}
	if list.first != nil {
		list.first.prev = node
	} else {
		list.last = node
	}
	list.first = node
	list.size++
	return node
}

func (list *LinkedList[T]) AddLast(value T) *Node[T] {
	node := &Node[T]{value: value, prev: list.last, list: list}
	if list.last != nil {
		list.last.next = node
	} else {
		list.first = node
	}
	list.last = node
	list.size++
	return node
}

func (list *LinkedList[T]) InsertAfter(position *Node[T], value T) *Node[T] {
	list.check(position)
	node := &Node[T]{value: value, prev: position, next: position.next, list: list}
	if position.next != nil {
		position.next.prev = node
	} else {
		list.last = node
	}
	position.next = node
	list.size++
	return node
}

func (list *LinkedList[T]) Contains(value T) bool {
	return list.IndexOf(value) != -1
}

func (list *LinkedList[T]) IndexOf(value T) int {
	index := 0
	for node := list.first; node != nil; node = node.next {
		if reflect.DeepEqual(node.value, value) {
			return index
		}
		index++
	}
	return -1
}

func (list *LinkedList[T]) RemoveFirst() {
	if list.first != nil {
		list.Remove(list.first)
	}
}

func (list *LinkedList[T]) RemoveLast() {
	if list.last != nil {
		list.Remove(list.last)
	}
}

func (list *LinkedList[T]) Remove(position *Node[T]) {
	list.check(position)
	if position.prev != nil {
		position.prev.next = position.next
	} else {
		list.first = position.next
	}
	if position.next != nil {
		position.next.prev = position.prev
	} else {
		list.last = position.prev
	}
	position.prev, position.next, position.list = nil, nil, nil
	list.size--
}

func (list *LinkedList[T]) Get(position *Node[T]) T {
	list.check(position)
	return position.value
}

func (list *LinkedList[T]) Set(position *Node[T], value T) {
	list.check(position)
	position.value = value
}

func (list *LinkedList[T]) Len() int {
	return list.size
}

func (list *LinkedList[T]) check(position *Node[T]) {
	if position == nil || position.list != list {
		panic("position does not belong to list")
	}
}
