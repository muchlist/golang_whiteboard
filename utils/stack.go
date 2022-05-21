package utils

import "errors"

type stackString struct {
	data []string
}

func NewStackString() *stackString {
	return &stackString{data: []string{}}
}

func (s *stackString) Push(text string) {
	s.data = append(s.data, text)
}

func (s *stackString) Top() (string, error) {
	if s.Empty() {
		return "", errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *stackString) Pop() (string, error) {
	if s.Empty() {
		return "", errors.New("stack is empty")
	}
	popReturn := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popReturn, nil
}

func (s *stackString) Empty() bool {
	return len(s.data) == 0
}

func (s *stackString) Size() int {
	return len(s.data)
}
