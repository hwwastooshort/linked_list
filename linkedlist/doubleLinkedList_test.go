package linkedlist

import "testing"

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
