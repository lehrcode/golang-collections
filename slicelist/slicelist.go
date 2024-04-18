package slicelist

import (
	"fmt"
	"strings"
)

type SliceList[E any] struct {
	slice []E
}

func (s *SliceList[E]) Push(value E) {
	s.slice = append(s.slice, value)
}

func (s *SliceList[E]) Pop() E {
	var value E
	var lastIndex = len(s.slice) - 1
	if lastIndex >= 0 {
		value = s.slice[lastIndex]
		s.slice = s.slice[:lastIndex]
	}
	return value
}

func (s SliceList[E]) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for index, element := range s.slice {
		if index > 0 {
			sb.WriteString(", ")
		}
		fmt.Fprint(&sb, element)
	}
	sb.WriteRune(']')
	return sb.String()
}

func New[T any]() *SliceList[T] {
	return &SliceList[T]{}
}
