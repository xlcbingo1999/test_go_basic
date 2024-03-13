package contextX

import (
	"context"
	"log"
	"time"
)

func calculate() {
	log.Println("beijing")
}

func Perform(ctx context.Context) {
	for {
		calculate()

		select {
		case <-ctx.Done(): // 接收到了结束的执行
			log.Println("finished Perform")
			return
		case <-time.After(500 * time.Millisecond):
			// 这也是一个常见的逻辑, 用<-time.After() 去执行Sleep
		}
	}
}

func RunContextExampleCancel() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	go Perform(ctx) // 启动一个子线程去处理具体的事务, 而子线程的取消是需要由context创建时确定的, 获得父线程决定的

	time.Sleep(20 * time.Second)
	cancel() // 此时可能是用户前端关闭了页面, 主routine接收到具体的情况后就调用这个来关掉子线程
}
