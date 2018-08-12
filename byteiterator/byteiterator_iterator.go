package byteiterator

type minimalIterator interface {
	Next() Maybe
}

type Iterator struct {
	minimalIterator
}

func (it *Iterator) Count() int {
	count := 0
	for {
		switch it.Next().(type) {
		case Just:
			count++
		case None:
			return count
		}
	}
}

func (it *Iterator) Last() Maybe {
	var last Maybe = None{}
	for {
		switch mt := it.Next().(type) {
		case Just:
			last = mt
		case None:
			return last
		}
	}
}

func (it *Iterator) Nth(n int) Maybe {
	var nth Maybe = None{}
	for v := 0; v <= n; v++ {
		nth = it.Next()
	}
	return nth
}

func (it *Iterator) ForEach(f func(x byte)) {
	for {
		switch mt := it.Next().(type) {
		case Just:
			f(mt.Value)
		case None:
			return
		}
	}
}

type sliceIterator struct {
	idx   int
	len   int
	slice []byte
}

func (s *sliceIterator) Next() Maybe {
	if s.idx >= s.len {
		return None{}
	}
	next := s.slice[s.idx]
	s.idx++
	return Just{next}
}

func FromSlice(s []byte) *Iterator {
	return &Iterator{
		minimalIterator: &sliceIterator{
			idx:   0,
			len:   len(s),
			slice: s,
		},
	}
}
