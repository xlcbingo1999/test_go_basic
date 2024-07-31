package contextX

import (
	"context"
	"log"
	"time"
)

// 只读的channel: <-chan int
// 只写的channel: chan<- int
// 可读可写的channel: chan int

func gen(ctx context.Context) <-chan int { // <-chan int 这是一个只读的管道, 只可以读取里面的内容, 而不能往里面写内容
	ch := make(chan int)
	go func() {
		var n int
		for { // 无限循环
			select {
			case <-ctx.Done():
				log.Printf("finished")
				return
			case <-time.After(500 * time.Millisecond): // 默认的时候会做
				ch <- n // 就是往管道里面写新的值
				n += 1
			}
		}
	}()
	return ch
}

func RunContextExampleBreak() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for n := range gen(ctx) { // 因为管道是一次一次阻塞的, 所以可以用range去遍历, 本质上就是 n <- gen(ctx)
		log.Println("get n: ", n)

		if n == 5 {
			cancel()
			break
		}
	}
	log.Println("all finished!")
}
