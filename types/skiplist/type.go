package skiplist

import (
	"math/rand"
)

// elementNode is a private struct to be extended by other public struct
type elementNode struct {
	// next nodes of a node
	next []*Element
}

// Element is a K-V node, saving elementNode as well
type Element struct {
	elementNode
	// key is a float64 type for comparing
	key   float64
	value interface{}
}

// Key allows retrieval of the key for a given Element
func (e *Element) Key() float64 {
	return e.key
}

// Value allows retrieval of the value for a given Element
func (e *Element) Value() interface{} {
	return e.value
}

func (e *Element) NextNodesNum() int {
	return len(e.next)
}

// Next returns the following Element or nil if we're at the end of the list.
// Only operates on the bottom level of the skip list (a fully linked list).
func (element *Element) Next() *Element {
	return element.next[0]
}

type SkipList struct {
	elementNode
	maxLevel       int
	Length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	prevNodesCache []*elementNode
}
