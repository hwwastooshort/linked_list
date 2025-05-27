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

func (list *DoubleLinkedList[T]) RemoveFirst() (T, error) {

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

func (list *DoubleLinkedList[T]) RemoveLast() (T, error) {

	var errorItem T

	if list.tail == nil {
		return errorItem, errors.New("cannot remove last element of an empty list")
	}

	returnedValue := list.tail.value
	list.tail = list.tail.previous

	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}

	list.size--

	return returnedValue, nil

}

func (list *DoubleLinkedList[T]) Get(index int) (T, error) {
	var zeroValue T

	if index < 0 || index >= list.size {
		return zeroValue, errors.New("index out of bounds")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func (list *DoubleLinkedList[T]) InsertAt(index int, value T) (T, error) {
	if index < 0 || index > list.size {
		var zeroValue T
		return zeroValue, errors.New("index out of bounds")
	}

	if index == 0 {
		return list.InsertFirst(value), nil
	}
	if index == list.size {
		return list.InsertLast(value), nil
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	newNode := &DoubleNode[T]{value: value}

	prev := current.previous
	newNode.previous = prev
	newNode.next = current
	prev.next = newNode
	current.previous = newNode

	list.size++
	return value, nil
}

func (list *DoubleLinkedList[T]) RemoveAt(index int) (T, error) {
	var zeroValue T

	if index < 0 || index >= list.size {
		return zeroValue, errors.New("index out of bounds")
	}

	if index == 0 {
		return list.RemoveFirst()
	}
	if index == list.size-1 {
		return list.RemoveLast()
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	prev := current.previous
	next := current.next
	prev.next = next
	next.previous = prev

	list.size--
	return current.value, nil
}
