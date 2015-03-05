package mutexmap

import (
	"encoding/binary"
	"hash/fnv"
	"sync"
	"unsafe"
)

type Mutexmap struct {
	size    int64
	mutexes []*sync.Mutex
}

func New(size int64) *Mutexmap {
	mutexes := make([]*sync.Mutex, size)
	for i := int64(0); i < size; i++ {
		mutexes[i] = &sync.Mutex{}
	}
	return &Mutexmap{size, mutexes}
}

func (m *Mutexmap) Get(ptr unsafe.Pointer) *sync.Mutex {
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
