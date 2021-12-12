package types

type String string

func NewString(s string) String {
	return String(s)
}

func (s String) Size() uint64 {
	return uint64(len(s))
}

func (s String) Len() uint64 {
	return uint64(len(s))
}