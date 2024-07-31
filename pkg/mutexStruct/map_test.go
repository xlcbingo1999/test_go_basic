package mutexstruct

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type mapInterface interface { // 构建一个接口, 只要实现了就可以了, 很方便
	Load(k interface{}) (v interface{}, ok bool)
	Store(k, v interface{})
}

func BenchmarkMutexMap(b *testing.B) {
	ms := [...]mapInterface{
		&MutexMap{data: map[interface{}]interface{}{}},
		&RWMutexMap{data: map[interface{}]interface{}{}},
		&sync.Map{}, // 发生足够多的读时，就将 dirty map 复制一份到 read map 上。 从而实现在 read map 上的读操作不再需要昂贵的 Mutex 操作。
	}

	// 测试对于同一个 key 的 n-1 并发读和 1 并发写的性能
	for _, m := range ms {
		b.Run(fmt.Sprintf("%T", m), func(b *testing.B) {
			var i int64
			b.RunParallel(func(pb *testing.PB) {
				// 记录并发执行的 goroutine id
				gid := int(atomic.AddInt64(&i, 1) - 1)

				if gid == 0 {
					// gid 为 0 的 goroutine 负责并发写
					for i := 0; pb.Next(); i++ {
						m.Store(0, i)
					}
				} else {
					// gid 不为 0 的 goroutine 负责并发读
					for pb.Next() {
						m.Load(0)
					}
				}
			})
		})
	}
}
