package iterator

//go:generate ./gen_iterator.sh int iiterator
//go:generate ./gen_iterator.sh int32 int32iterator
//go:generate ./gen_iterator.sh int64 int64iterator
//go:generate ./gen_iterator.sh uint uiiterator
//go:generate ./gen_iterator.sh uint32 uint32iterator
//go:generate ./gen_iterator.sh uint64 uint64iterator

//go:generate ./gen_iterator.sh string stringiterator
//go:generate ./gen_iterator.sh byte byteiterator
//go:generate ./gen_iterator.sh rune runeiterator
//go:generate ./gen_iterator.sh []byte bytesliceiterator
