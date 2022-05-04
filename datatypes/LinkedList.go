package datatypes

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	Head *listNode[T]
	Tail *listNode[T]
	Len  int
}

type listNode[T any] struct {
	Val    T
	Next   *listNode[T]
	Before *listNode[T]
}

func CreateList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Len: 0,
	}
}

func CreateListFromArr[T any](arr *[]T) *LinkedList[T] {
	list := &LinkedList[T]{}
	var currentNode *listNode[T]

	if len(*arr) == 0 {
		return CreateList[T]()
	}
	currentNode = &listNode[T]{
		Val: (*arr)[0],
	}
	list.Head = currentNode

	for _, ele := range (*arr)[1:] {
		tempNode := &listNode[T]{
			Val:    ele,
			Before: currentNode,
		}
		currentNode.Next = tempNode
		currentNode = tempNode
		list.Tail = currentNode

	}
	return list
}

func (list *LinkedList[T]) String() string {
	var b strings.Builder
	current := list.Head
	b.WriteString("{")
	for current != nil {
		b.WriteString(current.String() + ", ")
		current = current.Next
	}
	return b.String()[:b.Len()-2] + "}"
}

func (node *listNode[T]) String() string {
	return fmt.Sprint(node.Val)
}
