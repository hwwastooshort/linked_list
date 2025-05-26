package linkedlist

import (
	"testing"
)

func TestDoubleLinkedListSizeWithEmptyList(t *testing.T) {

	list := NewDoubleLinkedList[int]()

	expectedSize := 0
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TestDoubleLinkedListInsertFirstWithOneElement(t *testing.T) {

	list := NewDoubleLinkedList[int]()
	list.InsertFirst(100)

	expectedSize := 1
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TestDoubleLinkedListInsertFirstWithMultipleElements(t *testing.T) {

	list := NewDoubleLinkedList[int]()
	expectedSize := 31

	for i := range expectedSize {
		list.InsertFirst(i)
	}

	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TestDoubleLinkedListInsertLastWithEmptyList(t *testing.T) {

	list := NewDoubleLinkedList[int]()
	list.InsertLast(100)

	expectedSize := 1
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TestDoubleLinkedListInsertLastWithMultipleElements(t *testing.T) {

	list := NewDoubleLinkedList[int]()

	expectedSize := 31

	for i := range expectedSize {
		list.InsertLast(i)
	}

	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TestDoubleLinkedListRemoveFirstWithEmptyList(t *testing.T) {

	list := NewDoubleLinkedList[int]()

	_, err := list.removeFirst()

	if err == nil {
		t.Errorf("it should not be possible to remove the first item of an empty list")
	}

}

func TestDoubleLinkedListRemoveFirstWithOneElement(t *testing.T) {

	list := NewDoubleLinkedList[int]()
	list.InsertFirst(100)
	list.removeFirst()

	expectedSize := 0
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TetDoubleLinkedListRemoveFirstWithMultipleElements(t *testing.T) {

	list := NewDoubleLinkedList[int]()

	for i := range 100 {
		list.InsertFirst(i)
	}

	if list.Size() != 100 {
		t.Errorf("error while filling the array")
	}

	for range 100 {
		list.removeFirst()
	}

}
