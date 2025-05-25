package linkedlist

import (
	"errors"
	"fmt"
)

type DoubleNode[T any] struct {
	value    T
	previous *DoubleNode[T]
	next     *DoubleNode[T]
}

type DoubleLinkedList[T any] struct {
	head *DoubleNode[T]
	tail *DoubleNode[T]
	size int
}

func (list *DoubleLinkedList[T]) Size() int {
	return list.size
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (list *DoubleLinkedList[T]) String() string {
	if list.head == nil {
		return "empty"
	}

	result := ""
	current := list.head

	for current != nil {
		prevStr := "nil"
		nextStr := "nil"

		if current.previous != nil {
			prevStr = fmt.Sprintf("%v", current.previous.value)
		}
		if current.next != nil {
			nextStr = fmt.Sprintf("%v", current.next.value)
		}

		result += fmt.Sprintf("[prev: %v <- val: %v -> next: %v]", prevStr, current.value, nextStr)

		if current.next != nil {
			result += " <=> "
		}

		current = current.next
	}

	return result
}

func (list *DoubleLinkedList[T]) InsertFirst(newValue T) T {

	newNode := &DoubleNode[T]{value: newValue}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head     // old head becomes successor of new one
		list.head.previous = newNode // new head becomes predecessor of old one
		list.head = newNode
	}

	list.size++
	return newValue
}

func (list *DoubleLinkedList[T]) InsertLast(newValue T) T {

	newNode := &DoubleNode[T]{value: newValue}

	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.previous = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
	return newValue

}

func (list *DoubleLinkedList[T]) removeFirst() (T, error) {
	var errorItem T

	if list.head == nil {
		return errorItem, errors.New("cannot remove first element of an empty list")
	}

	returnedValue := list.head.value
	list.head = list.head.next

	if list.head != nil { // the head has already been moved, not redundant
		list.head.previous = nil
	} else {
		list.tail = nil
	}

	list.size--

	return returnedValue, nil
}
