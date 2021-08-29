package utils

type setMapBased struct {
	setMap map[string]struct{}
}

func NewSet() *setMapBased {
	return &setMapBased{
		setMap: make(map[string]struct{}),
	}
}

func (s *setMapBased) Add(text string) (success bool) {
	if _, exist := s.setMap[text]; exist {
		success = false
		return
	}
	s.setMap[text] = struct{}{}
	success = true
	return
}

func (s *setMapBased) Remove(text string) {
	delete(s.setMap, text)
}

func (s *setMapBased) Reveal() []string {
	strSlice := make([]string, 0)
	for key := range s.setMap {
		strSlice = append(strSlice, key)
	}
	return strSlice
}

func (s *setMapBased) Len() int {
	return len(s.setMap)
}
