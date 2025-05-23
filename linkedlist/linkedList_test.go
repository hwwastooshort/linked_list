package linkedlist

import (
	"testing"
)

func TestLinkedListWithNoEntries(t *testing.T) {
	list := NewLinkedList[int]()

	expectedSize := 0
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size: %d after initialization, but hast size: %d", expectedSize, actualSize)
	}

}

func TestLinkedListWithOneEntry(t *testing.T) {
	list := NewLinkedList[int]()

	list.InsertLast(10)
	expectedSize := 1
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size: %d but has size: %d", expectedSize, actualSize)
	}

}

func TestListWithMultipleEntries(t *testing.T) {
	list := NewLinkedList[int]()

	for i := range 10 {
		list.InsertLast(i)
	}

	expectedSize := 10
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size: %d but has size: %d", expectedSize, actualSize)
	}

}

func TestLinkedListRemoveLastWhenListIsEmpty(t *testing.T) {

	list := NewLinkedList[int]()

	list.RemoveLast()

	expectedSize := 0
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size: %d but has size: %d", expectedSize, actualSize)
	}
}

func TestLinkedListRemoveLastWithOneElement(t *testing.T) {

	list := NewLinkedList[int]()
	list.InsertLast(10)

	expectedSize := 1
	actualSize := list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size: %d but has size: %d", expectedSize, actualSize)
	}

	list.RemoveLast()

	expectedSize = 0
	actualSize = list.Size()

	if actualSize != expectedSize {
		t.Errorf("list should have size: %d but has size: %d", expectedSize, actualSize)
	}

}

func TestLinkedListGetWithNoElements(t *testing.T) {

	list := NewLinkedList[string]()

	_, err := list.Get(0)
	if err == nil {
		t.Errorf("it shouldn't be possible to access the 0th element of an empty list")
	}
}

func TestLinkedListGetWithElements(t *testing.T) {

	list := NewLinkedList[int]()
	list.InsertLast(10)
	list.InsertLast(20)
	list.InsertLast(30)

	tests := []struct {
		index    int
		expected int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
	}

	for _, tt := range tests {
		value, err := list.Get(tt.index)
		if err != nil {
			t.Errorf("unexpected error for index %d: %v", tt.index, err)
		}
		if value != tt.expected {
			t.Errorf("expected value %d at index %d, got %d", tt.expected, tt.index, value)

		}

	}

}

func TestLinkedListRemoveAtWithEmptyList(t *testing.T) {

	list := NewLinkedList[int]()
	_, err := list.RemoveAt(0)

	if err == nil {
		t.Errorf("it should not be possible to remove the 0th element of an empty list")
	}

}

func TestLinkedListRemoveAtWithOneElement(t *testing.T) {

	list := NewLinkedList[int]()
	insertedValue := 10
	list.InsertLast(insertedValue)

	if list.Size() != 1 {
		t.Errorf("the list should have 1 element but has %d", list.Size())
	}

	removedValue, err := list.RemoveAt(0)

	if removedValue != insertedValue {
		t.Errorf("the removed value was expected to be %d but is %d", insertedValue, removedValue)
	}

	if err != nil {
		t.Error("removing element from list went wrong")
	}

}

func TestLinkedListRemoveAtWithListWithMultilpeValues(t *testing.T) {

	list := NewLinkedList[int]()

	for i := range 1000 {
		list.InsertLast(i)
	}

	for i := range 1000 {

		value, err := list.Get(i)
		if err != nil {
			t.Errorf("something went wrong while trying to access the elements of the list")
		}

		if value != i {
			t.Errorf("values don't match in loop")
		}
	}

	for range 200 {

		_, err := list.RemoveAt(0)
		if err != nil {
			t.Errorf("Something went wrong while trying to delete elements of the list")
		}

	}

	expecteCurrentSize := 1000 - 200
	actualCurrentSize := list.Size()

	if actualCurrentSize != expecteCurrentSize {
		t.Errorf("list doesnt have correct length")
	}

}

func TestLinkedListInsertFirst(t *testing.T) {
	list := NewLinkedList[int]()

	list.InsertFirst(100)

	if list.Size() != 1 {
		t.Errorf("list should have size 1 but has size %d", list.Size())
	}

	numbersToInsert := []int{1, 2, 4}

	for _, num := range numbersToInsert {
		list.InsertFirst(num)
	}

	//list should look like this: 4 -> 2 -> 1 -> 100 -> nil
	expected := []int{4, 2, 1, 100}

	if list.Size() != len(expected) {
		t.Errorf("expected size %d but got %d", len(expected), list.Size())
	}

	for i, exp := range expected {
		value, err := list.Get(i)
		if err != nil {
			t.Fatalf("unexpected error at index %d: %v", i, err)
		}
		if value != exp {
			t.Errorf("expected %d at index %d but got %d", exp, i, value)
		}
	}
}

func TestLinkedListInsertAtAsLastElement(t *testing.T) {

	list := NewLinkedList[int]()
	list.InsertAt(0, 10)

	if list.Size() != 1 {
		t.Errorf("list should have size %d but hast size %d", 1, list.Size())
	}

	value, err := list.Get(0)
	if err != nil {
		t.Error("Something went wrong: ", err)
	}

	if value != 10 {
		t.Errorf("%d was expected but is actually %d", 10, value)
	}
}

func TestLinkedListInsertAt(t *testing.T) {
	list := NewLinkedList[int]()

	_, err := list.InsertAt(0, 10)
	if err != nil {
		t.Errorf("InsertAt(0, 10) failed on empty list: %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("expected size 1 but got %d", list.Size())
	}
	value, err := list.Get(0)
	if err != nil || value != 10 {
		t.Errorf("expected value 10 at index 0, got %d, err: %v", value, err)
	}

	_, err = list.InsertAt(0, 5)
	if err != nil {
		t.Errorf("InsertAt(0, 5) failed: %v", err)
	}

	if list.Size() != 2 {
		t.Errorf("expected size 2 but got %d", list.Size())
	}

	v0, _ := list.Get(0)
	if v0 != 5 {
		t.Errorf("expected 5 at index 0, got %d", v0)
	}

	v1, _ := list.Get(1)
	if v1 != 10 {
		t.Errorf("expected 10 at index 1, got %d", v1)
	}

	_, err = list.InsertAt(1, 7)
	if err != nil {
		t.Errorf("InsertAt(1, 7) failed: %v", err)
	}

	if list.Size() != 3 {
		t.Errorf("expected size 3 but got %d", list.Size())
	}

	expected := []int{5, 7, 10}
	for i, exp := range expected {
		val, err := list.Get(i)
		if err != nil || val != exp {
			t.Errorf("expected %d at index %d, got %d, err: %v", exp, i, val, err)
		}
	}

	_, err = list.InsertAt(list.Size(), 20)
	if err != nil {
		t.Errorf("InsertAt at end failed: %v", err)
	}
	if list.Size() != 4 {
		t.Errorf("expected size 4 but got %d", list.Size())
	}
	vEnd, _ := list.Get(3)
	if vEnd != 20 {
		t.Errorf("expected 20 at last index, got %d", vEnd)
	}

	_, err = list.InsertAt(-1, 99)
	if err == nil {
		t.Error("expected error for negative index, got nil")
	}

	_, err = list.InsertAt(list.Size()+1, 99)
	if err == nil {
		t.Error("expected error for out-of-bounds index, got nil")
	}
}
