package util

import (
	"hash"
	"hash/fnv"
)

type BloomFilter struct {
	bitmap       []bool
	hashfn       hash.Hash64
	numHashFuncs int
	n            int
	size         int
}

// NewBloomFilter creates BloomFilter instance
func NewBloomFilter(numHashFuncs, size int) *BloomFilter {
	return &BloomFilter{
		bitmap:       make([]bool, size),
		hashfn:       fnv.New64(),
		numHashFuncs: numHashFuncs,
		n:            0,
		size:         size,
	}
}

func (bf *BloomFilter) getHash(b []byte) (uint32, uint32) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	h1 := uint32(hash64 & ((1 << 32) - 1))
	h2 := uint32(hash64 >> 32)
	return h1, h2
}

// Add adds element to BloomFilter
func (bf *BloomFilter) Add(b []byte) {
	h1, h2 := bf.getHash(b)
	for i := 0; i < bf.numHashFuncs; i++ {
		index := (h1 + uint32(i)*h2) % uint32(bf.size)
		bf.bitmap[index] = true
	}
	bf.n++
}

// Contains return true if element is in BloomFilter
func (bf *BloomFilter) Contains(b []byte) bool {
	h1, h2 := bf.getHash(b)
	result := true
	for i := 0; i < bf.numHashFuncs && result; i++ {
		index := (h1 + uint32(i)*h2) % uint32(bf.size)
		result = result && bf.bitmap[index]
	}
	return result
}

// GetEleNum returns the number of elements in BloomFilter
func (bf *BloomFilter) GetEleNum() int {
	return bf.n
}