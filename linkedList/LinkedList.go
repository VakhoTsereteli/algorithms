package linkedlist

import "fmt"

type node[T comparable] struct {
	data T
	next *node[T]
}

type LinkedList[T comparable] struct {
	head *node[T]
	last *node[T]
	size int
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Insert(value T) {
	newNode := &node[T]{data: value, next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		ll.last.next = newNode
	}

	ll.last = newNode
	ll.size++
}

func (ll *LinkedList[T]) Print() {
	current := ll.head

	if ll.head == nil {
		return
	}

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (ll *LinkedList[T]) Delete(value T) {
	if ll.head == nil {
		return
	}

	if ll.head.data == value {
		ll.head = ll.head.next

		if ll.head == nil {
			ll.last = nil
		}
		return
	}

	prev := ll.head
	current := ll.head.next

	for current != nil {
		if current.data == value {
			prev.next = current.next

			if current.next == nil {
				ll.last = prev
			}
			break
		}

		prev = current
		current = current.next
	}
	ll.size--
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList[T]) Update(element T, value T) {
	current := ll.head

	for current != nil {
		if current.data == element {
			current.data = value
			break
		}

		current = current.next

	}
}

func (ll *LinkedList[T]) Search(value T) int {
	current := ll.head
	index := 0

	for current != nil {
		if current.data == value {
			return index
		}
		current = current.next
		index += 1
	}

	return -1
}

func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.last = nil
	ll.size = 0
}

func (ll *LinkedList[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		current := ll.head
		for current != nil {
			ch <- current.data
			current = current.next
		}
		close(ch)
	}()
	return ch
}

func (ll *LinkedList[T]) Find(_value T) bool {
	current := ll.head

	for current != nil {
		if current.data == _value {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList[T]) GetHead() *node[T] {
	return ll.head
}
