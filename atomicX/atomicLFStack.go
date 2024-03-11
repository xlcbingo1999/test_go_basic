package atomicX

import (
	"sync/atomic"
	"time"
	"unsafe"
)

type LFStack struct {
	Next unsafe.Pointer
	Item int
}

var lfhead unsafe.Pointer // 记录栈头信息

func (head *LFStack) Push(i int) *LFStack { // 强制逃逸
	new := &LFStack{Item: i}
	newptr := unsafe.Pointer(new)
	for {
		old := atomic.LoadPointer(&lfhead)
		new.Next = old

		if atomic.CompareAndSwapPointer(&lfhead, old, newptr) {
			break
		}
	}
	return new
}

func (head *LFStack) Pop() int {
	for {
		time.Sleep(time.Nanosecond) // 可以让CPU缓一缓
		old := atomic.LoadPointer(&lfhead)
		if old == nil {
			return 0
		}

		if lfhead == old {
			new := (*LFStack)(old).Next
			if atomic.CompareAndSwapPointer(&lfhead, old, new) {
				return 1
			}
		}
	}
}
