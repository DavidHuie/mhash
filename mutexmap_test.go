package mutexmap

import (
	"bytes"
	"testing"
	"unsafe"
)

func TestGet(t *testing.T) {
	m := New(2)
	y := &bytes.Buffer{}
	mutex := m.Get(unsafe.Pointer(y))
	mutex.Lock()
	mutex.Unlock()
}
