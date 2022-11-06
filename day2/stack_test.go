package day2_test

import "testing"

type Stack[T any] struct {
	list []T
}

func (s *Stack[T]) Add(e T) {
	s.list = append(s.list, e)
}

func (s *Stack[T]) Pop() (T, bool) {
	size := len(s.list)
	if size == 0 {
		var zero T
		return zero, false
	} else if size == 1 {
		val := s.list[0]
		s.list = nil
		return val, true
	} else {
		val := s.list[size-1]
		s.list = s.list[:size-1]
		return val, true
	}
}

func TestStack(t *testing.T) {
	var s Stack[int]
	s.Add(1)
	s.Add(2)
	s.Add(3)

	{
		val, _ := s.Pop()
		t.Logf("%v", val)
	}

	{
		val, _ := s.Pop()
		t.Logf("%v", val)
	}
}
