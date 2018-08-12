package uint32iterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value uint32
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
