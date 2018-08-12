package runeiterator

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
	for i := 0; i <= n; i++ {
		nth = it.Next()
	}
	return nth
}

func (it *Iterator) ForEach(f func(x rune)) {
	for {
		switch mt := it.Next().(type) {
		case Just:
			f(mt.Value)
		case None:
			return
		}
	}
}

func (it *Iterator) Chain(other *Iterator) *Iterator {
	return &Iterator{&chainIterator{front: it, back: other, state: chainBoth}}
}

type stepByIterator struct {
	internal *Iterator
	step     int
}

func (s *stepByIterator) Next() Maybe {
	next := s.internal.Next()
	for i := 0; i <= s.step; i++ {
		next = s.internal.Next()
	}
	return next
}

type chainIterator struct {
	front *Iterator
	back  *Iterator
	state chainState
}

type chainState int

const (
	chainBoth chainState = iota
	chainFront
	chainBack
)

func (c *chainIterator) Next() Maybe {
	switch c.state {
	case chainBoth:
		switch mt := c.Next().(type) {
		case Just:
			return mt
		case None:
			c.state = chainBack
			return c.back.Next()
		default:
			panic("invalid maybe type")
		}
	case chainFront:
		return c.front.Next()
	case chainBack:
		return c.back.Next()
	default:
		panic("invalid chain state")
	}
}

type sliceIterator struct {
	idx   int
	len   int
	slice []rune
}

func (s *sliceIterator) Next() Maybe {
	if s.idx >= s.len {
		return None{}
	}
	next := s.slice[s.idx]
	s.idx++
	return Just{next}
}

func FromSlice(s []rune) *Iterator {
	return &Iterator{
		minimalIterator: &sliceIterator{
			idx:   0,
			len:   len(s),
			slice: s,
		},
	}
}
