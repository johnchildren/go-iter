package int64iterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value int64
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
