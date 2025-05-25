package linkedlist

import (
	"errors"
	"fmt"
)

type SingleNode[T any] struct {
	value T
	next  *SingleNode[T]
}

type SingleLinkedList[T any] struct {
	head *SingleNode[T]
	size int
}

func (list *SingleLinkedList[T]) Size() int {
	return list.size
}

func NewSingleLinkedList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (list *SingleLinkedList[T]) InsertLast(newValue T) T {

	newNode := &SingleNode[T]{value: newValue}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

	list.size++
	return newValue

}

func (list *SingleLinkedList[T]) InsertFirst(newVaule T) T {

	if list.size == 0 {
		return list.InsertLast(newVaule)
	}

	newNode := &SingleNode[T]{value: newVaule}
	newNode.next = list.head
	list.head = newNode
	list.size++

	return newVaule
}

func (list *SingleLinkedList[T]) String() string {

	if list.head == nil {
		return "nil"
	}

	current := list.head
	listAsString := ""

	for current != nil {
		listAsString += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			listAsString += " -> "
		}
		current = current.next
	}

	listAsString += " -> nil"
	return listAsString

}

func (list *SingleLinkedList[T]) RemoveLast() (T, error) {
	var errorItem T

	if list.head == nil {
		return errorItem, errors.New("list is empty")
	}

	if list.head.next == nil {
		value := list.head.value
		list.head = nil
		list.size--
		return value, nil
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	value := current.next.value
	current.next = nil
	list.size--

	return value, nil
}

func (list *SingleLinkedList[T]) Get(index int) (T, error) {

	var errorItem T

	if index > list.size-1 {
		return errorItem, errors.New("index out of bounds")
	}

	if list.head == nil {
		return errorItem, errors.New("list is empty")
	}

	counter := 0
	current := list.head

	for counter < index {
		current = current.next
		counter++
	}

	return current.value, nil

}

func (list *SingleLinkedList[T]) RemoveAt(index int) (T, error) {

	var errorItem T

	if index < 0 || index >= list.size {
		return errorItem, errors.New("index out of bounds")
	}

	if index == 0 {
		removedElement := list.head.value
		list.head = list.head.next
		list.size--
		return removedElement, nil
	}

	current := list.head
	for range index - 1 {
		current = current.next
	}

	removedElement := current.next.value
	current.next = current.next.next
	list.size--

	return removedElement, nil

}

func (list *SingleLinkedList[T]) InsertAt(index int, newValue T) (T, error) {
	var errorItem T

	if index < 0 || index > list.size {
		return errorItem, errors.New("index out of bounds")
	}

	if index == 0 {
		return list.InsertFirst(newValue), nil
	}

	if index == list.size {
		return list.InsertLast(newValue), nil
	}

	newNode := &SingleNode[T]{value: newValue}
	current := list.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
	list.size++

	return newValue, nil
}
