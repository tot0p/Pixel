package core

import "fmt"

type Stack struct {
	i    uint8
	next *Stack
}

func (s Stack) String() string {
	if s.next == nil {
		return fmt.Sprintf("value %d", s.i)
	}
	return fmt.Sprintf("value %d", s.i) + fmt.Sprintf(" %s", *s.next)
}

func (s *Stack) Push(i uint8) {
	if s.next == nil {
		s.next = &Stack{i, nil}
		return
	}
	s.next.Push(i)
}

func (s *Stack) Pop() uint8 {
	if s.next == nil {
		return s.i
	}
	return s.next.Pop()
}
