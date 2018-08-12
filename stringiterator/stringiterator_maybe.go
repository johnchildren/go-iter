package stringiterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value string
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
