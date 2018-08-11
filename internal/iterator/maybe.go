package iterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	T T
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
