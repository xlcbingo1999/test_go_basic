package barrierX

import (
	"log"
	"sync"
	"time"
)

// barrier的作用: 等待一组goroutine在同一时刻完成某项操作, 然后再进行下一项操作, 是常见的同步机制。
// 注意: 一组线程在达到屏障前阻塞等待，直到所有线程都到达屏障时，屏障才会打开，所有线程才会继续执行。
// 注意: 在所有线程继续执行之前，可以通过屏障的构造函数指定一个回调函数，在所有线程到达屏障之后执行回调函数。
// barrier的竞品: WaitGroup类型可以利用Add和Done进行一些同步操作, Wait()会阻塞相关的主routine。但此时

type CyclicBarrier struct {
	n          int
	count      int
	cond       *sync.Cond // 条件变量
	beforeFunc func()
	afterFunc  func()
}

func NewCyclicBarrier(n int, beforeFunc func(), afterFunc func()) *CyclicBarrier {
	c := sync.NewCond(&sync.Mutex{})

	return &CyclicBarrier{
		n:          n,
		count:      0,
		cond:       c,
		beforeFunc: beforeFunc,
		afterFunc:  afterFunc,
	}
}

func (cb *CyclicBarrier) await() {
	cb.cond.L.Lock() // 利用管道里面的锁进行加锁, 达成一个临界区域, 只有一个Goroutine可以进入这里
	defer cb.cond.L.Unlock()

	cb.beforeFunc()

	cb.count += 1

	if cb.count == cb.n { // 当执行了n次锁之后, 就可以离开了
		cb.count = 0
		cb.cond.Broadcast() // 广播, 非常粗暴地通知(Signal)所有调用了cb.cond.Wait()的Goroutine
		cb.afterFunc()      // 特殊的Goroutine需要在这里先进行处理
		return
	}

	cb.cond.Wait() // 阻塞等, 条件变量的Wait()必须等到来自条件变量的通知
	cb.afterFunc()
}

func RunBarrier() {
	count := 5
	funcCount := 10
	barrier := NewCyclicBarrier(count, func() {
		log.Println("before Func")
	}, func() {
		log.Println("after Func")
	})

	for i := 0; i < funcCount; i++ {
		// 开启count个Goroutine, 每个函数都会调用await()用于等待一个批次的Goroutine到达屏障, 当足够的Goroutine到达后, 会一次性执行这些Goroutine的after
		go func(i int) {
			log.Println("Goroutine ", i, " before await")
			barrier.await()
			log.Println("Goroutine ", i, " after await")
		}(i)
		time.Sleep(1 * time.Second)
	}

}
