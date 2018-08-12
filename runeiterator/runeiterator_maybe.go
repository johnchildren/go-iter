package runeiterator

type Maybe interface {
	isMaybe()
}

type Just struct {
	Value rune
}

func (Just) isMaybe() {}

type None struct{}

func (None) isMaybe() {}
