# mutexmap

Use a mutexmap when you have a need for a large number of of locks, but little lock contention.

## Usage

```go
mmap := mutexmap.New(10)
y := &bytes.Buffer{}
mutex := mmap.Get(unsafe.Pointer(y))
mutex.Lock()
mutex.Unlock()
```
