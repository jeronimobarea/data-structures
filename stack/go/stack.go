package stack

import "errors"

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Size() == 0 {
		return nil, errors.New("empty stack")
	}
	index := s.Size() - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, nil
}

func (s Stack) Empty() bool {
	if s == nil {
		return true
	}
	return len(s) == 0
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Top() interface{} {
	if s.Size() == 0 {
		return nil
	}
	return s[s.Size()-1]
}
