package hashtable

import (
	"fmt"

	linkedlist "algorithms/linkedList"
)

type node[T comparable] struct {
	key  int
	data T
}

type HashTable[T comparable] struct {
	tableSize int
	table     []linkedlist.LinkedList[node[T]]
}

func (ht *HashTable[T]) hash(_key int) int {
	return _key % ht.tableSize
}

func New[T comparable](_size int) *HashTable[T] {
	return &HashTable[T]{
		tableSize: _size,
		table:     make([]linkedlist.LinkedList[node[T]], _size),
	}
}

func (ht *HashTable[T]) Insert(_key int, _value T) {
	index := ht.hash(_key)
	chain := &ht.table[index]
	newNode := node[T]{key: _key, data: _value}

	for element := range chain.Iterate() {
		if element.key == _key {
			element.data = _value
			return
		}
	}
	chain.Insert(newNode)
}

func (ht *HashTable[T]) Print() {
	for i := range ht.table {
		for element := range ht.table[i].Iterate() {
			fmt.Println(element.data)
		}
	}
}

func (ht *HashTable[T]) Search(_key int) (*T, bool) {
	index := ht.hash(_key)
	chain := &ht.table[index]

	for element := range chain.Iterate() {
		if element.key == _key {
			return &element.data, true
		}
	}
	return nil, false
}

func (ht *HashTable[T]) Delete(_key int) {
	index := ht.hash(_key)
	chain := &ht.table[index]

	for element := range chain.Iterate() {
		if element.key == _key {
			chain.Delete(element)
			return
		}
	}
}

func (ht *HashTable[T]) Update(_key int, _newValue T) {
	index := ht.hash(_key)
	chain := &ht.table[index]

	for element := range chain.Iterate() {
		if element.key == _key {
			chain.Update(element, node[T]{key: _key, data: _newValue})
			return
		}
	}
}

func (ht *HashTable[T]) Clear() {
	if ht.table == nil || ht.tableSize == 0 {
		return
	}

	for i := range ht.table {
		ht.table[i].Clear()
	}
}
