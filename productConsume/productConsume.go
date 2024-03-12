// refer: https://golang.design/under-the-hood/zh-cn/part1basic/ch05sync/cond/

package productConsume

import (
	"fmt"
	"sync"
)

func RunProductConsume() {
	cond := sync.NewCond(&sync.Mutex{})
	goods := 0

	// consume
	go func() {
		for {
			// 只有一个消费者可以进来
			cond.L.Lock()

			// 当没有东西的时候, 只能阻塞等待
			for goods == 0 {
				cond.Wait()
			}

			// 真正进行消费
			goods -= 1
			fmt.Println("consume: ", goods)

			// 需要生产者进行生产, 这里只是唤醒一个Goroutine, 但是不一定是生产者吧? 如果永远唤醒的都是消费者怎么办?
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// product
	for {
		// 生产者开始生产
		cond.L.Lock()

		// 如果生产的内容太多了, 就慢一些
		for goods == 10 {
			cond.Wait()
		}

		goods += 1
		fmt.Println("product: ", goods)

		cond.Signal()
		cond.L.Unlock()
	}
}
