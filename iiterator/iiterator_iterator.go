package iiterator

type Iterator interface {
	Next() Maybe
	Count() int
	Last() int
}

type sliceIterator struct {
	idx   int
	len   int
	slice []int
}

func (s *sliceIterator) Next() Maybe {
	if s.idx >= s.len {
		return None{}
	}
	next := s.slice[s.idx]
	s.idx++
	return Just{next}
}

func (s *sliceIterator) Count() int {
	return s.len
}

func (s *sliceIterator) Last() int {
	return s.slice[s.len-1]
}

func FromSlice(s []int) Iterator {
	return &sliceIterator{
		idx:   0,
		len:   len(s),
		slice: s,
	}
}
