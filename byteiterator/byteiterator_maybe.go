package byteiterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value byte
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
