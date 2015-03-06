# mhash

Use a mhash when you have a need for a large number of of locks, but little lock contention.

## Usage

```go
m := mhash.New(10)
y := &bytes.Buffer{}
mutex := m.Get(unsafe.Pointer(y))
mutex.Lock()
mutex.Unlock()
```
