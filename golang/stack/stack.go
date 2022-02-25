package stack

import "errors"

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("empty stack")
	}
	element := s.Top()
	index := s.Size() - 1
	*s = (*s)[:index]
	return element, nil
}

func (s Stack) Empty() bool {
	if s == nil {
		return true
	}
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Top() interface{} {
	size := s.Size()
	if size == 0 {
		return nil
	}
	return s[size-1]
}
