package int32iterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value int32
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
