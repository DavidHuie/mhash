package mhash

import (
	"encoding/binary"
	"hash/fnv"
	"sync"
	"unsafe"
)

type Mhash struct {
	size    uint64
	mutexes []*sync.Mutex
}

func New(size uint64) *Mhash {
	mutexes := make([]*sync.Mutex, size)
	for i := uint64(0); i < size; i++ {
		mutexes[i] = &sync.Mutex{}
	}
	return &Mhash{size, mutexes}
}

func (m *Mhash) Get(ptr unsafe.Pointer) *sync.Mutex {
	// Convert the pointer to a byte array
	intPtr := uint64(uintptr(ptr))
	bytes := make([]byte, 8)
	binary.PutUvarint(bytes, intPtr)

	// Hash the byte array
	hash := fnv.New64()
	hash.Write(bytes)
	key := int64(hash.Sum64() % uint64(m.size))
	return m.mutexes[key]
}
