package semaphoreX

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// 多值信号量, 可以用于让一组协程同时处理某个临界区域的内容

func doSomething(u string) {
	log.Println(u)
	time.Sleep(2 * time.Second)
}

const (
	Limit  = 2 // 同时进入临界区的goroutine上限
	Weight = 2 // 每个goroutine获取信号量资源的权重
)

func RunSemaphore() {
	urls := []string{
		"ex",
		"ex1",
		"ex2",
		"ex3",
		"ex4",
	}

	s := semaphore.NewWeighted(Limit)
	var w sync.WaitGroup

	for _, u := range urls {
		w.Add(1) // 阻塞一个
		go func(u string) {
			s.Acquire(context.Background(), Weight) // 获取Weight个信号量
			doSomething(u)
			s.Release(Weight) // 释放Weight个信号量
			w.Done()          // 结束一个
		}(u)
	}

	w.Wait()
	log.Println("All done")
}
