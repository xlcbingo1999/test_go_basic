package contextX

import (
	"fmt"
	"time"
)

// 演示一个不用上下文的场景下,  关闭context的方法

func RunNormalClose() {
	cancel := make(chan struct{})

	go func() { // 第一个协程, 需要在里面实现done和cancel的逻辑
		done := make(chan struct{})

		go func() { // 第二个协程, 模拟正常关闭的流程
			defer func() {
				done <- struct{}{}
			}()

			fmt.Println("do logic....")
			time.Sleep(2 * time.Second)
		}()

		select {
		case <-cancel:
			fmt.Println("cancel")
			<-done // 接收操作, 等别的线程来
		case <-done:
			fmt.Println("done")
			close(done) // 结束执行
		}
	}()

	time.Sleep(1 * time.Second)
	cancel <- struct{}{} // 主协程写入cancel的值, 用于关闭
}
