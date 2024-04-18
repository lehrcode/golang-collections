package linkedlist

import (
	"fmt"
	"strings"
)

type node[E any] struct {
	data E
	next *node[E]
}
type LinkedList[E any] struct {
	first *node[E]
}

func (l *LinkedList[E]) Push(value E) {
	var newNode = node[E]{data: value}
	if l.first == nil {
		l.first = &newNode
	} else {
		var lastNode = l.first
		for lastNode.next != nil {
			lastNode = lastNode.next
		}
		lastNode.next = &newNode
	}
}

func (l *LinkedList[E]) Pop() E {
	var value E
	if l.first != nil {
		if l.first.next == nil {
			value = l.first.data
			l.first = nil
		} else {
			var secondLastNode = l.first
			var lastNode = l.first.next
			for lastNode.next != nil {
				secondLastNode = lastNode
				lastNode = lastNode.next
			}
			value = lastNode.data
			secondLastNode.next = nil
		}
	}
	return value
}

func (l LinkedList[E]) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for e := l.first; e != nil; e = e.next {
		fmt.Fprint(&sb, e.data)
		if e.next != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

func (l *LinkedList[E]) Insert(index int, value E) {
	//TODO: implement Insert method
}

func (l *LinkedList[E]) Get(index int) E {
	//TODO: implement Get method
	var value E
	return value
}

func (l *LinkedList[E]) Delete(index int) {
	//TODO: implement Delete method
}

func (l *LinkedList[E]) Len() int {
	//TODO: implement Len method
	return 0
}

func (l *LinkedList[E]) Enqueue(value E) {
	//TODO: implement Enqueue method
}

func (l *LinkedList[E]) Deque() E {
	//TODO: implement Deque method
	var value E
	return value
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}
