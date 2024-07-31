package atomicX

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var mu sync.Mutex
var wg sync.WaitGroup

func normalAdd() {
	x++

	wg.Done()
}

func mutexAdd() {
	mu.Lock()
	defer mu.Unlock()

	x++
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func RunAtomic() {
	begin := time.Now()
	for i := 0; i < 50000; i++ {
		wg.Add(1)

		go atomicAdd()
		// go mutexAdd()
		// go normalAdd()
	}
	wg.Wait()
	end := time.Now()

	fmt.Println(x)
	fmt.Println(end.Sub(begin))
}
