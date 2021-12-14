package skiplist

import (
	"math"
	"math/rand"
	"time"
)

const (
	DefaultMaxLevel    int     = 18
	DefaultProbability float64 = 1 / math.E
)

type SkipList struct {
	head        *Node
	next        []*Node
	maxLevel    int
	length      int
	randSource  rand.Source
	probability float64
	probTable   []float64
}

// NewWithMaxLevel creates a new skip list with MaxLevel set to the provided number.
// maxLevel has to be int(math.Ceil(math.Log(N))) for DefaultProbability (where N is an upper bound on the
// number of elements in a skip list). Returns a pointer to the new list.
func NewWithMaxLevel(maxLevel int) *SkipList {
	if maxLevel < 1 || maxLevel > 64 {
		panic("maxLevel for a SkipList must be a positive integer <= 64")
	}

	return &SkipList{
		head:        &Node{next: make([]*Node, maxLevel)},
		next:        make([]*Node, maxLevel),
		maxLevel:    maxLevel,
		randSource:  rand.New(rand.NewSource(time.Now().UnixNano())),
		probability: DefaultProbability,
		probTable:   probabilityTable(DefaultProbability, maxLevel),
	}
}

// New creates a new skip list with default parameters. Returns a pointer to the new list.
func New() *SkipList {
	return NewWithMaxLevel(DefaultMaxLevel)
}

func (list *SkipList) Len() int { 
	return list.length
}

func (list *SkipList) Size() uint64 {
	var size uint64 = uint64(32 + 8*(len(list.probTable)+len(list.next)))
	node := list.Front()
	for node != nil {
		size += uint64(len(node.Value()) + len(node.next)*8 + 8)
		node = node.Next()
	}
	return size
}

// Front returns the first node of the list
func (list *SkipList) Front() *Node {
	return list.next[0]
}

// Set inserts a value in the list with the specified key, ordered by the key
// If the key exists, it updates the value in the existing node
// Returns a pointer to the new element
func (list *SkipList) Set(key float64, value string) *Node {
	var node *Node
	prevs := list.getPrevElementNodes(key)

	// if key exists, only update the value
	if node = prevs[0].next[0]; node != nil && node.key == key {
		node.value = value
		return node
	}

	// if key doesn't exist, create a new node
	node = &Node{
		next:  make([]*Node, list.randLevel()),
		key:   key,
		value: value,
	}

	for i := range node.next {
		node.next[i] = prevs[i].next[i]
		prevs[i].next[i] = node
	}

	list.length++
	return node
}

// Get finds an element by key. It returns element pointer if found, nil if not found.
func (list *SkipList) Get(key float64) *Node {
	var node *Node = list.head
	var next *Node

	// retrieve from list.maxLevel - 1 to 0, to achieve O(logN) time complexity
	for i := list.maxLevel - 1; i >= 0; i-- {
		next = node.next[i]
		for next != nil && next.key < key {
			node = next
			next = next.next[i]
		}
	}

	if next != nil && next.key == key {
		return next
	}
	return nil
}

// Remove deletes an element with given key from the list.
// Returns removed element pointer if found, nil if not found.
func (list *SkipList) Remove(key float64) *Node {
	prevs := list.getPrevElementNodes(key)

	// found the node and remove it
	if node := prevs[0].next[0]; node != nil && node.key == key {
		for i, v := range node.next {
			prevs[i].next[i] = v
		}
		list.length--
		return node
	}
	return nil
}

// Contains return if a key exists in the list, based on Get method
func (list *SkipList) Contains(key float64) bool {
	return list.Get(key) != nil
}

// getPrevElementNodes is the private search method that other functions use.
// Finds the previous nodes on each level relative to the current Element and caches them in prevNodesCache.
// Note that key doesn't have to exist.
func (list *SkipList) getPrevElementNodes(key float64) ([]*Node) {
	var node *Node = list.head
	var next *Node
	prevNodesCache := make([]*Node, list.maxLevel)

	for i := list.maxLevel - 1; i >= 0; i-- {
		next = node.next[i]
		for next != nil && key > next.key {
			node = next
			next = next.next[i]
		}
		prevNodesCache[i] = node
	}
	return prevNodesCache
}

// SetProbability changes the current P value of the list.
// It doesn't alter any existing data, only changes how future insert heights are calculated.
func (list *SkipList) SetProbability(newProbability float64) {
	list.probability = newProbability
	list.probTable = probabilityTable(list.probability, list.maxLevel)
}

// probabilityTable calculates in advance the probability of a new node having a given level.
// probability is in [0, 1], MaxLevel is (0, 64]
// Returns a table of floating point probabilities that each level should be included during an insert.
func probabilityTable(probability float64, MaxLevel int) (table []float64) {
	for i := 1; i <= MaxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}

// randLevel calculate the random level of a newly inserted node
func (list *SkipList) randLevel() (level int) {
	// Our random number source only has Int63(), so we have to produce a float64 from it
	r := float64(list.randSource.Int63()) / (1 << 63)

	level = 1
	for level < list.maxLevel && r < list.probTable[level] {
		level++
	}
	return
}
