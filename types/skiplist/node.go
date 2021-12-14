package skiplist

// Node is a K-V node, saving elementNode as well
type Node struct {
	next []*Node
	// key is a float64 type for comparing
	key   float64
	value string
}

// Key allows retrieval of the key for a given Element
func (e *Node) Key() float64 {
	return e.key
}

// Value allows retrieval of the value for a given Element
func (e *Node) Value() string {
	return e.value
}

// Next returns the following Element or nil if we're at the end of the list.
// Only operates on the bottom level of the skip list (a fully linked list).
func (element *Node) Next() *Node {
	return element.next[0]
}
