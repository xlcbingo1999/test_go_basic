package workerPool

import (
	"log"
	"runtime"
	"time"
)

type MyPoolJob struct {
	index int
}

// 这里的接收者是指针类型的
func (j *MyPoolJob) Run() {
	log.Println("start index: ", j.index)
	time.Sleep(1 * time.Second)
	log.Println("end index: ", j.index)
}

func RunWorkerPool() {
	poolNum := 100 * 20 // 最多支持20w的并发量
	jobQueueNum := 100
	workerPool := NewWorkerPool(poolNum, jobQueueNum)

	workerPool.Start()

	dataNum := 100 * 100 // 模拟出来的百万业务请求

	go func() {
		for i := 0; i < dataNum; i++ {
			task := &MyPoolJob{index: i}
			workerPool.JobQueue <- task
		}
	}()

	for {
		log.Println("runtime.NumGoroutine(): ", runtime.NumGoroutine())
		time.Sleep(3 * time.Second)
	}
}
