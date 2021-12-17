package types

import "errors"

type List struct {
	slice []string
}

func NewEmptyList() *List {
	return &List{slice: make([]string, 0)}
}

func NewList(strs []string) *List {
	list := &List{slice: make([]string, len(strs))}
	for i := 0; i < len(strs); i++ {
		list.slice[i] = strs[i]
	}
	return list
}

func (l *List) Size() uint64 {
	var size uint64 = 0
	for _, v := range l.slice {
		size += uint64(len(v))
	}
	return size
}

func (l *List) Len() int {
	return len(l.slice)
}

func (l *List) Get(i int) (string, error) {
	if i >= len(l.slice) {
		err := errors.New("index out of bounds")
		return "", err
	}
	return l.slice[i], nil
}

func (l *List) GetAll() []string {
	return l.slice
}

func (l *List) Set(i int, s string) error {
	if i >= len(l.slice) {
		err := errors.New("index out of bounds")
		return err
	}
	l.slice[i] = s
	return nil
}

func (l *List) Add(s string) {
	l.slice = append(l.slice, s)
}

func (l *List) Remove(i int) error {
	if i >= len(l.slice) {
		err := errors.New("index out of bounds")
		return err
	}
	l.slice = append(l.slice[:i], l.slice[i+1:]...)
	return nil
}
