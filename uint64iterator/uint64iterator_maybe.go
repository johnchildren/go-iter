package uint64iterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value uint64
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
