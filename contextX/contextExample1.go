package contextX

import (
	"context"
	"fmt"
	"time"
)

func RunContextExample1() {
	// context.Background()是顶级的祖先上下文
	// ctx是新建的子上下文
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // 一秒钟就过期

	defer cancel() // 这个是一个函数变量, 上下文需要关闭的

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main ", ctx.Err())
	default:
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done(): // 接收了一个关闭信号, 一般会出现在上下文过期, 或者手动关闭的时候出现
		fmt.Println("handle ", ctx.Err())
	case <-time.After(duration): // 超过了一段时间后, 就会往这里面写内容
		fmt.Println("process request with ", duration)
	}
}
