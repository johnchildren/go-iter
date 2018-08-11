package iiterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	T int
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
