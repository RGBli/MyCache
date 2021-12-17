package types

type String struct {
	s string
}

func NewString(s string) *String {
	return &String{s: s}
}

func (s *String) Size() uint64 {
	return uint64(len(s.s))
}

func (s *String) ToString() string {
	return s.s
}

func (s *String) Len() int {
	return len(s.s)
}
