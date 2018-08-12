package uiiterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value uint
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
