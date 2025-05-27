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

	_, err := list.RemoveFirst()

	if err == nil {
		t.Errorf("it should not be possible to remove the first item of an empty list")
	}

}

func TestDoubleLinkedListRemoveFirstWithOneElement(t *testing.T) {

	list := NewDoubleLinkedList[int]()
	list.InsertFirst(100)
	list.RemoveFirst()

	expectedSize := 0
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size %d but has size %d", expectedSize, actualSize)
	}

}

func TestDoubleLinkedListRemoveFirstWithMultipleElements(t *testing.T) {

	list := NewDoubleLinkedList[int]()

	for i := range 100 {
		list.InsertFirst(i)
	}

	if list.Size() != 100 {
		t.Errorf("error while filling the array")
	}

	for range 100 {
		list.RemoveFirst()
	}

	if list.Size() != 0 {
		t.Errorf("array should have size 0 and thus should be empty")
	}

}

func TestDoubleLinkedListRemoveLastItemWithEmptyList(t *testing.T) {

	list := NewDoubleLinkedList[int]()

	_, err := list.RemoveLast()

	if err == nil {
		t.Errorf("it should not be possible to remove the last item of an empty list")
	}

}

func TestDoubleLinkedListRemoveLastItemWithOneElement(t *testing.T) {

	list := NewDoubleLinkedList[int]()
	list.InsertFirst(100)

	value, err := list.RemoveLast()

	if err != nil {
		t.Errorf("error while trying to remove last element")
	}

	if value != 100 {
		t.Errorf("list.RemoveLast returned the wrong value, 100 was expected but %d recieved", value)
	}
}

func TestDoubleLinkedListGet(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	for i := 0; i < 10; i++ {
		list.InsertLast(i * 10)
	}

	value, err := list.Get(0)
	if err != nil || value != 0 {
		t.Errorf("expected 0 at index 0, got %v, err: %v", value, err)
	}

	value, err = list.Get(5)
	if err != nil || value != 50 {
		t.Errorf("expected 50 at index 5, got %v, err: %v", value, err)
	}

	_, err = list.Get(10)
	if err == nil {
		t.Errorf("expected error when accessing out-of-bounds index")
	}
}

func TestDoubleLinkedListInsertAt(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.InsertLast(1)
	list.InsertLast(3)
	list.InsertLast(4)

	_, err := list.InsertAt(1, 2) // list should become: 1 2 3 4
	if err != nil {
		t.Errorf("unexpected error while inserting: %v", err)
	}

	val, _ := list.Get(1)
	if val != 2 {
		t.Errorf("expected value 2 at index 1, got %v", val)
	}

	_, err = list.InsertAt(5, 99) // invalid
	if err == nil {
		t.Errorf("expected error when inserting at invalid index")
	}
}

func TestDoubleLinkedListRemoveAt(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	for i := 1; i <= 5; i++ {
		list.InsertLast(i) // list: 1 2 3 4 5
	}

	val, err := list.RemoveAt(2) // remove "3"
	if err != nil || val != 3 {
		t.Errorf("expected to remove 3 at index 2, got %v, err: %v", val, err)
	}

	val, _ = list.Get(2)
	if val != 4 {
		t.Errorf("expected 4 at index 2 after removal, got %v", val)
	}

	_, err = list.RemoveAt(10)
	if err == nil {
		t.Errorf("expected error when removing at invalid index")
	}

	if list.Size() != 4 {
		t.Errorf("expected size 4 after removal, got %d", list.Size())
	}
}
