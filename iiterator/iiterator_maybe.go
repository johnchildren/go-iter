package iiterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value int
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
