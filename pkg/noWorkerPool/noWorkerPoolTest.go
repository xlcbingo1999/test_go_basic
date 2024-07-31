package noWorkerPool

import (
	"fmt"
	"log"
	"time"
)

// 业务代码, 实现了Job的接口
type MyJob struct {
	index int
}

func (j *MyJob) Run() {
	log.Println("start index: ", j.index)
	time.Sleep(1 * time.Second)
	log.Println("end index: ", j.index)
}

func chanBufferRunJob() {
	jobQ := make(chan Job, 1000) // 缓冲队列, 避免过载

	go func() {
		for {
			select {
			case j := <-jobQ:
				j.Run()
			case <-time.After(500 * time.Millisecond):
				// 啥也不干
				fmt.Println("time wait 500 ms")
			}
		}
	}()

	// 开始发送一个请求
	job := &MyJob{index: 0}
	jobQ <- job

}

func RunNoWorkerPool() {
	chanBufferRunJob()
}
