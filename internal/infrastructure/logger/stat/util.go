package stat

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const mutexLocked = 1 << iota

// The mutex which supports try-locking.
type Mutex struct {
	sync.Mutex
}

// TryLock acquires the lock only if it is free at the time of invocation.
func (tl *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&tl.Mutex)), 0, mutexLocked)
}

// SliceHeader is a safe version of SliceHeader used within this project.
type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// StringHeader is a safe version of StringHeader used within this project.
type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

func CurrentTimeMillis() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}
