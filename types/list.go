package types

import "errors"

type List []string

func NewEmptyList() List {
	return List(make([]string, 0))
}

func NewList(strs []string) List {
	return List(strs)
}

func (l List) Size() uint64 {
	var size uint64 = 0
	for _, v := range l {
		size += uint64(len(v))
	}
	return size
}

func (l List) Len() int {
	return len(l)
}

func (l List) Get(i int) (string, error) {
	if i >= len(l) {
		err := errors.New("index out of bounds")
		return "", err
	}
	return l[i], nil
}

func (l List) GetAll() []string {
	return []string(l)
}

func (l List) Set(i int, s string) error {
	if i >= len(l) {
		err := errors.New("index out of bounds")
		return err
	}
	l[i] = s
	return nil
}

func (l List) Add(s string) {
	l = append([]string(l), s)
}

func (l List) Remove(i int) error {
	if i >= len(l) {
		err := errors.New("index out of bounds")
		return err
	}
	l = append([]string(l)[:i], l[i+1:]...)
	return nil
}
