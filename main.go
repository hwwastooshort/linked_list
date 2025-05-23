package main

import (
	"fmt"

	"github.com/hwwastooshort/linked_list/linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList[int]()

	for i := range 10 {
		list.InsertLast(i)
	}

	fmt.Println(list)

}
