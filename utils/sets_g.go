package utils

type setMap[T comparable] struct {
	setMap map[T]struct{}
}

func NewSetG[T comparable]() *setMap[T] {
	return &setMap[T]{
		setMap: make(map[T]struct{}),
	}
}

func (s *setMap[T]) Add(val T) (success bool) {
	if _, exist := s.setMap[val]; exist {
		success = false
		return
	}
	s.setMap[val] = struct{}{}
	success = true
	return
}

func (s *setMap[T]) Remove(val T) {
	delete(s.setMap, val)
}

func (s *setMap[T]) Reveal() []T {
	strSlice := make([]T, 0)
	for key := range s.setMap {
		strSlice = append(strSlice, key)
	}
	return strSlice
}

func (s *setMap[T]) Len() int {
	return len(s.setMap)
}
